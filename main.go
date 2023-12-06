package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ws", createWs)
	// router.RunTLS(":8888", "/home/faixin/IMdemo/chat.faixin.cn/fullchain.pem", "/home/faixin/IMdemo/chat.faixin.cn/privkey.pem")
	router.Run(":8888")
}

//交叉编译至linux平台
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server *.go
