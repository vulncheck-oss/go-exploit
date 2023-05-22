package transform

import (
	"testing"
)

func TestEscapeHTML(t *testing.T) {
	escaped := EscapeHTML("<script>")

	if escaped != "&lt;script&gt;" {
		t.Fatal(escaped)
	}

	t.Log(escaped)
}

func TestEscapeXML(t *testing.T) {
	escaped := EscapeXML("<script>")

	if escaped != "&lt;script&gt;" {
		t.Fatal(escaped)
	}

	t.Log(escaped)
}
