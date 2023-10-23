package main

import (
	"fmt"
	"net/http"
	"server_manager/db"
	"server_manager/service"
)

func main() {
	db.Init()
	service.InitCheckTask()
	service.InitNotifyPush()
	fmt.Println("http://ip:8001")

	// 启动HTTP服务器
	http.ListenAndServe(":8001", service.RegisterRoute())

}
