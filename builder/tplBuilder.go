package builder

import (
	"github.com/farseer-go/utils/file"
	"strings"
)

// TplBuilder 根据模板文件，生成文件
func TplBuilder(tplContent string, content map[string]string, path string) {
	c := TplContent(tplContent, content)
	file.WriteString(path, c)
}

// TplContent 根据模板文件，生成内容
func TplContent(tplContent string, content map[string]string) string {
	for k, v := range content {
		tplContent = strings.ReplaceAll(tplContent, k, v)
	}
	return tplContent
}
