#  Cli Document

## Global Flags

全局参数，-c指定配置文件的位置，不指定的话，默认在当前目录寻找`config/config.yaml`和`$HOME/.config.yaml`下寻找配置文件。

```bash
go-blog -c config.yaml
# or 
go-blog --config config.yaml
```


## init

初始化目录，-d指定初始化目录的位置。会在指定目录创建配置文件。

```bash
go-blog init -d /opt/go-blog
```


## 