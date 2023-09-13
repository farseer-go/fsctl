// @area api
package demoApp

import (
	"{projectName}/domain/demo"
)

// Hello 演示
// repository通过container自动注入实现进来
// @get {area}/hello
// @filter jwt auth
// @message 查询成功
func Hello(name string, repository {do}.Repository) string {
	return "hello:" + name
}