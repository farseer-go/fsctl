package parse

import (
	"fmt"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/exec"
	"github.com/farseer-go/utils/file"
	"os"
	"strings"
)

const modulePrefix = "module "

// GetRootPackage 得到包名
func GetRootPackage(rootPath string) string {
	receiveOutput := make(chan string, 100)
	exec.RunShell("go list", receiveOutput, nil, rootPath, false)
	result := collections.NewListFromChan(receiveOutput)
	if result.Count() == 0 {
		fmt.Printf(utils.Red("当前目录没有go.mod文件\n"))
	}
	packageName := result.First()
	if strings.Contains(packageName, "go.mod file not found") {
		fmt.Printf(utils.Red("当前目录没有go.mod文件\n"))
	}
	return packageName
}

// ExistsGoMod 判断go.mod文件是否存在
func ExistsGoMod(rootPath string) bool {
	goModPath := rootPath + "go.mod"
	return file.IsExists(goModPath)
}

// GetModRequire 得到依赖包
func GetModRequire(rootPath string) collections.List[string] {
	goModPath := rootPath + "go.mod"
	goModContent := file.ReadAllLines(goModPath)
	if len(goModContent) == 0 {
		fmt.Printf(utils.Red("无法读取go.mod文件\n"))
		os.Exit(0)
	}
	lst := collections.NewList[string]()
	for _, content := range goModContent {
		if strings.HasPrefix(content, "module ") {
			continue
		}
		if strings.Contains(content, "// indirect") {
			continue
		}
		if strings.Contains(content, " v") {
			content = strings.ReplaceAll(content, "\t", "")
			lastIndex := strings.Index(content, " v")
			content = content[:lastIndex]
			lst.Add(content)
		}
	}
	return lst
}
