package protocol

import (
	"net"
	"strconv"

	"github.com/vulncheck-oss/go-exploit/output"
)

func UDPConnect(host string, port int) (*net.UDPConn, bool) {
	target := host + ":" + strconv.Itoa(port)
	output.PrintfFrameworkStatus("Connecting to " + target)
	udpAddr, err := net.ResolveUDPAddr("udp", target)
	if err != nil {
		output.PrintFrameworkError("ResolveUDPAddr failed: " + err.Error())

		return nil, false
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		output.PrintFrameworkError("Connection failed: " + err.Error())

		return nil, false
	}

	return conn, true
}

func UDPWrite(conn *net.UDPConn, data []byte) bool {
	written, err := conn.Write(data)
	if err != nil {
		output.PrintFrameworkError("Server write failed: " + err.Error())

		return false
	}
	if written != len(data) {
		output.PrintFrameworkError("Failed to write all data")

		return false
	}

	return true
}

func UDPReadAmount(conn *net.UDPConn, amount int) ([]byte, bool) {
	reply := make([]byte, amount)
	count, err := conn.Read(reply)
	if err != nil {
		output.PrintFrameworkError("Failed to read from the socket: " + err.Error())

		return nil, false
	}
	if count != amount {
		output.PrintFrameworkError("Failed to read specified amount from the socket")

		return nil, false
	}

	return reply, true
}
