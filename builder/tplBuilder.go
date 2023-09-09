package builder

import (
	"github.com/farseer-go/utils/file"
	"strings"
)

// TplBuilder 根据模板文件，生成文件
func TplBuilder(tplContent string, content map[string]string, path string) {
	for k, v := range content {
		tplContent = strings.ReplaceAll(tplContent, k, v)
	}
	file.WriteString(path, tplContent)
}
