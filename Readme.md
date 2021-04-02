### 一：项目介绍

| 版本   | 更新日期   | 更新内容 |
| :----- | :--------- | :------- |
| v0.0.0 | 2021.03.01 | 项目开始 |

#### 1. 项目说明

实现 C/S 客户端服务端交互的后端程序

```


后端：

```

### 二：项目启动

##### 1. 后端

初始化数据库

```
$ yum install mysql

$ mysql

> create database server_client DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
> create user 'odemo'@'%' identified by 'xxxxx';
> grant all privileges on odemo.*  to 'odemo'@'%';
```

运行项目

```

$ git clone https://github.com/Ohimma/server-client.git
$ cd server-client/tcp/server
$ go mod init tcp-server

$ go run main.go


```

### 三：项目结构

### 四：其他
