# go-blog
基于Gin、Vue两个框架开发的一个前后端分离的博客系统，数据库采用的是mysql。

## 快速开始

### 准备工作

1. MySQL 服务器，参照此文档进行安装 [MySQL安装]()

2. SMTP 账号，可以使用qq邮箱或者163邮箱的smtp服务。  

3. 一台云主机  

4. 将MySQL账号及SMTP账号配置到server的配置文件中

### 安装

1. server端

```bash
# 进入server端目录
cd server
# 下载go依赖
go mod download
# 启动main.go并指定配置文件位置
go run main.go -c configs/config.yaml
```

2. web端

```bash
# 进入web端目录
cd web
# 安装node的依赖
yarn install
# 运行前端服务
yarn serve
```