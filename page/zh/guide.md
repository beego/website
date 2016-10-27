你好！这篇 **快速入门** 将带你学习和使用创建基于 beego 的开发项目，熟悉 beego 的相关概念和工具操作，以及引导深入阅读 beego 的相关文档。

beego 是基于 Go 语言实现 MVC 概念的 Web 框架，提供丰富的 API 和组件库，帮助编写高生产力和高性能的 Web 应用。在使用 beego 之前，你需要先了解：

- Go 语言的基本语法，相关命令行的使用如 `go install`,`go build` 等
- 对 HTTP 的标准库 `net/http` 有一定了解。beego 是基于 `net/http` 开发的。
- 对 MVC 的概念有一定了解

# 安装

安装 beego 之前请先确认 Go 语言已经安装成功，设置好 `GOPATH`，**并将 `GOPATH/bin` 添加到系统 `PATH` 中**。 beego 需要 `Go 1.4+` 支持。 使用 `go get` 直接安装 beego：

```bash

# -v 打印更多下载信息
# -u 拉取最近代码，即使本地已经存在

go get -v -u github.com/astaxie/beego

```

beego 配合 `bee` 工具，可以更快速的创建和开发 beego 应用。安装 `bee` 工具可以直接从 [Github](#) 下载编译好的二进制文件，加入到系统 `PATH` 使用。或者自己编译安装：

```bash

go get -v -u github.com/beego/bee

# 下载并编译 bee 。二进制文件在 GOPATH/bin 目录

```