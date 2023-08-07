package main

import (
	"github.com/farseer-go/fsctl/cmd"
	"os"
)

// fsctl route
// fsctl ioc
// fsctl -r
func main() {
	args := os.Args
	if len(args) == 1 {
		cmd.Commands["err"].Execute(args)
		return
	}

	c, isExists := cmd.Commands[args[1]]
	if !isExists {
		cmd.Commands["err"].Execute(args)
		return
	}
	c.Execute(args)
}
