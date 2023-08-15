package cmd

func init() {
	c := new{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type new struct {
}

func (receiver new) Execute(args []string) {
}

func (receiver new) FullCommand() string {
	return "new"
}

func (receiver new) ShortCommand() string {
	return "-n"
}

func (receiver new) CommandDesc() string {
	return "新建项目（脚手架）"
}
