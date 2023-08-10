package parse

import (
	"fmt"
	"github.com/farseer-go/fsctl/utils"
	"github.com/farseer-go/utils/file"
	"os"
	"strings"
)

const modulePrefix = "module "

func GetRootPackage(rootPath string) string {
	goModPath := rootPath + "go.mod"
	goModContent := file.ReadAllLines(goModPath)
	if len(goModContent) == 0 {
		fmt.Printf(utils.Red("无法读取go.mod文件"))
		os.Exit(0)
	}
	for _, content := range goModContent {
		if strings.HasPrefix(content, modulePrefix) {
			return content[len(modulePrefix):]
		}
	}
	return ""
}
