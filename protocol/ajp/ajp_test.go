package ajp

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestStructCreation(t *testing.T) {
	req := createForwardRequest("127.0.0.1", 80, false, "/hello", []string{"accept", "text/html"}, []string{})
	if req.protocol != "HTTP/1.1" {
		t.Error("Unexpected protocol specified")
	}

	// loop through the headers looking for the defaults
	foundHost := false
	foundAccept := false
	fmt.Println(req.headers)
	for _, header := range req.headers {
		if strings.HasPrefix(header, "accept") {
			foundAccept = true
		}
		if strings.HasPrefix(header, "host") {
			foundHost = true
		}
	}

	if !foundHost {
		t.Error("Missing Host header")
	}
	if !foundAccept {
		t.Error("Missing Accept header")
	}
}

func TestStructSerialize(t *testing.T) {
	req := createForwardRequest("127.0.0.1", 80, false, "/hello", []string{}, []string{})
	setGetForwardRequest(&req)
	serialized := serializeForwardRequest(req)

	if serialized[0] != 0x12 || serialized[1] != 0x34 {
		t.Error("Invalid magic")
	}

	if serialized[4] != 0x02 {
		t.Errorf("Invalid code: %d", serialized[4])
	}

	if serialized[5] != 0x02 {
		t.Errorf("Invalid method: %d", serialized[5])
	}

	if serialized[6] != 0x00 && serialized[7] != 0x08 {
		t.Errorf("Invalid protocol length")
	}

	if !bytes.Equal(serialized[8:16], []byte("HTTP/1.1")) {
		t.Errorf("Invalid protocol version %s", string(serialized[8:16]))
	}
}
