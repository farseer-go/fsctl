package cmd

import "fmt"

func init() {
	c := err{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type err struct {
}

func (receiver err) Execute(args []string) {
	fmt.Println("参数错误，可以输入：fsctl help 查看更多信息")
}

func (receiver err) FullCommand() string {
	return "err"
}

func (receiver err) ShortCommand() string {
	return ""
}

func (receiver err) CommandDesc() string {
	return "错误的命令提示"
}
