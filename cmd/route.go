package cmd

import (
	"fmt"
	"github.com/farseer-go/fsctl/parse"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"go/ast"
	"os"
	"strings"
)

func init() {
	str, _ := os.Getwd()
	c := &route{projectPath: str + "/"}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type route struct {
	projectPath string
	routePath   string
}

func (receiver *route) Execute(args []string) {
	receiver.routePath = receiver.projectPath + "route.go"
	receiver.checkRoute()

	var routeComments []parse.RouteComment
	// 解析整个项目
	parse.ASTDir(receiver.projectPath, func(filePath string, astFile *ast.File, funcDecl *ast.FuncDecl) {
		if funcDecl.Doc == nil {
			return
		}
		rc := parse.RouteComment{IocNames: make(map[string]string)}

		// 解析头部注解：区域
		if astFile.Doc != nil {
			for _, comment := range astFile.Doc.List {
				// 得到注解
				ant := parse.GetAnnotation(comment.Text)
				// 解析
				rc.ParsePackageComment(ant)
			}
		}
		// 解析是否有注解
		for _, comment := range funcDecl.Doc.List {
			// 得到注解
			ant := parse.GetAnnotation(comment.Text)
			// 解析
			rc.ParseFuncComment(ant)
		}

		// 解析成功
		if rc.IsHaveComment() {
			// 移除相对路径和文件名，得到包路径
			rc.PackagePath = parse.GetRootPackage(receiver.projectPath) + "/" + filePath[len(receiver.projectPath):strings.LastIndex(filePath, "/")]
			// 解析函数类型
			rc.ParseFuncType(astFile, funcDecl)
			if rc.Area != "" {
				rc.Area = strings.TrimPrefix(rc.Area, "/")
				rc.Area = strings.TrimSuffix(rc.Area, "/")
			}

			fmt.Printf("找到路由：area=%s, [%s]%s ==> %s.%s\n", rc.Area, rc.Method, rc.Url, rc.PackageName, rc.FuncName)

			rc.Url = strings.TrimPrefix(rc.Url, "/")
			rc.Url = "/" + rc.Area + "/" + rc.Url
			rc.Url = strings.Replace(rc.Url, "{action}", rc.FuncName, -1)
			routeComments = append(routeComments, rc)
		}
	})

	// 生成route.go文件
	parse.BuildRoute(receiver.routePath, routeComments)
}

func (receiver *route) FullCommand() string {
	return "route"
}

func (receiver *route) ShortCommand() string {
	return "-r"
}

func (receiver *route) CommandDesc() string {
	return "动态路由配置"
}

// 检查根目录route.go文件是否为fsctl工具生成
func (receiver *route) checkRoute() {
	if file.IsExists(receiver.routePath) && !parse.CheckIsRoute(receiver.routePath) {
		if file.ReadString(receiver.routePath) != "" {
			fmt.Printf(utils.Red("route.go文件不是fsctl工具生成，请手动删除./route.go后再重新运行此命令\n"))
			os.Exit(0)
		}
	}
	// 删除route.go文件
	file.Delete(receiver.routePath)
}
