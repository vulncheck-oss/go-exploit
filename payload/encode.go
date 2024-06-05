// Payload related functions and actions
//
// The payload package contains a collection of universally applicable functions for payloads, sub-packages
// containing specific payloads, and any specific payloads that do not fit into the other sub package types.
package payload

import (
	"regexp"
)

func EncodeCommandBrace(cmd string) string {
	escaped := regexp.MustCompile(`([{,}])`).ReplaceAllString(cmd, `\$1`)

	return "{" + regexp.MustCompile(`\s+`).ReplaceAllString(escaped, ",") + "}"
}

func EncodeCommandIFS(cmd string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllLiteralString(cmd, "${IFS}")
}
