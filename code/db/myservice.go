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
	Url    string `json:"url"`
	Local  string `json:"local"`
	Branch string `json:"branch"`
}
type OpenCheckT struct {
	IpPort           string `json:"ip_port"`
	OuterNet         string `json:"outer_net"`
	InternalNet      string `json:"internal_net"`
	IpPortState      bool   `json:"ip_port_state"`
	InternalNetState bool   `json:"internal_net_state"`
	OuterNetState    bool   `json:"outer_net_state"`
}
type MyserviceT struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Desc              string     `json:"desc"`
	HostId            string     `json:"host_id"`
	ServiceType       string     `json:"service_type"`  //qianhou   hou
	DeployMethod      string     `json:"deploy_method"` //qianhou   hou
	GitHou            GitT       `json:"git_hou"`
	GitQian           GitT       `json:"git_qian"`
	IsOpenCheck       bool       `json:"is_open_check"`
	IsOpenCheckQian   bool       `json:"is_open_check_qian"`
	IsOpenCheckHou    bool       `json:"is_open_check_hou"`
	CheckTimeInterval string     `json:"check_time_interval"`
	OpenCheckQian     OpenCheckT `json:"open_check_qian"`
	OpenCheckHou      OpenCheckT `json:"open_check_hou"`
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
