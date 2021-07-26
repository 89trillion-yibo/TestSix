# 技术文档

### 1.整体框架

基于fyne开发聊天界面，与第五题的websocket服务端连接完善聊天室，使用protobuf数据格式与服务端通信，在窗口中输入用户名与服务端地址，请求连接，连接成功后进入聊天室，下方输入聊天内容，点击"send"发送，左边维护一个在线用户列表，退出聊天室，点击"Disconnect"断开连接



### 2.目录结构

```
├── app
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── conntion.go
│   ├── model
│   │   └── client.go
│   └── ws
│       └── operating.go
├── protobuf
│   ├── message.pb.go
│   └── message.proto
└── view
    └── viewWindow.go

```



### 3.代码逻辑分层

| 层       | 文件夹                    | 主要职责                                | 调用关系                     | 其它说明     |
| -------- | ------------------------- | --------------------------------------- | ---------------------------- | ------------ |
| ws       | internal/ws/operating.go  | 创建websocket连接并发送，监听服务端消息 | 调用路由层                   | 不可同层调用 |
| ctrl层   | internal/ctrl/conntion.go | 参数验证                                | 被view调用                   | 不可同层调用 |
| view层   | view/viewWindow.go        | 启动客户端界面                          | 调用ctrl                     | 不可同层调用 |
| model层  | internal/model/client.go  | 数据模型，定义界面控件                  | 被其他层调用                 | 不可同层调用 |
| protobuf | protobuf                  | 存放protobuf相关文件                    | 被handler、service、view调用 | 不可同层调用 |



### 4.存储设计

消息数据

| 内容         | 数据类形 | Key         |
| ------------ | -------- | ----------- |
| 消息内容     | string   | MessageText |
| 消息类型     | string   | MessageType |
| 用户名       | string   | User        |
| 在线用户列表 | []string | UserList    |



### 5.UI界面

初始界面


<img width="800" alt="主界面" src="https://user-images.githubusercontent.com/87186547/126996512-63f04cc6-cb8a-4d5a-ac59-0067ed650db1.png">


用户连接

<img width="800" alt="用户连接" src="https://user-images.githubusercontent.com/87186547/126996550-7f858356-5b60-4228-8aa9-9faf6c50187d.png">


用户发消息

<img width="800" alt="用户发消息" src="https://user-images.githubusercontent.com/87186547/126996574-a42b27bd-6c23-4f3b-ba84-e79f811d81a6.png">


重复连接

<img width="800" alt="重复连接" src="https://user-images.githubusercontent.com/87186547/126996593-fa3d75e6-26a6-40b3-a796-902fba9566a2.png">


断开连接

<img width="800" alt="断开连接" src="https://user-images.githubusercontent.com/87186547/126996621-0c612a01-1d5a-4fea-a9ef-6126d26c9285.png">


未连接发消息

<img width="799" alt="未连接发消息" src="https://user-images.githubusercontent.com/87186547/126996657-bbfeb228-a93b-4497-b802-72934fab65ae.png">



### 6.第三方库

### fyne

```
https://github.com/fyne-io/fyne
```

### websocket

```
github.com/gorilla/websocket
```



### 7.流程图

![未命名文件 (8)](https://user-images.githubusercontent.com/87186547/126996721-5235c44d-ca10-4389-aced-c6a1f86f5433.jpg)

