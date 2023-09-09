package main

import (
	"fmt"
	"github.com/farseer-go/fsctl/cmd"
	"os"
)

func main() {
	// 参数个数不对
	if len(os.Args) == 1 {
		cmd.Commands["err"].Execute(os.Args)
		return
	}

	c, isExists := cmd.Commands[os.Args[1]]
	// 没有找到对应的命令
	if !isExists {
		cmd.Commands["err"].Execute(os.Args)
		return
	}
	// 命令正确，执行命令
	c.Execute(os.Args)

	fmt.Print("成功...")
}
