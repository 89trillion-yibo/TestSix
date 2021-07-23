package view

import (
	"awesomeProject/testsix/internal/ctrl"
	"awesomeProject/testsix/internal/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func InitView() {
	myApp := app.New()
	myWin := myApp.NewWindow("Entry")
	myWin.Resize(fyne.NewSize(800,700))

	//name输入框
	model.UserName = widget.NewEntry()
	model.UserName.SetPlaceHolder("input name")

	//status状态
	status := widget.NewLabel("status:")
	model.Status = widget.NewEntry()
	model.Status.SetText("no")
	//服务器地址输入框
	model.Server = widget.NewEntry()
	model.Server.SetPlaceHolder("input server")
	//连接按钮
	conButton := widget.NewButton("connection", conButton)

	//断开连接按钮
	disconButton := widget.NewButton("Disconnect", disButton)
	conBox := container.NewHBox(widget.NewLabel("server"), model.Server, layout.NewSpacer(), conButton, disconButton)
	line := canvas.NewLine(color.Black)
	nameBox := container.NewHBox(widget.NewLabel("Name"), model.UserName,layout.NewSpacer(),status,model.Status)
	//topbox
	boxtop := container.NewVBox(nameBox,conBox,line)

	//在线用户列表
	model.UserList = widget.NewLabel("")
	left := widget.NewCard("userlist", "", container.NewScroll(model.UserList))

	//显示信息框
	model.Infomation = widget.NewLabel("")
	Central := widget.NewCard("message", "", container.NewScroll(model.Infomation))

	//输入消息框
	model.Message = widget.NewMultiLineEntry()
	bottom := widget.NewCard("", "", model.Message)
	button := widget.NewButton("send",sendButton)
	boxbuttom := container.NewVBox(bottom, button)
	//边界布局
	border := container.NewBorder(boxtop, boxbuttom, left, nil,Central)
	myWin.SetContent(border)
	myWin.ShowAndRun()
}


//连接按钮调用
func conButton()  {
	ctrl.Connect(model.UserName.Text,model.Server.Text)
}

//发送消息按钮调用
func sendButton() {
	ctrl.WriteMessage(model.UserName.Text,model.Message.Text)
	model.Message.SetText("")
	model.Message.Refresh()
}

//断开连接调用
func disButton() {
	ctrl.DisConnect(model.UserName.Text)
}