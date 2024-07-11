package cmd

import (
	"fmt"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/path"
	"github.com/farseer-go/fsctl/parse"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"go/ast"
	"os"
	"strings"
)

// 自动配置路由
func init() {
	str, _ := os.Getwd()
	c := &route{projectPath: str + path.PathSymbol}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type route struct {
	projectPath    string // 项目根目录
	routePath      string // route.go路径
	topPackageName string // 顶级包名
}

func (receiver *route) Execute(args []string) {
	receiver.routePath = receiver.projectPath + "route.go"
	receiver.topPackageName = parse.GetRootPackage(receiver.projectPath)
	receiver.checkRoute()

	codeReviewCount := 0
	var routeComments []parse.RouteComment
	// 解析整个项目
	parse.AstDirFuncDecl(receiver.projectPath, func(filePath string, astFile *ast.File, funcDecl *ast.FuncDecl) {
		// 统计转成linux目录格式
		filePath = strings.ReplaceAll(filePath, path.PathSymbol, "/")
		if funcDecl.Doc == nil {
			return
		}
		rc := parse.RouteComment{IocNames: make(map[string]string), ProjectPath: receiver.projectPath, TopPackageName: receiver.topPackageName}
		// 解析头部注解：区域
		if astFile.Doc != nil {
			for _, comment := range astFile.Doc.List {
				// 得到注解
				ant := parse.GetAnnotation(comment.Text)
				// 解析
				rc.ParsePackageComment(ant)
			}
		}
		// 路由地址 Key = {area}/order/{action}, Value = Method (get/post/put/delete)
		mapRoute := make(map[string]string)
		// 解析是否有注解
		for _, comment := range funcDecl.Doc.List {
			// 得到注解
			ant := parse.GetAnnotation(comment.Text)
			// 解析
			rc.ParseFuncComment(ant, mapRoute)
		}

		// 解析成功
		if rc.IsHaveComment(mapRoute) {
			// 移除相对路径和文件名，得到包路径
			rc.PackagePath = receiver.topPackageName + "/" + filePath[len(receiver.projectPath):strings.LastIndex(filePath, "/")]
			// 解析函数类型
			rc.ParseFuncType(astFile, funcDecl)
			if rc.Area != "" {
				rc.Area = strings.TrimPrefix(rc.Area, "/")
				rc.Area = strings.TrimSuffix(rc.Area, "/")
			}
			if !strings.HasPrefix(rc.Area, "/") {
				rc.Area = "/" + rc.Area
			}
			if !strings.HasSuffix(rc.Area, "/") {
				rc.Area = rc.Area + "/"
			}

			for url, method := range mapRoute {
				url = strings.TrimPrefix(url, "/")
				url = rc.Area + url
				url = strings.Replace(url, "{action}", rc.FuncName, -1)
				if rc.CodeReview {
					fmt.Printf("找到路由%s：area=%s, [%s]%s ==> %s.%s\n", flog.Red("（已审核）"), rc.Area, flog.Green(method), url, flog.Blue(rc.PackageName), flog.Blue(rc.FuncName))
				} else {
					fmt.Printf("找到路由：area=%s, [%s]%s ==> %s.%s\n", rc.Area, flog.Green(method), url, flog.Blue(rc.PackageName), flog.Blue(rc.FuncName))
				}
				rc.Url = url
				rc.Method = method
				routeComments = append(routeComments, rc)

				// 统计代码审核过的数量
				if rc.CodeReview {
					codeReviewCount++
				}
			}
		}
	})

	// 生成route.go文件
	parse.BuildRoute(receiver.routePath, routeComments)

	rate := fmt.Sprintf("%.0f", float64(codeReviewCount)/float64(len(routeComments))*100)
	fmt.Printf("共生成路由：%s条，代码审核：%s条（%s%%）\n", utils.Red(len(routeComments)), utils.Blue(codeReviewCount), utils.Green(rate))
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
