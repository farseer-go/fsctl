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
sudo curl -L -o /usr/local/bin/fsctl https://github.com/farseer-go/fsctl/releases/download/v0.10.1/fsctl.$(uname -s).$(uname -m) && sudo chmod +x /usr/local/bin/fsctl
```