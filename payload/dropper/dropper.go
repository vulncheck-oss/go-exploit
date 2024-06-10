// File dropper download and execute payloads.
//
// The dropper package contains all the code for download and execute payloads. Unlike the other payloads
// this package is necessarily OS dependent for both the download and execution portions.
package dropper

type Dropper interface{}

type (
	UnixPayload    struct{}
	WindowsPayload struct{}
	GroovyPayload  struct{}
	PHPPayload     struct{}
)

var (
	Unix    = &UnixPayload{}
	Windows = &WindowsPayload{}
	Groovy  = &GroovyPayload{}
	PHP     = &PHPPayload{}
)
