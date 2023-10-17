package db

import tiny "github.com/Yiwen-Chan/tinydb"

var MyserviceDb *MyserviceDbT

type MyserviceDbT struct {
	TableName  string     `json:"table_name"`
	MyserviceT MyserviceT `json:"myservice_t"`
}

type MyserviceT struct {
	ID        string `json:"id"`
	IP        string `json:"ip"`
	SshUser   string `json:"ssh_user"`
	SshPasswd string `json:"ssh_passwd"`
	Desc      string `json:"desc"`
}

func NewMyserviceDb() {
	MyserviceDb = &MyserviceDbT{
		TableName: "./web_manager_db/myservice.json",
	}
}
func (ms *MyserviceDbT) getTable() (*tiny.Table[HostT], error) {
	storage, err := tiny.JSONStorage(ms.TableName)
	if err != nil {
		return nil, err
	}
	database, err := tiny.TinyDB(storage)
	defer database.Close()
	table := tiny.GetTable[HostT](database)
	return table, nil
}

func (ms *MyserviceDbT) Insert(data HostT) error {
	table, err := ms.getTable()
	if err != nil {
		return err
	}
	err = table.Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (ms *MyserviceDbT) Update(data func(s HostT) HostT, condition func(s HostT) bool) error {
	table, err := ms.getTable()
	if err != nil {
		return err
	}
	err = table.Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (ms *MyserviceDbT) QueryRow(condition func(s HostT) bool) ([]HostT, error) {
	var data = make([]HostT, 0)
	table, err := ms.getTable()
	if err != nil {
		return data, err
	}
	data, err = table.Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
