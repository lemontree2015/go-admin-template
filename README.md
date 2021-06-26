# go-admin-template
该项目使用golang + Vue搭建前后端分离框架，包含整套后台系统基础功能，可做为后台基础框架搭建后台管理系统。

## 后端Golang使用到的插件
-  [Gin](https://github.com/gin-gonic/gin) (web api接口)
-  [Jwt](https://github.com/dgrijalva/jwt-go) (用户登陆认证)
-  [Casbin](https://github.com/casbin/casbin) (角色权限管理)
-  [Gorm](https://github.com/jinzhu/gorm) (数据库，在此基础上可拓展分库、分表、读写分离等)
-  [Cli](https://github.com/urfave/cli) (命令行运行插件)
-  [Logrus](https://github.com/sirupsen/logrus) (日志插件)

## 运行环境搭建
1.拉取代码

```sh
$ git clone https://github.com/lemontree2015/go-admin-template.git
```

2.安装拓展
```sh
$ cd go-admin-template
$ go mod tidy
$ go mod vendor
```

3.创建数据库并导入测试数据
```sh
$ use dbname
$ source sql文件
```

4.修改conf目录配制文件

## 项目构建
构建Linux环境执行文件：
```sh
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../bin/go_admin .
```

## 运行
```sh
$ cd bin
$ ./go_admin mgr -c ../conf/local.toml -m ../conf/model.conf
$ ./go_admin api -c ../conf/local.toml
```


## 前端界面基于[vue-admin-template](https://github.com/PanJiaChen/vue-admin-template)，权限控制可精细到按钮级别，前端组件参考:
-  [ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue) 

## 线上效果预览
> admin  /  123456(管理员)

> test  /  1234(测试账号)

演示地址：[http://demo.sscmgroup.com/](http://demo.sscmgroup.com/)

## 如有疑问联系微信
![image](http://img1.sscmgroup.com/avatar/wx.jpg)
