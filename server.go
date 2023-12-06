package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// UpGrader用来将http协议升级成WebSocket协议
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 创建客户端实例
func createWs(c *gin.Context) {
	// 创建 WebSocket 连接客户端
	client, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	//用于处理客户端发来的JSON
	var reciveMsg reciveMsg
	for {
		// 监听接受信息
		_, message, err := client.ReadMessage() // 读取客户端发送的消息
		if err == nil {
			//将客户端发来的JSON转化为reciveMsg对象
			json.Unmarshal(message, &reciveMsg)

			switch reciveMsg.Event {
			case "message":
				//如果消息类型是图片的话，计算其hash值，并将hash和base64值存在imgs map中
				if reciveMsg.Type == "img" {
					hash := hashImg(reciveMsg.Msg)
					imgs[hash] = reciveMsg.Msg
					fmt.Println("====用户 ", reciveMsg.Name, reciveMsg.Time, " 发送图片,图片ID:", hash)
					reciveMsg.Msg = hash

					//判断是否是私聊消息，分类存入消息池
					if reciveMsg.Receiver != "" {
						privateHisMessages = append(privateHisMessages, reciveMsg)
					} else {
						historyMessages = append(historyMessages, reciveMsg)
					}

					reciveMsg.Msg = imgs[hash]
				}

				if reciveMsg.Type == "text" {
					//判断是否是私聊消息，分类存入消息池
					if reciveMsg.Receiver != "" {
						privateHisMessages = append(privateHisMessages, reciveMsg)
					} else {
						historyMessages = append(historyMessages, reciveMsg)
					}

					fmt.Println("====用户 ", reciveMsg.Name, reciveMsg.Time, " 发送文本信息:", reciveMsg.Msg)
					//backdoors
					if reciveMsg.Msg == "/clear" {
						historyMessages = historyMessages[:0]
					}

				}

				//私聊判断
				if reciveMsg.Receiver != "" {
					sendPrivateMsg(reciveMsg, "message", reClients[reciveMsg.Receiver])
				} else {
					sendMsg(reciveMsg, "message")
				}
			case "userNum":
				saveClient(client, reciveMsg.Name)
				// appendUsers(&users, reciveMsg.Name)
				user := userListAvatar{
					Name:   reciveMsg.Name,
					Avatar: reciveMsg.Ava,
				}
				appendUsersAvatar(&userList, user)

				reciveMsg.UserList = userList

				sendMsg(reciveMsg, "updateUserList")

				fmt.Println("在线用户表: ", reciveMsg.UserList)

			case "login":
				data := login(reciveMsg.Name, reciveMsg.Pwd)
				privateMsg(data, client)
			default:
				fmt.Println("无效的Event: ", reciveMsg.Event)
			}

		} else {
			fmt.Println("client.ReadMessage err:", err)
			break
		}
	}

	defer func() {
		//移除用户切片元素
		removeUsersAvatar(&userList, reciveMsg.Name)
		reciveMsg.UserList = userList
		sendMsg(reciveMsg, "updateUserList")
		// 关闭客户端连接
		client.Close()
		// 删除客户端并广播离开消息
		delClient(client, reciveMsg.Name)
	}()
}
