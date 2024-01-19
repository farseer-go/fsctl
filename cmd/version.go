package cmd

import (
	"fmt"
	"github.com/farseer-go/fsctl/utils"
)

const ver = "v0.13.0.beta1"
const farseerVer = "v0.12.0"

func init() {
	c := &version{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type version struct {
}

func (receiver *version) Execute(args []string) {
	fmt.Println("工具版本：", utils.Yellow(ver))
	fmt.Println("框架版本：", utils.Yellow(farseerVer))
}
func (receiver *version) FullCommand() string {
	return "version"
}

func (receiver *version) ShortCommand() string {
	return "-v"
}

func (receiver *version) CommandDesc() string {
	return "显示当前工具版本"
}
