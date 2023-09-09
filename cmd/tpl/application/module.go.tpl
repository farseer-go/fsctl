package application

import (
	"github.com/farseer-go/fs/modules"
	"{projectName}/domain"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{domain.Module{}}
}

func (module Module) PostInitialize() {
}
