package cli

import (
	"bufio"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/protocol"
)

// A very basic reverse/bind shell handler.
func Basic(conn net.Conn) {
	// Create channels for communication between goroutines.
	responseCh := make(chan string)
	quit := make(chan struct{})

	// Use a WaitGroup to wait for goroutines to finish.
	var wg sync.WaitGroup

	// Goroutine to read responses from the server.
	wg.Add(1)
	go func() {
		defer wg.Done()
		responseBuffer := make([]byte, 1024)
		for {
			select {
			case <-quit:
				return
			default:
				_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
				bytesRead, err := conn.Read(responseBuffer)
				if err != nil && !strings.Contains(err.Error(), "i/o timeout") {
					// things have gone sideways, but the command line won't know that
					// until they attempt to execute a command and the socket fails.
					// i think that's largely okay.
					return
				}
				if bytesRead > 0 {
					// I think there is technically a race condition here where the socket
					// could have move data to write, but the user has already called exit
					// below. I that that's tolerable for now.
					responseCh <- string(responseBuffer[:bytesRead])
				}
			}
		}
	}()

	// Goroutine to handle responses and print them.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for response := range responseCh {
			select {
			case <-quit:
				return
			default:
				output.PrintShell(response)
			}
		}
	}()

	for {
		// read user input until they type 'exit\n' or the socket breaks
		// note that ReadString is blocking, so they won't know the socket
		// is broken until they attempt to write something
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		ok := protocol.TCPWrite(conn, []byte(command))
		if !ok || command == "exit\n" {
			break
		}
	}

	// signal for everyone to shutdown
	quit <- struct{}{}
	close(responseCh)

	// wait until the go routines are clean up
	wg.Wait()
}
