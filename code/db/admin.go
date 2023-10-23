package db

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

var AdminDb *adminDbT

type adminDbT struct {
	TableName string `json:"table_name"`
	storage   *tiny.StorageJSON
	database  *tiny.Database
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
	storage, err := tiny.JSONStorage(AdminDb.TableName)
	if err != nil {
		panic(err)
	}
	AdminDb.storage = storage

	database, err := tiny.TinyDB(AdminDb.storage)
	if err != nil {
		panic(err)
	}
	AdminDb.database = database
	AdminDb.initAdminuser()
}
func (ad *adminDbT) initAdminuser() {
	a, err := ad.QueryRow(func(s AdminT) bool {
		if s.Name == "admin" {
			return true
		}
		return false
	})
	if err != nil {
		println("init admin fail : " + err.Error())
	} else {
		if len(a) == 0 {
			ad.Insert(AdminT{
				Token:  "1",
				Name:   "admin",
				Passwd: "admin123..",
			})
		}
	}
}
func (ad *adminDbT) getTable() *tiny.Table[AdminT] {
	if ad.database == nil || ad.storage == nil {
		NewAdminDb()
	}
	return tiny.GetTable[AdminT](ad.database)
}

func (ad *adminDbT) Insert(data AdminT) error {

	err := ad.getTable().Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminDbT) Update(data func(s AdminT) AdminT, condition func(s AdminT) bool) error {

	err := ad.getTable().Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
func (ad *adminDbT) QueryRowOne(condition func(s AdminT) bool) (AdminT, error) {
	var tmp AdminT

	ls, err := ad.getTable().Select(condition)
	if err != nil {
		return tmp, err
	}
	if len(ls) != 1 {
		return tmp, nil
	}
	return ls[0], nil
}
func (ad *adminDbT) QueryRow(condition func(s AdminT) bool) ([]AdminT, error) {

	data, err := ad.getTable().Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (ad *adminDbT) Del(token string) error {
	_, err := ad.getTable().Delete(func(s AdminT) bool {
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
