package db

import (
	"errors"
	tiny "github.com/Yiwen-Chan/tinydb"
)

var MyHostDb *myHostDbT

type myHostDbT struct {
	TableName string `json:"table_name"`
	HostT     HostT  `json:"host_t"`
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
}
func (msh *myHostDbT) getTable() (*tiny.Table[HostT], *tiny.Database, error) {
	storage, err := tiny.JSONStorage(msh.TableName)
	if err != nil {
		return nil, nil, err
	}
	database, err := tiny.TinyDB(storage)

	table := tiny.GetTable[HostT](database)
	return table, database, nil
}

func (msh *myHostDbT) Insert(data HostT) error {
	table, database, err := msh.getTable()
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
func (msh *myHostDbT) Update(data func(s HostT) HostT, condition func(s HostT) bool) error {
	table, database, err := msh.getTable()
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
func (msh *myHostDbT) QueryRowOne(condition func(s HostT) bool) (HostT, error) {
	var tmp HostT
	table, database, err := msh.getTable()
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
func (msh *myHostDbT) QueryRow(condition func(s HostT) bool) ([]HostT, error) {
	var data = make([]HostT, 0)
	table, database, err := msh.getTable()
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
func (msh *myHostDbT) Del(id string) error {
	table, database, err := msh.getTable()
	if err != nil {
		return err
	}
	defer database.Close()
	_, err = table.Delete(func(s HostT) bool {
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
