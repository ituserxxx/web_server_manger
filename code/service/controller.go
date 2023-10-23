package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server_manager/db"
)

type resp struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

func respHandle(w http.ResponseWriter, code int, data interface{}) ([]byte, error) {
	response := resp{
		Data: data,
		Code: code,
	}
	w.Header().Set("Content-Type", "application/json") // 设置响应头的Content-Type为"application/json"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}
func respOk(w http.ResponseWriter, data interface{}) {
	jsonResponse, err := respHandle(w, 33, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))

}
func respErr(w http.ResponseWriter, data string) {
	jsonResponse, err := respHandle(w, 44, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, string(jsonResponse))
}

var UserController = &userController{}
var MyhostController = &myhostController{}
var MyserviceController = &myserviceController{}
var NotifyController = &notifyController{}

// ------------UserController
type userController struct {
}
type reqMucLogin struct {
	db.AdminT
}

func (muc *userController) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMucLogin
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}

	d, err := db.AdminDb.QueryRowOne(func(s db.AdminT) bool {
		if s.Name == params.Name && s.Passwd == params.Passwd {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	d2 := db.AdminT{}
	if d == d2 {
		respErr(w, "no data")
		return
	}
	respOk(w, d)
}

func (muc *userController) Info(w http.ResponseWriter, r *http.Request) {
	xToken := r.Header.Get("X-token")
	d, err := db.AdminDb.QueryRowOne(func(s db.AdminT) bool {
		if s.Token == xToken {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	d2 := db.AdminT{}
	if d == d2 {
		respErr(w, "no data")
		return
	}
	respOk(w, d)
}

// ------------MyhostController
type myhostController struct {
}

type reqMhcAdd struct {
	db.HostT
}
type reqMhcUpdate struct {
	db.HostT
}
type reqMhcDel struct {
	Id string `json:"id"`
}

func (mhc *myhostController) Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMhcAdd
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	insertData := db.HostT{
		ID:        Tool.Krand(),
		IP:        params.IP,
		SshUser:   params.SshUser,
		SshPasswd: params.SshPasswd,
		Desc:      params.Desc,
	}
	err = db.MyHostDb.Insert(insertData)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, insertData)
}
func (mhc *myhostController) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMhcUpdate
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	err = db.MyHostDb.Update(func(s db.HostT) db.HostT {
		s = db.HostT{
			ID:        params.ID,
			IP:        params.IP,
			SshUser:   params.SshUser,
			SshPasswd: params.SshPasswd,
			Desc:      params.Desc,
		}
		return s
	}, func(s db.HostT) bool {
		if s.ID == params.ID {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	d, err := db.MyHostDb.QueryRow(func(s db.HostT) bool {
		if s.ID == params.ID {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	if len(d) != 1 {
		respErr(w, "query is fial over")
		return
	}
	respOk(w, d[0])
}
func (mhc *myhostController) Del(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMhcDel
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	err = db.MyHostDb.Del(params.Id)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, "")
}
func (mhc *myhostController) GetList(w http.ResponseWriter, r *http.Request) {

	l, err := db.MyHostDb.QueryRow(func(s db.HostT) bool {
		return true
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, l)
}

// ------------MyserviceController
type myserviceController struct {
}
type reqMscAdd struct {
	db.MyserviceT
}
type reqMscUpdate struct {
	db.MyserviceT
}
type reqMscDel struct {
	Id string `json:"id"`
}

func (msc *myserviceController) Add(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMscAdd
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}

	for i, _ := range params.CheckList {
		params.CheckList[i].ID = Tool.Krand()
	}
	insertData := db.MyserviceT{
		ID:           Tool.Krand(),
		Name:         params.Name,
		Desc:         params.Desc,
		HostId:       params.HostId,
		ServiceType:  params.ServiceType,
		DeployMethod: params.DeployMethod,

		IsOpenCheck: params.IsOpenCheck,
		GitList:     params.GitList,
		CheckList:   params.CheckList,
	}
	err = db.MyserviceDb.Insert(insertData)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	if params.IsOpenCheck {
		CheckTask.Add(insertData.ID, insertData)
	}
	respOk(w, insertData)
}
func (msc *myserviceController) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMscUpdate
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	CheckTask.CanenlByTaskId(params.ID)
	for i, _ := range params.CheckList {
		params.CheckList[i].ID = Tool.Krand()
	}
	upMod := db.MyserviceT{
		ID:           params.ID,
		Name:         params.Name,
		Desc:         params.Desc,
		HostId:       params.HostId,
		ServiceType:  params.ServiceType,
		DeployMethod: params.DeployMethod,
		IsOpenCheck:  params.IsOpenCheck,
		GitList:      params.GitList,
		CheckList:    params.CheckList,
	}
	err = db.MyserviceDb.Update(func(s db.MyserviceT) db.MyserviceT {
		s = upMod
		return s
	}, func(s db.MyserviceT) bool {
		if s.ID == params.ID {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}

	if params.IsOpenCheck {
		CheckTask.Add(params.ID, upMod)
	} else {
		CheckTask.CanenlByTaskId(params.ID)
	}
	respOk(w, "ok")
}
func (msc *myserviceController) Del(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqMscDel
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	err = db.MyserviceDb.Del(params.Id)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, "")
}
func (msc *myserviceController) GetList(w http.ResponseWriter, r *http.Request) {

	l, err := db.MyserviceDb.QueryRow(func(s db.MyserviceT) bool {
		return true
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, l)
}

// ------------notifyController
type notifyController struct {
}
type reqNtcUpdate struct {
	db.NotifyT
}

func (ntc *notifyController) Get(w http.ResponseWriter, r *http.Request) {
	d, err := db.NotifyDb.QueryRowOne(func(s db.NotifyT) bool {
		return true
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, d)
}
func (ntc *notifyController) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respErr(w, err.Error())
		return
	}
	var params reqNtcUpdate
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respErr(w, err.Error())
		return
	}
	upMod := db.NotifyT{
		ID:          params.ID,
		DdIsOpen:    params.DdIsOpen,
		DdWebhook:   params.DdWebhook,
		QywxIsOpen:  params.QywxIsOpen,
		QywxWebhook: params.QywxWebhook,
	}
	err = db.NotifyDb.Update(func(s db.NotifyT) db.NotifyT {
		s = upMod
		return s
	}, func(s db.NotifyT) bool {
		if s.ID == params.ID {
			return true
		}
		return false
	})
	if err != nil {
		respErr(w, err.Error())
		return
	}
	respOk(w, "ok")
}
