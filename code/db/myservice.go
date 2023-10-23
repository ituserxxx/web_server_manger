package db

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

var MyserviceDb *myserviceDbT

type myserviceDbT struct {
	TableName  string `json:"table_name"`
	storage    *tiny.StorageJSON
	database   *tiny.Database
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
	storage, err := tiny.JSONStorage(MyserviceDb.TableName)
	if err != nil {
		panic(err)
	}
	MyserviceDb.storage = storage

	database, err := tiny.TinyDB(MyserviceDb.storage)
	if err != nil {
		panic(err)
	}
	MyserviceDb.database = database

}
func (ms *myserviceDbT) getTable() *tiny.Table[MyserviceT] {
	if ms.database == nil || ms.storage == nil {
		NewMyserviceDb()
	}
	return tiny.GetTable[MyserviceT](ms.database)
}

func (ms *myserviceDbT) Insert(data MyserviceT) error {

	//defer ms.database.Close()
	err := ms.getTable().Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (ms *myserviceDbT) Update(data func(s MyserviceT) MyserviceT, condition func(s MyserviceT) bool) error {

	err := ms.getTable().Update(data, condition)
	if err != nil {
		return err
	}
	return nil
}
func (ms *myserviceDbT) QueryRowOne(condition func(s MyserviceT) bool) (MyserviceT, error) {
	var tmp MyserviceT

	ls, err := ms.getTable().Select(condition)
	if err != nil {
		return tmp, err
	}
	if len(ls) != 1 {
		return tmp, nil
	}
	return ls[0], nil
}
func (ms *myserviceDbT) QueryRow(condition func(s MyserviceT) bool) ([]MyserviceT, error) {

	data, err := ms.getTable().Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (ms *myserviceDbT) Del(id string) error {

	_, err := ms.getTable().Delete(func(s MyserviceT) bool {
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
