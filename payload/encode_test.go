package payload

import (
	"testing"
)

func TestEncodeCommandBrace(t *testing.T) {
	encoded := EncodeCommandBrace("foo bar baz")

	if encoded != "{foo,bar,baz}" {
		t.Fatal(encoded)
	}

	t.Log(encoded)
}

func TestEncodeCommandIFS(t *testing.T) {
	encoded := EncodeCommandIFS("foo bar baz")

	if encoded != "foo${IFS}bar${IFS}baz" {
		t.Fatal(encoded)
	}

	t.Log(encoded)
}

func TestEncodeEchoBase64ToBash(t *testing.T) {
	encoded := EncodeEchoBase64ToBash("whoami; id;")

	if encoded != "echo d2hvYW1pOyBpZDs=|base64 -d|bash" {
		t.Fatal(encoded)
	}

	t.Log(encoded)
}

func TestEncodeSelfRemovingCron(t *testing.T) {
	cron, xploit := SelfRemovingCron("root", "/etc/cron.d/hi", "/tmp/test", "id")

	if cron != "* * * * * root /bin/sh /tmp/test\n" {
		t.Fatal(cron)
	}

	if xploit != "#!/bin/sh\n\nrm -f /etc/cron.d/hi\nrm -f /tmp/test\nid\n" {
		t.Fatal(xploit)
	}
}
