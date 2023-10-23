package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"server_manager/db"
	"strings"
	"time"
)

var NotifyPush *notifyPush

func InitNotifyPush() {
	NotifyPush = &notifyPush{
		ChanList: make(chan string, 5),
	}
	go NotifyPush.startListen()
}

type notifyPush struct {
	ChanList chan string
}

// 企业微信通知
func (ntp *notifyPush) startListen() {
	for {
		select {
		case msg := <-ntp.ChanList:
			l, err := db.NotifyDb.QueryRow(func(s db.NotifyT) bool {
				return true
			})
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			if len(l) == 1 {
				d := l[0]
				if d.QywxIsOpen && d.QywxWebhook != "" {
					ntp.sendQiYeWeiXinPush(d.QywxWebhook, msg)
				}
				if d.DdIsOpen && d.DdWebhook != "" {
					ntp.sendDingDingPush(d.DdWebhook, msg)
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
func (ntp *notifyPush) sendQiYeWeiXinPush(url, content string) {
	msgStr := `{
	"msgtype": "markdown",
	"markdown": {
		"content": "%s" }
	}`
	msgStr = fmt.Sprintf(msgStr, content)
	data := []byte(msgStr)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() { _ = res.Body.Close() }()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("push Qywx result ", string(body))
}
func (ntp *notifyPush) sendDingDingPush(url, content string) {
	msgStr := `{
	"msgtype": "markdown", 
	"markdown": {
		"title": "服务监控",
		"text":"%s" }}`

	ContentSlice := strings.Split(content, "\n")
	ContentStr := strings.Join(ContentSlice, "\n\n")
	msgStr = fmt.Sprintf(msgStr, ContentStr)

	data := []byte(msgStr)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = res.Body.Close() }()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("push DingDing result ", string(body))
}
