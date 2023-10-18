package main

import (
	"fmt"
	"net/http"
	"server_manager/db"
)

func main() {
	db.StartDb()
	router := http.NewServeMux()
	handler := corsMiddleware(router)

	StartRoute(router)
	fmt.Println("172.16.20.34:8001")
	// 启动HTTP服务器
	if err := http.ListenAndServe("172.16.20.34:8001", handler); err != nil {
		panic(err)
	}
}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域的域名列表
		// 设置允许的请求头部
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-token")

		// 对于预检请求（OPTIONS），直接返回成功
		if r.Method == "OPTIONS" {
			return
		}

		// 继续执行下一个处理程序
		next.ServeHTTP(w, r)
	})
}

// 检查字符串是否在字符串切片中
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
