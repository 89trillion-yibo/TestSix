package ws

import (
	"awesomeProject/testsix/internal/model"
	"awesomeProject/testsix/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
)

var Conntion *websocket.Conn

//创建连接
func CareatCon(username string,server string) {
	header := http.Header{}
	header.Add("name",username)
	conn, _, err := websocket.DefaultDialer.Dial(server, header)
	if err!=nil{
		fmt.Println(err)
		model.Infomation.SetText(model.Infomation.Text+"\n"+"server message:Sorry,error in connecting")
		return
	}
	Conntion = conn
	model.Status.SetText("ok")
	go Read(conn)
}

//发送消息给服务器
func Write(username string,message string)  {
	meg := protobuf.Message{
		MessageType: "talk",
		MessageText: message,
		User: username,
	}
	marshal, _ := proto.Marshal(&meg)
	err := Conntion.WriteMessage(websocket.TextMessage, marshal)
	fmt.Println("发送消息",meg)
	if err!=nil{
		fmt.Println(err)
	}
}

//与服务器断开连接
func DisCon(username string) {
	meg := protobuf.Message{
		User: username,
		MessageType: "exit",
	}
	marshal, _ := proto.Marshal(&meg)
	err := Conntion.WriteMessage(websocket.TextMessage, marshal)
	fmt.Println("断开连接")
	if err!=nil{
		fmt.Println(err)
	}
	err = Conntion.Close()
	if err!=nil{
		fmt.Println(err)
	}
}

//请求在线用户列表
func RequestUserlist() {
	meg := protobuf.Message{
		MessageType: "userlist",
	}
	marshal, _ := proto.Marshal(&meg)
	err := Conntion.WriteMessage(websocket.TextMessage, marshal)
	fmt.Println("请求用户列表")
	if err!=nil{
		fmt.Println(err)
	}
}

//监听服务端发来的消息
func Read(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err!=nil{
			conn.Close()
			return
		}
		meg := protobuf.Message{}
		err = proto.Unmarshal(p, &meg)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("meg接受消息",meg)
		if len(meg.MessageText) != 0 && meg.MessageType == "talk" {
			model.Infomation.SetText(model.Infomation.Text+"\n"+meg.MessageText)
			//发起请求用户列表
			RequestUserlist()
		}
		if meg.MessageType == "userlist"{
			fmt.Println("进入判断")
			var list string
			for _,v := range meg.UserList{
				list += v + "\n"
			}
			model.UserList.SetText(list)
		}
	}
}