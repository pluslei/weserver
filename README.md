## weserver

微信聊天室后台，系统采用beego框架,后台界面采用Amazeui,系统消息推送采用阿里云mqtt形式

VERSION = "V0.12.28"

## 其它相关
系统中采用了第三方包,但经过修改，源码被修改到[wechat](git@git.haoyue.me:iuan/wechat.git)

## 获取安装
go get git@git.haoyue.me:hlive/weserver.git

执行以下命令
首先你应该先有beego 环境。
 
1.然后把源码放在gopath的src目录下。

2.利用go run 运行程序，或bee run (若无法执行bee run ,请下载 go get [github.com/beego/bee](https://github.com/beego/bee))

## 接口文档

具体的接口参考 [Wiki](http://git.haoyue.me:8080/hlive/weserver/wiki)

## 系统数据库初始化
将当前目录切换到`weserver`中
1. win系统中 `weserver.exe db`
2. linux系统中 `./weserver db`

