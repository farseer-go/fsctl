package repository

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"{projectName}/domain/{do}"
	"{projectName}/infrastructure/repository/context"
)

// Init{Do} 初始化仓储
func Init{Do}() {
	container.Register(func() {do}.Repository {
		return &{Do}Repository{}
	})
}

// {Do}Repository 仓储实现
// @register order.Repository
type {Do}Repository struct {
	data.IRepository[{do}.DomainObject]
}
