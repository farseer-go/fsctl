package cmd

type ICmd interface {
	Execute(args []string)
	FullCommand() string  // 完整命令
	ShortCommand() string // 简写命令
	CommandDesc() string  // 命令描述
}

// Commands 命令列表
var Commands = make(map[string]ICmd)
