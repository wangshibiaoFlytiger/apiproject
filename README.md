_**若本项目给您带来收获, 还请您动动小拇指,右上角给点个赞哈,万分感谢哈哈!!!**_

项目原始仓库地址: https://github.com/wangshibiaoFlytiger/apiproject  
欢迎关注官方微信公众号,及时收到一手技术资料!  
![微信公众号二维码](<https://gitee.com/wangshibiao/blog_picBed2/raw/master/images/qrcode_for_gh_66d82451d714_258(1).jpg> "微信公众号二维码")

# 1. 博客地址

https://www.sofineday.com

这 2 天刚开始搭建, 后面会逐步完善功能,性能和内容等, 望收藏, 有任何疑问请提交 issue, 谢谢!

_**欢迎加作者微信645102170或进群共同交流, 请扫下方二维码. 请备注 sofineday:smile:**_  
![](https://i.loli.net/2020/07/16/5X1HohKICW7knlL.jpg)

> 目标是归纳出平时经常用到的重要知识点, 作为日后的枕边工具书

# 2. 工程创建

mkdir ./apiproject

cd ./apiproject

go mod init apiproject

go build

# 3. 集成 gin web 框架

go get github.com/gin-gonic/gin

go get github.com/go-sql-driver/mysql

go get github.com/jinzhu/gorm

# 4. 项目规范

## 4.1 包名追加前缀

model 层: m\_

dao 层: d\_

service 层: s\_

controller 层: c\_

## 4.2 文件名或 struct 名称定义规范

每一层的文件名或类名以所属层为后缀, model 层例外

# 5. 测试代码

执行单元测试, 需要进入 test 目录, 然后执行 go test

# 6. live reload 自动编译并运行

go get -u github.com/cosmtrek/air

项目根目录创建 liveReload.conf 文件:

编辑配置项 bin,full_bin,exclude_file(将 rice-box.go 移除)

启动 liveReload 服务

```
 air -c ./liveReload.conf
```

# 7. 集成自动测试框架 goconvey

go get github.com/smartystreets/goconvey

# 8. 支持跨域

go get github.com/gin-contrib/cors

> 此外还需要如下显示指定 ugorji/go 的版本, 使用 > v1.1.2 版本,最新版本会出现包冲突错误:cannot load github.com/ugorji/go/codec: ambiguous import: found github.com/ugorji/go/codec in multiple modules
> go get github.com/ugorji/go@v1.1.2

# 9. 指定运行环境

通过指定命令行参数--profile 参数来加载对应的配置文件, 取值范围:dev,test,pro.

可通过执行帮助命令查看./apiproject help

# 10. 打包静态文件到可执行文件

go get github.com/GeertJohan/go.rice/...

执行 go build 之前, 执行如下命令, 将项目中所有引用 rice.MustFindBox 或 rice.FindBox 的所属文件路径, 作为--import-path 的参数

```
rice -v --import-path "./router" --import-path "./config" embed-go
```

**注意事项**

1. rice.MustFindBox 的路径参数不能为绝对路径, 且必须为字符串字面量, 不能为字符串变量
2. 最终的静态文件路径不能以./开头, 否则报错, 如错误写法 box.MustBytes("./config_dev.ini"), 正确写法 box.MustBytes("config_dev.ini")
3. 每次修改相应的静态文件, 记得随时执行 rice 命令,　否则会导致程序读取的文件内容不是最新的
4. 把 box 理解为目录即可

# 11. 压缩可执行文件

## 11.1 优化编译参数

go build 增加如下编译参数

```
   -ldflags "-s -w"
```

## 11.2 使用 upx 压缩工具进一步压缩

下载地址:https://github.com/upx/upx/releases

使用方法

```
upx --brute 可执行文件路径
```

# 12. git 分支约定

## 12.1 master 分支放通用功能, 目标是打造成 1 个通用框架

## 12.2 创建实际项目的操作

创建项目主分支

```
git checkout -b 项目名称_master
```

创建针对该项目的分支

```
git checkout 项目名称_master
git checkout -b 项目名称_branch_当前时间
```

合并项目分支到项目主分支

```
git checkout 项目名称_master
```

合并项目分支到项目主分支

合并通用框架的递增功能到项目分支(可以作为通用的功能一定要在 master 分支中开发)

# 13. 分布式 ID 生成库

go get github.com/bwmarrin/snowflake

# 14. kafka 消息队列

confluent-kafka-go 在后期使用过程中, kafka 服务端报未知异常, 故舍弃该客户端, 改用 sarama

# 15. 纯真 ip 归属地查询

## 15.1 安装

直接执行如下命令会提示安装失败

```
malformed module path "github.com/kayon/iploc/...": double dot
```

经实践, 顺序执行如下 2 个命令即可完成安装, 原因未知

```
go get -u github.com/kayon/iploc/
go get -u github.com/kayon/iploc/...
```

## 15.2 下载纯真数据库文件

iploc-fetch qqwry.gbk.dat

## 15.3 将 IP 库文件的编码转换为 UTF-8

iploc-conv -s qqwry.gbk.dat -d qqwry.utf8.dat

## 16. 请作者喝杯:coffee:

**若本项目给您带来收获，就请作者喝杯咖啡吧, 请备注 sofineday, 谢谢:smile:**
| 微信 | 支付宝 |
| ------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| ![微信](https://cdn.jsdelivr.net/gh/wangshibiaoFlytiger/blog_picBed1/images/微信.jpg) | ![支付宝](https://cdn.jsdelivr.net/gh/wangshibiaoFlytiger/blog_picBed1/images/支付宝.jpg) |
