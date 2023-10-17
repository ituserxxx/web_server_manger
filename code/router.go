package main

import (
	"fmt"
	"net/http"
)

func StartRoute() {
	// GET请求处理
	http.HandleFunc("/hello", helloHandler)
	// POST请求处理
	http.HandleFunc("/dev-api/myhost/add", middleware(MyhostController.Add))
	http.HandleFunc("/dev-api/myhost/getList", middleware(MyhostController.GetList))

}

// 自定义中间件函数
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 中间件逻辑处理
		fmt.Println("Executing middleware")

		// 检查是否有授权头
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// 调用下一个处理函数
		next(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取路由参数
	name := r.URL.Query().Get("name")

	// 返回响应
	fmt.Fprintf(w, "Hello, %s!", name)
}
