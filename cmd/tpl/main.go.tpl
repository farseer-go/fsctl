package main

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/webapi"
)

func main() {
	fs.Initialize[StartupModule]("{projectName}")

	webapi.RegisterRoutes(route...)

	// 让所有的返回值，包含在core.ApiResponse中
	webapi.UseApiResponse()
	// 使用静态文件 在根目录./wwwroot中的文件
	webapi.UseStaticFiles()
	// 运行web服务，端口配置在：farseer.yaml Webapi.Url 配置节点
	webapi.Run()
}
