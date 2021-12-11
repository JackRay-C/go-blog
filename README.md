# go-blog
基于Gin、Vue两个框架开发的一个前后端分离的博客系统，数据库采用的是mysql。

* 在线示例 [www.renhj.cc](https://www.renhj.cc)
* 文档地址 [go-blog.renhj.cc](go-blog.renhj.cc)


## 一、快速开始

### 1.1、准备工作

1. MySQL 服务器，参照此文档进行安装 [MySQL安装]()

2. SMTP 账号，可以使用qq邮箱或者163邮箱的smtp服务。  

3. 一台云主机  

4. 将MySQL账号及SMTP账号配置到server的配置文件中

### 1.2、安装

#### 1.2.1、server端

```bash
# 进入server端目录
cd server
# 下载go依赖
go mod download
# 启动main.go并指定配置文件位置
go run main.go -c configs/config.yaml
```

#### 1.2.2、web端

```bash
# 进入web端目录
cd web
# 安装node的依赖
yarn install
# 运行前端服务
yarn serve
```


## 二、部署

#### 2.1、server端

1. 编译

```bash
go build 
```

2. 上传文件

3. 启动服务

#### 2.2、web端

1. 编译

2. 上传

3. 启动服务

> #### 2.3、Nginx
> 可以额外通过nginx做反向代理使用一个域名和端口来进行集群配置和隐藏后端ip，
> nginx的配置可以下载nginx.conf，修改其中的域名和HTTPS证书位置。

## 三、编译工具

* Go - 后端server的编译工具
* Webpack - 前端node的编辑工具

## 四、贡献指南

## 五、版本管理

## 六、作者

## 七、许可证

## 八、示例及文档地址