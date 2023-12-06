package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// 广播消息
func broadcastMsg(mt int, message []byte) {
	for client := range clients {
		client.WriteMessage(mt, message) // 将消息发送给所有客户端
	}
}

// 私发消息
func privateMsg(message interface{}, targetClient *websocket.Conn) {
	if targetClient == nil {
		// 如果目标客户端不存在
		fmt.Println("Target client not found")
		return
	}

	dataJson, _ := json.Marshal(message)
	err := targetClient.WriteMessage(1, dataJson)

	if err != nil {
		// 处理发送私发消息的错误
		fmt.Println("Failed to send private message:", err)
	}
}

// 调用私发消息发送历史消息到指定客户端
func sendHistoryMessages(reciveMsg reciveMsg, client *websocket.Conn) {

	var data = messageDataS{
		Event:    "hismessage",
		Ava:      reciveMsg.Ava,
		Message:  reciveMsg.Msg, // 将消息内容转为字符串
		UserName: reciveMsg.Name,
		Time:     reciveMsg.Time,
		Type:     reciveMsg.Type,
	}
	// dataJson, _ := json.Marshal(data) // 转换消息数据为 JSON
	privateMsg(data, client) // 私发 JSON 消息
}

// 发送消息
func sendMsg(reciveMsg reciveMsg, event string) {
	var data = messageDataS{
		Event:    event,
		Ava:      reciveMsg.Ava,
		Message:  reciveMsg.Msg,
		UserName: reciveMsg.Name,
		Time:     reciveMsg.Time,
		Type:     reciveMsg.Type,
		UserList: reciveMsg.UserList,
	}
	dataJson, _ := json.Marshal(data)
	broadcastMsg(1, dataJson)
}

// sendprivate
func sendPrivateMsg(reciveMsg reciveMsg, event string, client *websocket.Conn) {
	var data = messageDataS{
		Event:    event,
		Ava:      reciveMsg.Ava,
		Message:  reciveMsg.Msg,
		UserName: reciveMsg.Name,
		Time:     reciveMsg.Time,
		Type:     reciveMsg.Type,
		UserList: reciveMsg.UserList,
	}
	dataJson, _ := json.Marshal(data)
	client.WriteMessage(1, dataJson)
}
