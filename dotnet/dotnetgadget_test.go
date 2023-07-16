package dotnet

import (
	"testing"
)

func TestTextFormattingRunPropertiesBinaryFormatter(t *testing.T) {
	want, err := ReadGadget("TextFormattingRunProperties", "BinaryFormatter")
	if err != nil {
		t.Fatal(err)
	}

	// Dynamically test the placeholder command
	got := TextFormattingRunPropertiesBinaryFormatter("mspaint.exe")

	if got != string(want) {
		t.Fatalf("%q", got)
	}

	t.Logf("%q", got)
}
