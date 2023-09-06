package transform

import (
	"encoding/base64"
	"encoding/binary"
	"strings"

	"github.com/vulncheck-oss/go-exploit/output"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		output.PrintFrameworkError(err.Error())

		return ""
	}

	return string(decoded)
}

func Title(s string) string {
	return cases.Title(language.Und, cases.NoLower).String(s)
}

// PackLittleInt32 packs a little-endian 32-bit integer as a string.
func PackLittleInt32(n int) string {
	var packed strings.Builder

	err := binary.Write(&packed, binary.LittleEndian, int32(n))
	if err != nil {
		output.PrintFrameworkError(err.Error())
	}

	return packed.String()
}

// PackLittleInt64 packs a little-endian 64-bit integer as a string.
func PackLittleInt64(n int) string {
	var packed strings.Builder

	err := binary.Write(&packed, binary.LittleEndian, int64(n))
	if err != nil {
		output.PrintFrameworkError(err.Error())
	}

	return packed.String()
}

// PackBigInt16 packs a big-endian 16-bit integer as a string.
func PackBigInt16(n int) string {
	var packed strings.Builder

	err := binary.Write(&packed, binary.BigEndian, int16(n))
	if err != nil {
		output.PrintFrameworkError(err.Error())
	}

	return packed.String()
}

// PackBigInt32 packs a big-endian 32-bit integer as a string.
func PackBigInt32(n int) string {
	var packed strings.Builder

	err := binary.Write(&packed, binary.BigEndian, int32(n))
	if err != nil {
		output.PrintFrameworkError(err.Error())
	}

	return packed.String()
}
