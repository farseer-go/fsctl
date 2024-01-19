# fsctl
farseer-go 编译工具

```shell
* fsctl help	| -h	查看帮助
* fsctl version	| -v	显示当前工具版本
* fsctl new	| -n	新建项目（脚手架）
* fsctl route	| -r	动态路由配置
* fsctl add	| -a	添加领域
* fsctl mod	| -m	更新mod到最新版本
```

## 如何安装
```shell
sudo curl -L -o /usr/local/bin/fsctl "https://github.com/farseer-go/fsctl/releases/download/v0.13.0/fsctl.$(uname -s).$(uname -m)" && sudo chmod +x /usr/local/bin/fsctl
```

## 新建项目
```shell
[root@test ~]# fsctl new testgo
共生成路由：0条
成功...
[root@test ~]# tree testgo/
testgo/
├── application
│   └── module.go
├── domain
│   └── module.go
├── farseer.yaml
├── go.mod
├── infrastructure
│   ├── module.go
│   └── repository
│       ├── context
│       │   ├── mysqlContext.go
│       │   └── redisContext.go
│       └── model
├── interfaces
│   └── module.go
├── main.go
├── route.go
├── startupModule.go
└── wwwroot

8 directories, 11 files
```