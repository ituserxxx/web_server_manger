package db

import (
	"fmt"
	"os"
)

var dbFileList = []string{"myhost.json", "myservice.json", "admin.json", "notify.json"}

func Init() {
	initDbFile()
	NewMyserviceDb()
	NewMyhostDb()
	NewAdminDb()
	NewNotifyDb()

}

func initDbFile() {
	dbDir := "web_manager_db"

	// 判断目录是否存在
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			fmt.Println("创建目录失败:", err)
			return
		}

	}

	for _, s := range dbFileList {
		ns := dbDir + "/" + s
		if _, err := os.Stat(ns); os.IsNotExist(err) {
			// 文件不存在，创建文件
			file, err := os.Create(ns)
			if err != nil {
				fmt.Println("创建文件失败：", err)
				return
			}
			defer file.Close()
		}
	}

}
