package cmd

import "fmt"

const ver = "v0.8.0"

func init() {
	Commands["-v"] = version{}
	Commands["version"] = version{}
}

type version struct {
}

func (receiver version) Execute(args []string) {
	fmt.Println("当前版本：", Yellow(ver))
}
func (receiver version) FullCommand() string {
	return "version"
}

func (receiver version) ShortCommand() string {
	return "-v"
}

func (receiver version) CommandDesc() string {
	return "显示当前工具版本"
}
