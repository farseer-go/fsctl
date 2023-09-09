package cmd

import (
	_ "embed"
	"fmt"
	"github.com/farseer-go/fsctl/builder"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"os"
)

func init() {
	str, _ := os.Getwd()
	c := &newProject{rootPath: str + "/"}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type newProject struct {
	rootPath          string // 根目录地址
	projectPath       string // 项目地址
	projectName       string // 项目名称
	startupModulePath string // startupModule.go
}

//go:embed tpl/go.mod.tpl
var modTpl string

//go:embed tpl/main.go.tpl
var mainTpl string

//go:embed tpl/startupModule.go.tpl
var startupModuleBuilderTpl string

//go:embed tpl/application/module.go.tpl
var applicationModuleTpl string

//go:embed tpl/domain/module.go.tpl
var domainModuleTpl string

//go:embed tpl/infrastructure/module.go.tpl
var infrastructureModuleTpl string

//go:embed tpl/interfaces/module.go.tpl
var interfacesModuleTpl string

//go:embed tpl/infrastructure/repository/context/mysqlContext.go.tpl
var mysqlContextTpl string

//go:embed tpl/infrastructure/repository/context/redisContext.go.tpl
var redisContextTpl string

//go:embed tpl/infrastructure/model/po.go.tpl
var poTpl string

//go:embed tpl/domain/demo/domainObject.go.tpl
var domainObjectTpl string

//go:embed tpl/domain/demo/repository.go.tpl
var repositoryTpl string

//go:embed tpl/application/demoApp/app.go.tpl
var helloAppTpl string

//go:embed tpl/farseer.yaml.tpl
var farseerYamlTpl string

func (receiver *newProject) Execute(args []string) {
	if len(args) != 3 {
		fmt.Printf(utils.Red("参数不正确，新建项目需要填写项目名称。如：fsctl new project1"))
		os.Exit(0)
	}

	receiver.projectName = args[2]
	receiver.projectPath = receiver.rootPath + args[2] + "/"
	receiver.startupModulePath = receiver.projectPath + "startupModule.go"

	// 项目目录必须为空
	if file.IsExists(receiver.projectPath) {
		file.Delete(receiver.projectPath)
		//fmt.Printf(utils.Red(fmt.Sprintf("目录%s已存在，请先删除该目录", receiver.projectName)))
		//os.Exit(0)
	}

	// 创建目录
	file.CreateDir766(receiver.projectPath)
	file.CreateDir766(receiver.projectPath + "application/demo")
	file.CreateDir766(receiver.projectPath + "domain/demo")
	file.CreateDir766(receiver.projectPath + "infrastructure/repository/context/")
	file.CreateDir766(receiver.projectPath + "infrastructure/repository/model/")
	file.CreateDir766(receiver.projectPath + "interfaces")
	file.CreateDir766(receiver.projectPath + "wwwroot")

	// 模板变量
	tplValue := map[string]string{
		"{projectName}": receiver.projectName,
		"{farseerVer}":  farseerVer,
		"{Do}":          "Demo",
		"{do}":          "demo",
	}
	// domain
	builder.TplBuilder(domainModuleTpl, tplValue, receiver.projectPath+"domain/module.go")            // domain/module.go
	builder.TplBuilder(domainObjectTpl, tplValue, receiver.projectPath+"domain/demo/domainObject.go") // domain/demo/domainObject.go
	builder.TplBuilder(repositoryTpl, tplValue, receiver.projectPath+"domain/demo/repository.go")     // domain/demo/repository.go
	// infrastructure
	builder.TplBuilder(infrastructureModuleTpl, tplValue, receiver.projectPath+"infrastructure/module.go")                  // infrastructure/module.go
	builder.TplBuilder(mysqlContextTpl, tplValue, receiver.projectPath+"infrastructure/repository/context/mysqlContext.go") // infrastructure/repository/context/mysqlContext.go
	builder.TplBuilder(redisContextTpl, tplValue, receiver.projectPath+"infrastructure/repository/context/redisContext.go") // infrastructure/repository/context/redisContext.go
	builder.TplBuilder(poTpl, tplValue, receiver.projectPath+"infrastructure/repository/model/demoPO.go")                   // infrastructure/repository/model/demoPO.go
	// application
	builder.TplBuilder(applicationModuleTpl, tplValue, receiver.projectPath+"application/module.go") // application/module.go
	builder.TplBuilder(helloAppTpl, tplValue, receiver.projectPath+"application/demo/app.go")        // application/demo/app.go
	// interfaces
	builder.TplBuilder(interfacesModuleTpl, tplValue, receiver.projectPath+"interfaces/module.go") // interfaces/module.go

	// 根目录
	builder.TplBuilder(modTpl, tplValue, receiver.projectPath+"go.mod")                            // go.mod
	builder.TplBuilder(mainTpl, tplValue, receiver.projectPath+"main.go")                          // main.mod
	builder.TplBuilder(startupModuleBuilderTpl, tplValue, receiver.projectPath+"startupModule.go") // startupModule.mod
	builder.TplBuilder(farseerYamlTpl, tplValue, receiver.projectPath+"farseer.yaml")              // farseer.yaml
	// rotue.go
	r := route{rootPath: receiver.projectPath, routePath: "route.go"}
	r.Execute(args)
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
