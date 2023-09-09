package cmd

import (
	_ "embed"
	"fmt"
	"github.com/farseer-go/fsctl/builder"
	"github.com/farseer-go/fsctl/parse"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"os"
	"strings"
)

func init() {
	str, _ := os.Getwd()
	c := &add{projectPath: str + "/"}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type add struct {
	projectPath string // 根目录地址
	lowerName   string // 首字母小写
	upperName   string // 首字母大写
	projectName string // 项目名称
	contextPath string // 数据库上下文
}

//go:embed tpl/domain/demo/domainObject.go.tpl
var domainObjectTpl string

//go:embed tpl/domain/demo/repository.go.tpl
var repositoryTpl string

//go:embed tpl/infrastructure/model/po.go.tpl
var poTpl string

//go:embed tpl/application/demoApp/app.go.tpl
var helloAppTpl string

//go:embed tpl/infrastructure/repository/context/domainSet.tpl
var domainSetTpl string

// Execute fsctl add user
func (receiver *add) Execute(args []string) {
	if len(args) != 3 {
		fmt.Printf(utils.Red("参数不正确，添加新的领域命令，如：fsctl add xxx"))
		os.Exit(0)
	}
	receiver.check()
	receiver.lowerName = utils.FirstLower(args[2])
	receiver.upperName = utils.FirstUpper(args[2])
	receiver.projectName = parse.GetRootPackage(receiver.projectPath)
	receiver.contextPath = receiver.projectPath + "infrastructure/repository/context/mysqlContext.go"

	// 模板变量
	tplValue := map[string]string{
		"{projectName}": receiver.projectName,
		"{farseerVer}":  farseerVer,
		"{Do}":          receiver.upperName,
		"{do}":          receiver.lowerName,
	}

	// domain
	file.CreateDir766(receiver.projectPath + "domain/" + receiver.lowerName)
	builder.TplBuilder(domainObjectTpl, tplValue, receiver.projectPath+"domain/"+receiver.lowerName+"/domainObject.go") // domain/lowerName/domainObject.go
	builder.TplBuilder(repositoryTpl, tplValue, receiver.projectPath+"domain/"+receiver.lowerName+"/repository.go")     // domain/lowerName/repository.go

	// infrastructure
	file.CreateDir766(receiver.projectPath + "infrastructure/repository/model")
	builder.TplBuilder(poTpl, tplValue, receiver.projectPath+"infrastructure/repository/model/"+receiver.lowerName+"PO.go") // infrastructure/repository/model/demoPO.go
	// 找到上下文，然后添加DomainSet
	contextContent := file.ReadAllLines(receiver.contextPath)
	if len(contextContent) > 0 {
		contextBeginIndex := -1
		contextEndIndex := -1
		for i := 0; i < len(contextContent); i++ {
			// 找到mysqlContext的定义
			if strings.HasPrefix(contextContent[i], "type ") && strings.HasSuffix(contextContent[i], "Context struct {") {
				contextBeginIndex = i
				continue
			}
			// 找到mysqlContext的结束符号
			if contextBeginIndex > -1 && contextContent[i] == "}" {
				contextEndIndex = i
				break
			}
		}
		// 找到了，则在末尾插入DomainSet
		if contextEndIndex > -1 {
			domainSetContent := builder.TplContent(domainSetTpl, tplValue)
			contextContent = append(contextContent[:contextEndIndex], append([]string{domainSetContent}, contextContent[contextEndIndex:]...)...)
			file.WriteString(receiver.contextPath, strings.Join(contextContent, "\n"))
		}
	}
	// application
	file.CreateDir766(receiver.projectPath + "application/" + receiver.lowerName)
	builder.TplBuilder(helloAppTpl, tplValue, receiver.projectPath+"application/"+receiver.lowerName+"/app.go") // application/demo/app.go
}

func (receiver *add) FullCommand() string {
	return "add"
}

func (receiver *add) ShortCommand() string {
	return "-a"
}

func (receiver *add) CommandDesc() string {
	return "添加领域"
}

// 检查是否在项目的根目录中
func (receiver *add) check() {
	if !file.IsExists(receiver.projectPath + "go.mod") {
		fmt.Printf(utils.Red("当前目录并不是go项目，请到go项目的根目录中重新执行fsctl命令"))
		os.Exit(0)
	}
}
