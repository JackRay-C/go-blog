# Cli Document

## init

初始化目录，-d指定初始化目录的位置。会在指定目录创建目录及配置文件。

```bash
go-blog init -d /opt/go-blog
```

## start

```bash
go-blog server start -c /opt/go-blog/go-blog.yaml
```

## stop

```bash
go-blog server stop
```

## restart

```bash
go-blog server restart
```

## ui

[blog ui](https://github.com/JackRay-C/go-blog-ui)

## project directory

```bash
.  
├── cmd           # 命令实现
├── conf          # 配置文件
├── docs          # 文档
├── internal      # 内部包
├── logs          # 日志
├── pkg         
  ├── api         # 接口
  ├── global      # 全局变量
  ├── middleware  # 中间件
  ├── model       # 模型
  ├── resources   # 资源
  ├── routes      # 路由
  ├── services    # 服务
  ├── utils       # 工具
```

## install directory

```bash
.  
├── bin         # 可执行文件
├── conf        # 配置文件
├── logs        # 日志
├── public      # 前端静态文件存储
└── uploads     # 上传文件
```