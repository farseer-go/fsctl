// @area api
package {do}App

import (
	"{projectName}/domain/demo"
)

// Hello 演示
// @get {Do}/{action}
func Hello(name string, repository {do}.Repository) string {
	return "hello:" + name
}