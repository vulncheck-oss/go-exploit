package payload_test

import (
	"testing"

	"github.com/vulncheck-oss/go-exploit/payload"
)

func TestEncodeCommandBrace(t *testing.T) {
	encoded := payload.EncodeCommandBrace("foo bar baz")

	if encoded != "{foo,bar,baz}" {
		t.Fatal(encoded)
	}

	t.Log(encoded)
}

func TestEncodeCommandIFS(t *testing.T) {
	encoded := payload.EncodeCommandIFS("foo bar baz")

	if encoded != "foo${IFS}bar${IFS}baz" {
		t.Fatal(encoded)
	}

	t.Log(encoded)
}
