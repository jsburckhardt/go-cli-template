// main package for the sample
package main

import (
	"sample/cmd"
)

var (
	version = "edge"
	commit  = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
