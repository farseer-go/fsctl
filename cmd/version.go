package cmd

import (
	"fmt"

	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fsctl/utils"
)

func init() {
	c := &version{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type version struct {
}

func (receiver *version) Execute(args []string) {
	fmt.Println("工具版本：", utils.Yellow(core.Version))
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
