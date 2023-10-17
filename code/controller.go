package main

import (
	"fmt"
	"net/http"
)

var MyhostController = &myhostController{}
var MyserviceController = &myserviceController{}

type myhostController struct {
}

func (mhc *myhostController) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login successful! Welcome, %s!", "xxx")
}
func (mhc *myhostController) Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login successful! Welcome, %s!", "xxx")
}

type myserviceController struct {
}
