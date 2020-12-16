# 记录目前遇到的困难和解决方式


## 版本同步
选择使用github.

## 开发环境选择
因为最开始后端选手使用的开发平台不同(MacBook Pro,ArchLinux,Windows)，为了避免出现你的代码在我这里不能跑的问题选择了容器化部署(见[开发环境](docker-compose-dev.yml)).
最后部署到服务器上也准备使用docker.

## docker容器通信
数据库PostgreSQL(端口映射5203:5203)选择使用docker容器部署,而go(端口映射8080:8080)开发环境(见[go](server/Dockerfile))也选择了服务器部署.于是两容器间的通信自然而然成了一大困扰.
经过搜索最近决定使用宿主机IP实现容器间通信.
### 宿主机IP
在Linux下安装Docker的时候，会在宿主机安装一个虚拟网卡docker0，我们可以使用宿主机在docker0上的IP地址来代替localhost。
首先，使用如下命令查询宿主机IP地址：
```
$ ip addr show docker0
6: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:fc:06:86:fc brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:fcff:fe06:86fc/64 scope link 
       valid_lft forever preferred_lft forever

```
可以发现宿主机的IP是***172.17.0.1***，那么在[server.go](server/server.go)中将与PostgreSQL连接时的URI改为`"PostgreSQL://" + "172.17.0.1" + ":5203",`即可.

但是，在Windows和macOS平台下并没有docker0虚拟网卡，这时候可以使用***host.docker.internal***这个特殊的DNS名称来解析宿主机IP。

## URL约定
localhost:8080/login            登录
localhost:8080/test             考试
localhost:8080/practice         测试
localhost:8080/create           DIY
localhost:8080/create/test      DIY+考试
localhost:8080/create/practice  DIY+测试
localhost:8080/commit           提交

## JSON约定
```
LevelChoose
{
    "op":[]         //运算符
    "nums":"3"      //运算数个数
    "maxnum":"20"   //运算数的最大值
    "qnums":"10"    //expression个数
    "anssize":"100" //答案的最大值
}

commit
{
    "owner":"name"
        "expressions":[
        {
            "exp":"1+2+3"
        },
        {
            "exp":"2+3+4"
        }
    ]
}

expressions
{
    "expressions":{
        {
            "Expression":"11+11",
            "Answer":"22"
        },
        {
            "Expression":"1+2+3",
            "Answer":"6"
        }
    }
}
```
