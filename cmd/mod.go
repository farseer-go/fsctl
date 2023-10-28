package cmd

import (
	"fmt"
	"github.com/farseer-go/fsctl/parse"
	"github.com/farseer-go/fsctl/utils"
	"os"
)

// 更新当前项目的所有Mod
func init() {
	str, _ := os.Getwd()
	c := &mod{projectPath: str + "/", receiveOutput: make(chan string, 100)}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type mod struct {
	projectPath   string // 项目根目录
	receiveOutput chan string
}

func (receiver *mod) Execute(args []string) {
	go receiver.print()
	lst := parse.GetModRequire(receiver.projectPath)
	lst.For(func(index int, packagePath *string) {
		cmd := "go get " + *packagePath
		fmt.Printf("%d/%d %s\n", index+1, lst.Count(), utils.Blue(cmd))
		utils.RunShell(cmd, receiver.receiveOutput, nil, receiver.projectPath)
	})
}

func (receiver *mod) print() {
	for output := range receiver.receiveOutput {
		fmt.Println(output)
	}
}
func (receiver *mod) FullCommand() string {
	return "mod"
}

func (receiver *mod) ShortCommand() string {
	return "-m"
}

func (receiver *mod) CommandDesc() string {
	return "更新mod到最新版本"
}
