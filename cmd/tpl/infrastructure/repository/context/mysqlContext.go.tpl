package context

import (
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/core"
	"{projectName}/domain/demo"
	"{projectName}/infrastructure/repository/model"
)

// MysqlContext 初始化数据库上下文
var MysqlContext *mysqlContext

// mysqlContext 数据库上下文
type mysqlContext struct {
	// 手动使用事务时必须定义
	core.ITransaction
	// 获取原生ORM框架（不使用TableSet或DomainSet）
	data.IInternalContext
	// 定义数据库表 订单 映射TableSet
	Demo data.DomainSet[model.DemoPO, demo.DomainObject] `data:"name=farseer_go_demo"`
}

// InitMysqlContext 初始化上下文
func InitMysqlContext() {
	MysqlContext = data.NewContext[mysqlContext]("default", true)
}
