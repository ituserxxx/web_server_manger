package db

import (
	tiny "github.com/Yiwen-Chan/tinydb"
)

var NotifyDb *notifyDbT

type notifyDbT struct {
	TableName string `json:"table_name"`
	storage   *tiny.StorageJSON
	database  *tiny.Database
	NotifyDbT NotifyT `json:"notify_db_t"`
}

type NotifyT struct {
	ID          string `json:"token"`
	QywxWebhook string `json:"qywx_webhook"`
	QywxIsOpen  bool   `json:"qywx_is_open"`
	DdWebhook   string `json:"dd_webhook"`
	DdIsOpen    bool   `json:"dd_is_open"`
}

func NewNotifyDb() {
	NotifyDb = &notifyDbT{
		TableName: "web_manager_db/notify.json",
	}
	storage, err := tiny.JSONStorage(NotifyDb.TableName)
	if err != nil {
		panic(err)
	}
	NotifyDb.storage = storage

	database, err := tiny.TinyDB(NotifyDb.storage)
	if err != nil {
		panic(err)
	}
	NotifyDb.database = database
	NotifyDb.init()
}
func (nt *notifyDbT) init() {
	d, err := nt.QueryRowOne(func(s NotifyT) bool {
		return true
	})
	d2 := NotifyT{}
	if err != nil || d == d2 {
		nt.Insert(NotifyT{
			ID:          "1",
			QywxWebhook: "",
			QywxIsOpen:  false,
			DdWebhook:   "",
			DdIsOpen:    false,
		})
	}
}
func (nt *notifyDbT) getTable() *tiny.Table[NotifyT] {
	if nt.database == nil || nt.storage == nil {
		NewNotifyDb()
	}
	return tiny.GetTable[NotifyT](nt.database)
}

func (nt *notifyDbT) Insert(data NotifyT) error {

	err := nt.getTable().Insert(data)
	if err != nil {
		return err
	}
	return nil
}
func (nt *notifyDbT) Update(data func(s NotifyT) NotifyT, condition func(s NotifyT) bool) error {

	err := nt.getTable().Update(data, condition)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
func (nt *notifyDbT) QueryRowOne(condition func(s NotifyT) bool) (NotifyT, error) {
	var tmp NotifyT

	ls, err := nt.getTable().Select(condition)
	if err != nil {
		return tmp, err
	}
	if len(ls) != 1 {
		return tmp, nil
	}
	return ls[0], nil
}
func (nt *notifyDbT) QueryRow(condition func(s NotifyT) bool) ([]NotifyT, error) {
	data, err := nt.getTable().Select(condition)
	if err != nil {
		return data, err
	}
	return data, nil
}
