package db

import (
	"errors"
	tiny "github.com/Yiwen-Chan/tinydb"
)

var AdminDb *adminDbT

type adminDbT struct {
	TableName string `json:"table_name"`
	AdminDbT  AdminT `json:"admin_db_t"`
}

type AdminT struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

func NewAdminDb() {
	AdminDb = &adminDbT{
		TableName: "web_manager_db/admin.json",
	}
	a, err := AdminDb.QueryRow(func(s AdminT) bool {
		if s.Name == "admin" {
			return true
		}
		return false
	})
	if err != nil {
		println("init admin fail : " + err.Error())
	} else {
		if len(a) == 0 {
			AdminDb.Insert(AdminT{
				Token:  "1",
				Name:   "admin",
				Passwd: "admin123..",
			})
		}
	}
}
func (ad *adminDbT) getTable() (*tiny.Table[AdminT], *tiny.Database, error) {
	storage, err := tiny.JSONStorage(ad.TableName)
	if err != nil {
		return nil, nil, err
	}
	database, err := tiny.TinyDB(storage)

	table := tiny.GetTable[AdminT](database)
	return table, database, nil
}

func (ad *adminDbT) Insert(data AdminT) error {
	table, database, err := ad.getTable()
	if err != nil {
		return err
	}
	defer database.Close()
	err = table.Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminDbT) Update(data func(s AdminT) AdminT, condition func(s AdminT) bool) error {
	table, database, err := ad.getTable()
	if err != nil {
		return err
	}
	defer database.Close()
	err = table.Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
func (ad *adminDbT) QueryRowOne(condition func(s AdminT) bool) (AdminT, error) {
	var tmp AdminT
	table, database, err := ad.getTable()
	if err != nil {
		return tmp, err
	}
	defer database.Close()
	ls, err := table.Select(condition)
	if err != nil {
		return tmp, err
	}
	if len(ls) != 1 {
		return tmp, errors.New("query result is no")
	}
	return ls[0], nil
}
func (ad *adminDbT) QueryRow(condition func(s AdminT) bool) ([]AdminT, error) {
	var data = make([]AdminT, 0)
	table, database, err := ad.getTable()
	if err != nil {
		return data, err
	}
	defer database.Close()
	data, err = table.Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (ad *adminDbT) Del(token string) error {
	table, database, err := ad.getTable()
	if err != nil {
		return err
	}
	defer database.Close()
	_, err = table.Delete(func(s AdminT) bool {
		if s.Token == token {
			return true
		}
		return false
	})
	if err != nil {
		return err
	}
	return nil
}
