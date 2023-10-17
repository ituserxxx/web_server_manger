package db

import (
	"fmt"
	"testing"

	tiny "github.com/Yiwen-Chan/tinydb"
)

type student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func Test_a(t *testing.T) {
	storage, err := tiny.JSONStorage("test.json")
	database, err := tiny.TinyDB(storage)
	defer database.Close()
	table := tiny.GetTable[student](database)
	//table.Insert(student{1, "test"})
	//table.Insert(student{2, "test"})
	//table.Insert(student{3, "test"})
	students, err := table.Select(func(s student) bool { return true })
	if err != nil {
		return
	}
	fmt.Printf("\n students\n%#v", students)

	err = table.Update(func(s student) student {
		s.ID = 444
		return s
	}, func(s student) bool {
		if s.ID == 1 {
			return true
		}
		return false
	})
	if err != nil {
		panic(err.Error())
	}

	//students2, err := table.Select(func(s student) bool {
	//	if s.ID == 9 {
	//		return true
	//	}
	//	return false
	//})
	////if err != nil {
	////	return
	////}
	//fmt.Printf("\n students2\n%#v", students2)

	//students3, err := table.Delete(func(s student) bool {
	//	if s.ID == 2 {
	//		return true
	//	}
	//	return false
	//})
	//if err != nil {
	//	return
	//}
	//fmt.Printf("\n \n%#v", students3)

	students, err = table.Select(func(s student) bool { return true })
	if err != nil {
		return
	}
	fmt.Printf("\n students\n%#v", students)
}
