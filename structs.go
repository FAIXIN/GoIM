package main

import (
	"github.com/gorilla/websocket"
)

// 连接成功结构体
type connectedDataS struct {
	Event            string          `json:"event"`
	HistoryClientNum int             `json:"historyClientNum"`
	OnlineClientNum  int             `json:"onlineClientNum"`
	Client           *websocket.Conn `json:"client"`
}

// 广播消息结构体
type messageDataS struct {
	Event    string      `json:"event"`
	Ava      string      `json:"ava"`
	Message  string      `json:"message"`
	UserName string      `json:"userName"`
	Time     string      `json:"time"`
	Type     string      `json:"type"`
	Result   bool        `json:"result"`
	Info     string      `json:"info"`
	UserList interface{} `json:"userList"`
}

// 来自客户端的JSON
type reciveMsg struct {
	OldName  string      `json:"oldName"`
	Ava      string      `json:"ava"`
	Receiver string      `json:"receiver"`
	Event    string      `json:"event"`
	Name     string      `json:"name"`
	Msg      string      `json:"msg"`
	Time     string      `json:"time"`
	Type     string      `json:"type"`
	Pwd      string      `json:"pwd"`
	UserList interface{} `json:"userList"`
}

// 用户名头像列表
type userListAvatar struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// 存储数据
var (
	// 初始化历史客户端数为 0
	historyClientNum = 0

	//历史消息池
	historyMessages []reciveMsg

	//历史私聊消息池
	privateHisMessages []reciveMsg

	// //在线用户列表
	// users []string

	userList []userListAvatar

	//用户名到客户端连接的映射
	reClients = make(map[string]*websocket.Conn)

	//客户端连接到用户名的映射
	clients = make(map[*websocket.Conn]string)

	//缓存图片
	imgs = make(map[string]string)

	//用户名到密码的映射
	pwds = make(map[string]string)
)
