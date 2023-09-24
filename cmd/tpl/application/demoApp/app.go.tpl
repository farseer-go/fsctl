// @area api
package {do}App

import (
	"{projectName}/domain/demo"
)

// Hello 演示
// @get {Do}/{action}
// @filter jwt auth
// @message 查询成功
func Hello(name string, repository {do}.Repository) string {
	return "hello:" + name
}