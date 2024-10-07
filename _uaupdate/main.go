package main

import (
	"os"
	"regexp"

	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/protocol"
)

func main() {
	uri := "https://www.whatismybrowser.com/guides/the-latest-user-agent/chrome"
	resp, body, ok := protocol.HTTPSendAndRecv("GET", uri, "")
	if !ok {
		return
	}

	if resp.StatusCode != 200 {
		output.PrintfError("Unexpected status code: %d %s", resp.StatusCode, body)

		return
	}

	// looking in the body for the latest Chrome on Windows whatever
	matches := regexp.MustCompile(`<li><span class="code">(Mozilla/\d+.\d+ \(Windows NT [^<]+)</span></li>`).FindStringSubmatch(body)
	if len(matches) != 0 {
		_ = os.WriteFile("../protocol/http-user-agent.txt", []byte(matches[1]), 0o644)
	}
}
