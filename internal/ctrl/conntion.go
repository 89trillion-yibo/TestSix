package ctrl

import (
	"awesomeProject/testsix/internal/model"
	"awesomeProject/testsix/internal/ws"
)
//创建连接
func Connect(username string,server string) {
	if len(username) == 0 || len(server) == 0 {
		model.Infomation.SetText(model.Infomation.Text+"\n"+"your name or server is null")
		return
	}
	if model.Status.Text == "ok"{
		model.Infomation.SetText(model.Infomation.Text+"\n"+"Sorry, you cannot repeat the connection")
		return
	}
	ws.CareatCon(username,server)
}

//发送消息
func WriteMessage(username string,message string) {
	if len(username) == 0 || len(message) == 0 {
		model.Infomation.SetText(model.Infomation.Text+"\n"+"your name or message is null")
		return
	}
	ws.Write(username,message)
}

//断开连接
func DisConnect(username string) {
	if model.Status.Text != "ok"{
		model.Infomation.SetText(model.Infomation.Text+"\n"+"you are not connected yet")
		return
	}
	ws.DisCon(username)
	model.Status.SetText("no")
}

