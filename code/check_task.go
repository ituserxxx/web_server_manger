package main

import (
	"fmt"
	"server_manager/db"
	"strconv"
	"sync"
	"time"
)

var CheckTask *check_test

type check_test struct {
	task  map[string][]*time.Ticker
	mutex sync.Mutex
}

func NewCheckTask() {
	CheckTask = &check_test{
		task: make(map[string][]*time.Ticker),
	}
	CheckTask.init()
}

func (ct *check_test) init() {
	l, err := db.MyserviceDb.QueryRow(func(s db.MyserviceT) bool {
		return true
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range l {
		if t.IsOpenCheck {
			ct.Add(t.ID, t.CheckList)
		}
	}
}
func (ct *check_test) Add(taskId string, items []db.CheckT) {
	ct.CanenlByTaskId(taskId)
	ct.mutex.Lock()
	defer ct.mutex.Unlock()

	var ti = make([]*time.Ticker, 0)
	for _, item := range items {
		sj, _ := strconv.Atoi(item.Time)
		if sj > 5 && item.Value != "" {
			tk := time.NewTicker(time.Duration(sj) * time.Second)
			go func() {
				for range tk.C {
					fmt.Printf("\n --check-- value-%v", item.Value)
					//var state bool
					//switch item.Type {
					//case "tcp":
					//	if Tool.CheckTcp(item.Value) != "" {
					//		state = true
					//	}
					//case "http":
					//	if Tool.CheckHttp(item.Value) != "" {
					//		state = true
					//	}
					//default:
					//	state = false
					//}
					//fmt.Printf("\n --check-- state-%v", state)
					//db.MyserviceDb.Update(
					//	func(s db.MyserviceT) db.MyserviceT {
					//		for i, v2 := range s.CheckList {
					//			if v2.ID == item.ID {
					//				s.CheckList[i].State = state
					//			}
					//		}
					//		return s
					//	},
					//	func(s db.MyserviceT) bool {
					//		for _, v2 := range s.CheckList {
					//			if v2.ID == item.ID {
					//				return true
					//			}
					//		}
					//		return false
					//	})
				}
			}()
			ti = append(ti, tk)
		}
	}
	ct.task[taskId] = ti
	fmt.Printf("\n --check--list------%d", len(ct.task[taskId]))

}

func (ct *check_test) CanenlByTaskId(taskId string) {
	ct.mutex.Lock()
	defer ct.mutex.Unlock()
	if tl, ok := ct.task[taskId]; ok {
		for _, timer := range tl {
			timer.Stop()
		}
		delete(ct.task, taskId)
	}
}
