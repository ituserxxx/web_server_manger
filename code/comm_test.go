package main

import (
	"fmt"
	"server_manager/db"
	"testing"
)

func Test_b(t *testing.T) {
	i := 0
	for i < 3 {
		println(Tool.Krand())
		i++
	}
}
func Test_a(t *testing.T) {
	//println(Tool.CheckTcp("172.16.9.103:6008"))
	println(Tool.CheckHttp("http://172.16.9.103:6003"))
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
