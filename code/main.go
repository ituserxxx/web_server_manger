package main

import (
	"net/http"
	"server_manager/db"
)

func main() {
	db.StartDb()
	StartRoute()
	// 启动HTTP服务器
	if err := http.ListenAndServe("localhost:9528", nil); err != nil {
		panic(err)
	}
}
