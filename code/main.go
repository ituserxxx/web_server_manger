package main

import (
	"fmt"
	"net/http"
	"server_manager/db"
)

func main() {
	db.Init()
	router := http.NewServeMux()
	handler := corsMiddleware(router)

	InitRoute(router)
	NewCheckTask()
	fmt.Println(":8001")

	// 启动HTTP服务器
	http.ListenAndServe(":8001", handler)

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
			w.WriteHeader(http.StatusOK)
			return
		}

		// 继续执行下一个处理程序
		next.ServeHTTP(w, r)
	})
}
