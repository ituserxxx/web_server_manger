package main

import (
	"fmt"
	"net/http"
)

func InitRoute(router *http.ServeMux) {
	// GET请求处理
	router.HandleFunc("/hello", helloHandler)
	// POST请求处理

	router.HandleFunc("/dev-api/user/Login", middleware(UserController.Login))
	router.HandleFunc("/dev-api/user/Info", middleware(UserController.Info))

	router.HandleFunc("/dev-api/myhost/Add", middleware(MyhostController.Add))
	router.HandleFunc("/dev-api/myhost/Update", middleware(MyhostController.Update))
	router.HandleFunc("/dev-api/myhost/Del", middleware(MyhostController.Del))
	router.HandleFunc("/dev-api/myhost/GetList", middleware(MyhostController.GetList))

	router.HandleFunc("/dev-api/myservice/Add", middleware(MyserviceController.Add))
	router.HandleFunc("/dev-api/myservice/Update", middleware(MyserviceController.Update))
	router.HandleFunc("/dev-api/myservice/Del", middleware(MyserviceController.Del))
	router.HandleFunc("/dev-api/myservice/GetList", middleware(MyserviceController.GetList))
}

// 自定义中间件函数
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
