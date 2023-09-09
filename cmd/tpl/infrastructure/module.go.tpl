package infrastructure

import (
	"github.com/farseer-go/data"
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/queue"
	"github.com/farseer-go/redis"
	"{projectName}/infrastructure/repository/context"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	// 使用到了redis模块、data(orm)模块、eventBus（事件总线）模块、queue（本地队列）模块
	// 这些模块都是farseer-go内置的模块
	return []modules.FarseerModule{redis.Module{}, data.Module{}, eventBus.Module{}, queue.Module{}}
}

func (module Module) PostInitialize() {
	// 初始化数据库上下文
	context.InitMysqlContext()
}
