package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ws", createWs)
	router.Run(":8888")
}

//交叉编译至linux平台
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server *.go
