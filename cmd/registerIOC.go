package cmd

func init() {
	c := registerIOC{}
	Commands[c.ShortCommand()] = c
	Commands[c.FullCommand()] = c
}

type registerIOC struct {
}

func (receiver registerIOC) Execute(args []string) {
}

func (receiver registerIOC) FullCommand() string {
	return "ioc"
}

func (receiver registerIOC) ShortCommand() string {
	return "-i"
}

func (receiver registerIOC) CommandDesc() string {
	return "注册IOC（脚手架）"
}
