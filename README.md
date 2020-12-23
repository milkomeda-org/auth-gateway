![](https://img.shields.io/github/stars/milkomeda-org/auth-gateway)
![](https://img.shields.io/github/issues/milkomeda-org/auth-gateway)
![](https://img.shields.io/github/license/milkomeda-org/auth-gateway)

# auth-gateway
> 权限网关

## Go Mod

```shell
go mod init auth-gateway
export GOPROXY=http://mirrors.aliyun.com/goproxy/
```

## 运行

```shell
go run cmd/server.go
```

## 概述
本项目采用Go语言进行开发，权限架构设计及引用采用RBC，秉承包含完整权限管理的鼓励二开的网关性质应用
在每一个包下都有改包的详细帮助文档，请翻阅，这里只做大概介绍

## 设计标准
- 领域足够明确，只包含用户信息管理，授权鉴权处理，不包含业务逻辑处理，不准涉足他处
- 逻辑足够清晰，不能包含任何复杂重复的操作，必须达到开箱即用
- 性能足够好，不能存在时移性的不正常性能损耗，但不包括瞬时冲击造成的阻塞，必须通过常见压力测试
- 安全性足够，需要满足弹性伸缩，不能包含单进程内的独有状态，必须通过崩溃测试

## 使用手册
1. 配置环境变量式参数，因为程序需要这些参数，确认无误后，启动它
> 系统在首次启动时，会初始化数据库和超级权限账号(root:P@ssW0rd)，然后你应该用超级权限账号来建设组织架构

推荐的做法是:登录root->注册其他账号/添加组织,角色,职位,用户组,模块,路径等资源->然后赋予角色权限->将账号和资源绑定
特别的，我们通过角色权限来校验，而不是账号，所以只要账号绑定了角色就行。下面举个例子来说明绑定关系

- office - 组织
- group - 组
- position - 职位
- role - 角色
- user - 用户
- module - 模块

> office <-> position <-> role 组织下的职位的角色

> group <-> role 组的角色

> user <-> role 用户的角色

> user <-> group 用户的用户组 

> user <-> position 用户的职位

一些衍生关系这里就不细说了，例如职位是衍生自组织，详细请看docs文件下的说明

# 接口文档
接口文档放置在showdoc

文档扫描采用[content scanner](https://github.com/milkomeda-org/content-scanner)

扫描工具采用[interactive assistant](https://github.com/lauvinson/interactive-assistant)

# 路由转发

本项目包含转发你的业务请求到指定业务系统，在这之前会完成本项目的权限职责


