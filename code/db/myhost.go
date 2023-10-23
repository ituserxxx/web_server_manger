package db

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

var MyHostDb *myHostDbT

type myHostDbT struct {
	TableName string `json:"table_name"`
	storage   *tiny.StorageJSON
	database  *tiny.Database
	HostT     HostT `json:"host_t"`
}
type HostT struct {
	ID        string `json:"id"`
	IP        string `json:"ip"`
	SshUser   string `json:"ssh_user"`
	SshPasswd string `json:"ssh_passwd"`
	Desc      string `json:"desc"`
}

func NewMyhostDb() {
	MyHostDb = &myHostDbT{
		TableName: "web_manager_db/myhost.json",
	}
	storage, err := tiny.JSONStorage(MyHostDb.TableName)
	if err != nil {
		panic(err)
	}
	MyHostDb.storage = storage

	database, err := tiny.TinyDB(MyHostDb.storage)
	if err != nil {
		panic(err)
	}
	MyHostDb.database = database
}
func (msh *myHostDbT) getTable() *tiny.Table[HostT] {

	if msh.database == nil || msh.storage == nil {
		NewMyhostDb()
	}
	return tiny.GetTable[HostT](msh.database)
}

func (msh *myHostDbT) Insert(data HostT) error {

	err := msh.getTable().Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (msh *myHostDbT) Update(data func(s HostT) HostT, condition func(s HostT) bool) error {

	err := msh.getTable().Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
func (msh *myHostDbT) QueryRowOne(condition func(s HostT) bool) (HostT, error) {
	var tmp HostT

	ls, err := msh.getTable().Select(condition)
	if err != nil {
		return tmp, err
	}
	if len(ls) != 1 {
		return tmp, nil
	}
	return ls[0], nil
}
func (msh *myHostDbT) QueryRow(condition func(s HostT) bool) ([]HostT, error) {
	data, err := msh.getTable().Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (msh *myHostDbT) Del(id string) error {
	_, err := msh.getTable().Delete(func(s HostT) bool {
		if s.ID == id {
			return true
		}
		return false
	})
	if err != nil {
		return err
	}
	return nil
}
