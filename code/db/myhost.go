package db

import tiny "github.com/Yiwen-Chan/tinydb"

var MyHostDb *MyHostDbT

type MyHostDbT struct {
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
	MyHostDb = &MyHostDbT{
		TableName: "./web_manager_db/myservice.json",
	}
}
func (msh *MyHostDbT) getTable() (*tiny.Table[HostT], error) {
	storage, err := tiny.JSONStorage(msh.TableName)
	if err != nil {
		return nil, err
	}
	database, err := tiny.TinyDB(storage)
	defer database.Close()
	table := tiny.GetTable[HostT](database)
	return table, nil
}

func (msh *MyHostDbT) insert(data HostT) error {
	table, err := msh.getTable()
	if err != nil {
		return err
	}
	err = table.Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (msh *MyHostDbT) Update(data func(s HostT) HostT, condition func(s HostT) bool) error {
	table, err := msh.getTable()
	if err != nil {
		return err
	}
	err = table.Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (msh *MyHostDbT) QueryRow(condition func(s HostT) bool) ([]HostT, error) {
	var data = make([]HostT, 0)
	table, err := msh.getTable()
	if err != nil {
		return data, err
	}
	data, err = table.Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
