package service

import (
	"fmt"
	"server_manager/db"
	"strconv"
	"sync"
	"time"
)

var CheckTask *checkTask

type checkTask struct {
	task  map[string][]*time.Ticker
	mutex sync.Mutex
}

func InitCheckTask() {
	CheckTask = &checkTask{
		task: make(map[string][]*time.Ticker),
	}
	CheckTask.init()
}

func (ct *checkTask) init() {
	l, err := db.MyserviceDb.QueryRow(func(s db.MyserviceT) bool {
		return true
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range l {
		if t.IsOpenCheck {
			ct.Add(t.ID, t)
		}
	}
}
func (ct *checkTask) Add(serviceId string, service db.MyserviceT) {
	ct.CanenlByTaskId(serviceId)
	ct.mutex.Lock()
	defer ct.mutex.Unlock()

	var ti = make([]*time.Ticker, 0)
	for _, item := range service.CheckList {
		sj, _ := strconv.Atoi(item.Time)
		if sj > 5 && item.Value != "" {
			tk := time.NewTicker(time.Duration(sj) * time.Second)
			go func(it db.CheckT) {
				for range tk.C {
					var state bool
					switch it.Type {
					case "tcp":
						if Tool.CheckTcp(it.Value) == "" {
							state = true
						}
					case "http":
						if Tool.CheckHttp(it.Value) == "" {
							state = true
						}
					default:
						state = false
					}
					//fmt.Printf("\n --check-- ing ----value-%v state-%v", it.Value, it.State)
					err := db.MyserviceDb.Update(
						func(s db.MyserviceT) db.MyserviceT {
							for i, v2 := range s.CheckList {
								if v2.ID == it.ID {
									s.CheckList[i].State = state
								}
							}
							return s
						},
						func(s db.MyserviceT) bool {
							if s.ID == serviceId {
								return true
							}
							return false
						})
					if err != nil {
						fmt.Printf("\nerr := db.MyserviceDb.Update is %v", err.Error())
					}

					// 推送
					var contentTmp = "靓仔服务异常啦~ \n项目名称：%s\n检测项：%s\n检测结果：%v\n检测时间：%s\n"
					msg := fmt.Sprintf(contentTmp, service.Name, it.Value, state, time.Now().Format("2006-01-02 15:04:05"))
					if state == false {
						NotifyPush.ChanList <- msg
					}
				}
			}(item)
			ti = append(ti, tk)
		}
	}
	ct.task[serviceId] = ti
}
func (ct *checkTask) CanenlByTaskId(taskId string) {
	ct.mutex.Lock()
	defer ct.mutex.Unlock()
	if tl, ok := ct.task[taskId]; ok {
		for _, timer := range tl {
			timer.Stop()
		}
		delete(ct.task, taskId)
	}
}
