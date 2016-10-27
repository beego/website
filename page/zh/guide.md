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

`bee` 工具安装和配置成功后，可以在终端运行 `bee version` 测试。如果正确输出了版本号、GOPATH 等信息，你就可以使用 `bee` 工具创建第一个 beego 了。`bee` 工具更详细的文档可以阅读 [todo](#)。

# 创建 beego 项目

切换到 `GOPATH/src` 下启动终端（或命令行），执行命令：

```bash

bee new webapp

```

`bee` 工具会创建默认的 beego 项目在 `webapp` 目录。进入目录执行 `bee run`，就会编译和启动这个项目：

```bash

cd webapp # 记得要 cd 进 webapp 目录
bee run

```

访问 `http://localhost:8080` 就可以看到带有 beego 蜜蜂 logo 的默认欢迎页面。

来浏览一下 `webapp` 项目的结构：

    webapp
        |--- conf           # 配置文件目录
        |--- controllers    # MVC 中的控制器目录
        |--- models         # MVC 中的模型目录
        |--- routers        # 路由规则目录
        |--- static         # 静态内容目录
        |--- tests          # 集成测试目录
        |--- views          # MVC 中的视图的模板的目录
        |--- main.go        # 入口 go 文件

从入口文件 `main.go` 开始：

```go
package main

import (
	_ "webapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
} 
```

`main()` 函数是 go 程序的入口。`beego.Run()` 启动 beego 的 HTTP 服务。但是代码里只有启动 HTTP ，没有实际的处理逻辑。继续阅读需要看引入的 `webapp/routers` 包的内容。

##### 路由

`import` 引入了 `webapp/routers` 为 beego 注册路由规则。

打开 `webapp/routers/router.go` 可以看到 `beego` 最简单的路由：

```go
package routers

import (
	"webapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
} 
```

`beego.Router()` 是 beego 注册路由规则的方法。代码向规则 `/` 注册了控制器 `controllers.MainController` 作为访问 `/` 时的处理逻辑。这里只注册 URL `/` 的规则，没有注册别的 URL 请求。如果你访问 `http://localhost:8080/xyz` 会出现 404 页面。不同 URL 的规则需要你根据业务自行添加。**路由注册** 的详细内容可以阅读 [todo] 文档。 

##### 控制器

接下来就是看 `controller.MainController`，在 `webapp/controllers/default.go`：

```go
package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
} 
```

`MainController` 内嵌 `beego.Controller`，然后覆写 `beego.Controller.Get()` 方法，用来接受 HTTP GET 请求。`c.Data` 是控制器当前维持的数据，会用于渲染视图，输出 JSON 等。`c.TplName` 定义当前请求需要渲染的视图模板的文件名。`c.Data` 的数据会提供到 `index.tpl` 中参与渲染。

`beego.Controller` 是 beego 的基本控制器结构。定义 `Get()`、`Post()` 等方法接收不同方法的 HTTP 请求。可以阅读文档 [todo] 了解更多的控制器基本使用。

如果你需要处理 HTTP 的请求数据，阅读 【todo】 查看更多实例。如果你不需要渲染视图，而是输出 JSON，甚至下载文件，可以阅读 【todo】了解 beego 如何操作。

##### 视图

`c.TplName = "index.tpl"` 中的 `index.tpl` 是指 `webapp/views/index.tpl` 文件。文件很长，主要的内容在 `<body>` 标签：

```html
<body>
  <header>
    <h1 class="logo">Welcome to Beego</h1>
    <div class="description">
      Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>
</body>
```

可以看到， `c.Data["Website"]` 和 `c.Data["Email"]` 分别对应模板中的 `{{.Website}}` 和 `{{.Email}}`。 beego 使用标准库 `html/template` 来解析和渲染模板。你可以阅读 【todo】 了解 `html/template` 的使用方法。当然 beego 为渲染过程也添加了一些辅助函数，阅读 【todo】 去了解 beego 特有的模板辅助函数。

# 阅读更多

上面只是引导入门 beego 默认的项目内容。如果你需要更详细的了解 beego 的相关操作，请阅读完整的文档 【todo】 和 API 列表 【todo】 获取更多的信息。

如果你有使用的疑问，可以在 Github 发起 issue，或加入社区 【todo】 提问。

如果你有更好的想法或实现帮助 beego 提高设计和性能，欢迎在 `Github` 发起 Pull Request 帮助我们。

