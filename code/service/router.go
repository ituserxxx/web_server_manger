package service

import (
	"fmt"
	"net/http"
)

func RegisterRoute() http.Handler {
	router := http.NewServeMux()
	handler := corsMiddleware(router)

	fs := http.FileServer(http.Dir("dist"))

	// 建立路由规则，将所有请求交给静态文件处理器处理
	router.Handle("/", fs)

	// GET请求处理
	router.HandleFunc("/hello", helloHandler)
	// POST请求处理

	router.HandleFunc("/user/Login", middleware(UserController.Login))
	router.HandleFunc("/user/Info", middleware(UserController.Info))

	router.HandleFunc("/myhost/Add", middleware(MyhostController.Add))
	router.HandleFunc("/myhost/Update", middleware(MyhostController.Update))
	router.HandleFunc("/myhost/Del", middleware(MyhostController.Del))
	router.HandleFunc("/myhost/GetList", middleware(MyhostController.GetList))

	router.HandleFunc("/myservice/Add", middleware(MyserviceController.Add))
	router.HandleFunc("/myservice/Update", middleware(MyserviceController.Update))
	router.HandleFunc("/myservice/Del", middleware(MyserviceController.Del))
	router.HandleFunc("/myservice/GetList", middleware(MyserviceController.GetList))
	router.HandleFunc("/notify/Get", middleware(NotifyController.Get))
	router.HandleFunc("/notify/Update", middleware(NotifyController.Update))
	return handler
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
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取路由参数
	name := r.URL.Query().Get("name")

	// 返回响应
	fmt.Fprintf(w, "Hello, %s!", name)
}
