package cmd

import "fmt"

func init() {
	c := route{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type route struct {
}

func (receiver route) Execute(args []string) {
	fmt.Println("当前版本：", Yellow(ver))
}
func (receiver route) FullCommand() string {
	return "route"
}

func (receiver route) ShortCommand() string {
	return "-r"
}

func (receiver route) CommandDesc() string {
	return "显示当前工具版本"
}
