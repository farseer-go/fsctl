package builder

import (
	_ "embed"
	"github.com/farseer-go/utils/file"
	"strings"
)

//go:embed tpl/go.mod.tpl
var modTpl string

// ModBuilder 生成go.mod文件
func ModBuilder(path string, projectName string, farseerVer string) {
	contents := strings.ReplaceAll(modTpl, "{projectName}", projectName)
	contents = strings.ReplaceAll(contents, "{farseerVer}", farseerVer)
	file.WriteString(path, contents)
}
