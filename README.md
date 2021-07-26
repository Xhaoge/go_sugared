

项目目录结构：
go_sugared/
├── config
│   ├── config.go
│   └── dev_config.yaml
├── middleware
├── pkg
│   ├── msg
│   │   └── code.go
│   ├── loggig
│   │   ├── file.go
│   │   └── logging.go
│   └── util
│       ├── httpSend.go
│       └── utils.go
├── routers
│   ├── api
│   │   ├── common.go
│   │   ├── pic
│   │   │    ├── picModel.go
│   │   │    ├── picController.go
│   │   ├── room
│   │   │    ├── roomController.go
│   │   │    ├── roomModel.go
│   │   │    ├── storage.go
│   │   │    ├── transformer.go
│   │   └── user
│   │        ├── userController.go
│   │        ├── storage.go
│   │        └── userModel.go
│   └── router.go
├── runtime
├── iMissYou.txt
└── main.go




mian 入口
config：用于存储配置文件
middleware：应用中间件
models：应用数据库模型
pkg：第三方包
routers: 路由逻辑处理
runtime：应用运行时数据


/cmd
    目录中存储的都是当前项目中的可执行文件，该目录下的每一个子目录都应该包含我们希望有的可执行文件，如果我们的项目是一个 grpc 服务的话，
    可能在 /cmd/server/main.go 中就包含了启动服务进程的代码，编译后生成的可执行文件就是 server。我们不应该在 /cmd 目录中放置太多的代码，
    我们应该将公有代码放置到 /pkg 中并将私有代码放置到 /internal 中并在 /cmd 中引入这些包，保证 main 函数中的代码尽可能简单和少。

/api
    目录中存放的就是当前项目对外提供的各种不同类型的 API 接口定义文件了，其中可能包含类似 /api/protobuf-spec、/api/thrift-spec 或者 
    /api/http-spec 的目录，这些目录中包含了当前项目对外提供的和依赖的所有 API 文件：

/internal
    私有代码推荐放到/internal目录中，真正的项目代码应该写在 /internal/app 里，同时这些内部应用依赖的代码库应该在 /internal/pkg 子目录
    和 /pkg 中，没办法下图展示了一个使用 /internal 目录的项目结构：

/pkg 
    目录是 Go 语言项目中非常常见的目录，我们几乎能够在所有知名的开源项目（非框架）中找到它的身影，例如：
    prometheus 上报和存储指标的时序数据库
    istio 服务网格 2.0
    kubernetes 容器调度管理系统
    grafana 展示监控和指标的仪表盘
