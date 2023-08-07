package cmd

import (
	"fmt"
	"strings"
)

func init() {
	Commands["-h"] = help{}
	Commands["help"] = help{}
}

type help struct {
}

func (receiver help) Execute(args []string) {
	Commands["-v"].Execute(args)
	for k, c := range Commands {
		if strings.HasPrefix(k, "-") {
			fmt.Printf("fsctl %s\t| %s\t%s\r\n", Red(c.FullCommand()), Blue(c.ShortCommand()), Green(c.CommandDesc()))
		}
	}
}

func (receiver help) FullCommand() string {
	return "help"
}

func (receiver help) ShortCommand() string {
	return "-h"
}

func (receiver help) CommandDesc() string {
	return "查看帮助"
}
