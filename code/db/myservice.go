package db

import (
	"errors"
	tiny "github.com/Yiwen-Chan/tinydb"
)

var MyserviceDb *myserviceDbT

type myserviceDbT struct {
	TableName  string     `json:"table_name"`
	MyserviceT MyserviceT `json:"myservice_t"`
}
type GitT struct {
	Url  string `json:"url"` //git resp url
	Desc string `json:"desc"`
}
type CheckT struct {
	ID    string `json:"id"`
	Type  string `json:"type"`  // http tcp
	Value string `json:"value"` // url  ip:port
	Time  string `json:"time"`  // time interval (second)
	Desc  string `json:"desc"`
	State bool   `json:"state"`
}
type MyserviceT struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Desc         string   `json:"desc"`
	HostId       string   `json:"host_id"`
	ServiceType  string   `json:"service_type"`  //qianhou   hou
	DeployMethod string   `json:"deploy_method"` //qianhou   hou
	GitList      []GitT   `json:"git_list"`
	IsOpenCheck  bool     `json:"is_open_check"`
	CheckList    []CheckT `json:"check_list"`
}

func NewMyserviceDb() {
	MyserviceDb = &myserviceDbT{
		TableName: "web_manager_db/myservice.json",
	}
}
func (ms *myserviceDbT) getTable() (*tiny.Table[MyserviceT], *tiny.Database, error) {
	storage, err := tiny.JSONStorage(ms.TableName)
	if err != nil {
		return nil, nil, err
	}
	database, err := tiny.TinyDB(storage)

	table := tiny.GetTable[MyserviceT](database)
	return table, database, nil
}

func (ms *myserviceDbT) Insert(data MyserviceT) error {
	table, database, err := ms.getTable()
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
func (ms *myserviceDbT) Update(data func(s MyserviceT) MyserviceT, condition func(s MyserviceT) bool) error {
	table, database, err := ms.getTable()
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
func (ms *myserviceDbT) QueryRowOne(condition func(s MyserviceT) bool) (MyserviceT, error) {
	var tmp MyserviceT
	table, database, err := ms.getTable()
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
func (ms *myserviceDbT) QueryRow(condition func(s MyserviceT) bool) ([]MyserviceT, error) {
	var data = make([]MyserviceT, 0)
	table, database, err := ms.getTable()
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
func (ms *myserviceDbT) Del(id string) error {
	table, database, err := ms.getTable()
	if err != nil {
		return err
	}
	defer database.Close()
	_, err = table.Delete(func(s MyserviceT) bool {
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
