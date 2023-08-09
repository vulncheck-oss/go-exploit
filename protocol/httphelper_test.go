package protocol

import (
	"testing"
)

func TestParseCookies(t *testing.T) {
	cookies := parseCookies([]string{
		"cookie1=foo; path=/",
		"cookie2=bar;",
		"cookie3=baz",
	})

	if cookies != "cookie1=foo; cookie2=bar; cookie3=baz" {
		t.Fatal(cookies)
	}

	t.Log(cookies)
}

func TestBuildURI(t *testing.T) {
	uri := BuildURI("a", "file", "path")

	if uri != "/a/file/path" {
		t.Fatal(uri)
	}

	uri = BuildURI("a", "", "path")
	if uri != "/a/path" {
		t.Fatal(uri)
	}

	uri = BuildURI("")
	if uri != "/" {
		t.Fatal(uri)
	}
}
