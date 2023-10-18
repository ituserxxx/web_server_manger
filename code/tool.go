package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var Tool = &tool{}

type tool struct {
}

func (to *tool) ChecIpPortIsOk(ipPort string) string {
	// 设置连接超时时间
	conn, err := net.DialTimeout("tcp", ipPort, time.Duration(10)*time.Second)
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	return ""
}
func (to *tool) CheckUrlIsOk(urlStr string) string {
	// 发送 HTTP GET 请求
	response, err := http.Get(urlStr)
	if err != nil {
		fmt.Println("请求出错:", err)
		return err.Error()
	}
	defer response.Body.Close()

	// 判断响应状态码是否为 200 OK
	if response.StatusCode == http.StatusOK {
		return ""
	}
	return fmt.Sprintf("错误状态码：%d", response.StatusCode)

}
func (to *tool) Krand() string {
	kind := 1
	size := 5
	// 随机字符串
	//func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {

		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
	//}
}
