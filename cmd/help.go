package cmd

import (
	"fmt"
	"github.com/farseer-go/fsctl/utils"
	"strings"
)

func init() {
	c := &help{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type help struct {
}

func (receiver *help) Execute(args []string) {
	Commands["-v"].Execute(args)
	for k, c := range Commands {
		if strings.HasPrefix(k, "-") {
			fmt.Printf("fsctl %s\t| %s\t%s\r\n", utils.Red(c.FullCommand()), utils.Blue(c.ShortCommand()), utils.Green(c.CommandDesc()))
		}
	}
}

func (receiver *help) FullCommand() string {
	return "help"
}

func (receiver *help) ShortCommand() string {
	return "-h"
}

func (receiver *help) CommandDesc() string {
	return "查看帮助"
}
