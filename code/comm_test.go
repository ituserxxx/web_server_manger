package main

import (
	"fmt"
	"server_manager/db"
	"server_manager/service"
	"strconv"
	"testing"
	"time"
)

func Test_01(t *testing.T) {
	var a = db.NotifyT{}
	b := func() db.NotifyT { return db.NotifyT{} }()
	if a == b {
		println(1111)
	}
}
func Test_1(t *testing.T) {
	println(time.Now().Format("2006-01-02 15:04:05"))
}
func Test_b(t *testing.T) {
	db.Init()
	l, err := db.MyserviceDb.QueryRow(func(s db.MyserviceT) bool {
		return true
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t2 := range l {
		if t2.IsOpenCheck {

			var ti = make([]*time.Ticker, 0)
			for _, item := range t2.CheckList {
				sj, _ := strconv.Atoi(item.Time)
				if item.Value != "" {
					tk := time.NewTicker(time.Duration(sj) * time.Second)
					go func(it db.CheckT) {
						for range tk.C {
							fmt.Printf("\n --check-- gofunc()----value-%v", item.Value)
							err := db.MyserviceDb.Update(
								func(s db.MyserviceT) db.MyserviceT {
									for i, v2 := range s.CheckList {
										if v2.ID == it.ID {
											s.CheckList[i].State = true
										}
									}
									return s
								},
								func(s db.MyserviceT) bool {
									if s.ID == "apnls" {
										return true
									}
									return false
								})
							if err != nil {
								fmt.Printf("\nerr := db.MyserviceDb.Update is %v", err.Error())
							}
						}
					}(item)
					ti = append(ti, tk)
				}
			}
		}
	}
	time.Sleep(100 * time.Minute)

}
func Test_a(t *testing.T) {
	//println(Tool.CheckTcp("172.16.9.103:6008"))
	println(service.Tool.CheckHttp("http://172.16.9.103:6003"))
}
func Test_c(t *testing.T) {
	db.Init()
	id := "xl9lh"
	db.MyserviceDb.Update(func(s db.MyserviceT) db.MyserviceT {
		for i, v2 := range s.CheckList {
			if v2.ID == id {
				s.CheckList[i].State = true
			}
		}
		//fmt.Printf("111=>%#v", s)
		return s
	}, func(s db.MyserviceT) bool {
		for _, v2 := range s.CheckList {
			fmt.Printf("v2.id=%s  id=%s", v2.ID, id)
			if v2.ID == id {
				return true
			}
		}
		return false
	})
}
