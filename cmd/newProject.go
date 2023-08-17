package cmd

import (
	"fmt"
	"github.com/farseer-go/fsctl/builder"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"os"
)

func init() {
	c := &newProject{rootPath: "../demo/"}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type newProject struct {
	rootPath          string // 根目录地址
	projectPath       string // 项目地址
	projectName       string // 项目名称
	modPath           string // go.mod
	mainPath          string // main.go
	startupModulePath string // startupModule.go
}

func (receiver *newProject) Execute(args []string) {
	if len(args) != 3 {
		fmt.Printf(utils.Red("参数不正确，新建项目需要填写项目名称。如：fsctl new project1"))
		os.Exit(0)
	}

	// 项目目录必须为空
	receiver.projectName = args[2]
	receiver.projectPath = receiver.rootPath + args[2] + "/"
	receiver.modPath = receiver.projectPath + "go.mod"
	receiver.mainPath = receiver.projectPath + "main.go"
	receiver.startupModulePath = receiver.projectPath + "startupModule.go"

	if file.IsExists(receiver.projectPath) {
		fmt.Printf(utils.Red(fmt.Sprintf("目录%s已存在，请先删除该目录", receiver.projectName)))
		os.Exit(0)
	}

	// 创建目录
	file.CreateDir766(receiver.projectPath)
	builder.ModBuilder(receiver.modPath, receiver.projectName, farseerVer) // go.mod
}

func (receiver *newProject) FullCommand() string {
	return "new"
}

func (receiver *newProject) ShortCommand() string {
	return "-n"
}

func (receiver *newProject) CommandDesc() string {
	return "新建项目（脚手架）"
}
