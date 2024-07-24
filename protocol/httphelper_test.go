package protocol

import (
	"net/http"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	auth := BasicAuth("foo", "bar")

	if auth != "Basic Zm9vOmJhcg==" {
		t.Fatal(auth)
	}

	t.Log(auth)
}

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

func TestGenerateURL(t *testing.T) {
	uri := GenerateURL("google.com", 443, true, "/")
	if uri != "https://google.com:443/" {
		t.Fatal(uri)
	}
	uri = GenerateURL("google.com", 443, true, "/helloworld")
	if uri != "https://google.com:443/helloworld" {
		t.Fatal(uri)
	}
	uri = GenerateURL("::1", 1270, false, "/")
	if uri != "http://[::1]:1270/" {
		t.Fatal(uri)
	}
}

func TestCookieString(t *testing.T) {
	// Normally you might not want duplicates, but there are common bugs with that handling
	// so it might be wanted for hacks.
	testCookies := []*http.Cookie{
		{Name: "testname", Value: "testvalue"},
		{Name: "testname", Value: "testvalue"},
		{Name: "test2name", Value: "test2value"},
		{Name: "stuff", Value: "stuff"},
	}
	cookieStr := CookieString(testCookies)
	if cookieStr != `testname=testvalue; testname=testvalue; test2name=test2value; stuff=stuff;` {
		t.Fatal(cookieStr)
	}
}
