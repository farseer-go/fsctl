package builder

import (
	_ "embed"
	"github.com/farseer-go/utils/file"
	"strings"
)

//go:embed tpl/route.go.tpl
var routeTpl string

//go:embed tpl/routeItem.tpl
var routeItemTpl string

// RouteBuilder 生成route.go文件
func RouteBuilder(path string, imports string, routeItemBuilder func(routeItemTpl string) string) {
	contents := strings.ReplaceAll(routeTpl, "{import}", imports)
	contents = strings.ReplaceAll(contents, "{route}", routeItemBuilder(routeItemTpl))
	if file.ReadString(path) != contents {
		file.WriteString(path, contents)
	}
}
