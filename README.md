# go-blog
基于Gin、Vue两个框架开发的一个前后端分离的博客系统，数据库采用的是mysql。

在线示例 [www.renhj.cc](https://www.renhj.cc)

文档地址 [go-blog.renhj.cc](go-blog.renhj.cc)


## Getting Started

### Prerequisites

1. MySQL 服务器，参照此文档进行安装 [MySQL安装]()

2. SMTP 账号，可以使用qq邮箱或者163邮箱的smtp服务。  

3. 一台云主机  

4. 将MySQL账号及SMTP账号配置到server的配置文件中

### Installing

1. Server

```bash
# 进入server目录
cd server
# 下载go依赖
go mod download
# 启动main.go并指定配置文件位置
go run main.go -c configs/config.yaml
```

2. Web

```bash
# 进入web端目录
cd web
# 安装node的依赖
yarn install
# 运行前端服务
yarn serve
```


## Deployment

1. Server

1.1. 编译

```bash
go build 
```

1.2. 上传文件

1.3. 启动服务

2. web

2.1. 编译

2.2. 上传

2.3. 启动服务

> #### 3、Nginx
> 可以额外通过nginx做反向代理使用一个域名和端口来进行集群配置和隐藏后端ip，
> nginx的配置可以下载nginx.conf，修改其中的域名和HTTPS证书位置。

## Build With

* Go - 后端server的编译工具
* Webpack - 前端node的编辑工具

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

## Authors

* JackRay-C 

See also the list of contributors who participated in this project.

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details

## 八、示例及文档地址