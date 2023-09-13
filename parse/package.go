package parse

import (
	"fmt"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"os"
	"strings"
)

const modulePrefix = "module "

// GetRootPackage 得到包名
func GetRootPackage(rootPath string) string {
	goModPath := rootPath + "go.mod"
	goModContent := file.ReadAllLines(goModPath)
	if len(goModContent) == 0 {
		fmt.Printf(utils.Red("无法读取go.mod文件\n"))
		os.Exit(0)
	}
	for _, content := range goModContent {
		if strings.HasPrefix(content, modulePrefix) {
			return content[len(modulePrefix):]
		}
	}
	return ""
}

// ExistsGoMod 判断go.mod文件是否存在
func ExistsGoMod(rootPath string) bool {
	goModPath := rootPath + "go.mod"
	return file.IsExists(goModPath)
}
