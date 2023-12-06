package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// 存储当前客户端
func saveClient(client *websocket.Conn, name string) {
	clients[client] = name   // 将当前客户端添加到映射中
	reClients[name] = client //

	historyClientNum++ // 增加历史客户端数

	// 广播连接成功消息
	var data = connectedDataS{
		Event:            "connected",
		HistoryClientNum: historyClientNum,
		OnlineClientNum:  len(clients), // 在线客户端数为映射的长度
		Client:           client,       //连接客户端地址
	}

	var dataJson, _ = json.Marshal(data) // 转换消息数据为 JSON
	broadcastMsg(1, dataJson)            // 广播 JSON 消息

	//遍历并且向当前客户端发送历史消息
	for _, msg := range historyMessages {
		if msg.Type == "img" {
			// 通过图片hash值取出图片值中的值
			if msg.Type == "img" {
				msg.Msg = imgs[msg.Msg]
			}
		}
		sendHistoryMessages(msg, client)
	}
}

// 删除当前客户端
func delClient(client *websocket.Conn, name string) {

	// 从映射中删除客户端
	delete(clients, client)
	delete(reClients, name)

	// 广播客户端离开消息
	var data = connectedDataS{
		Event:           "connected",
		OnlineClientNum: len(clients),
	}
	var dataJson, _ = json.Marshal(data)
	broadcastMsg(websocket.TextMessage, dataJson)
}
