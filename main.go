package main

import (
	"LemonMovie/global"
	"LemonMovie/models"
	_ "LemonMovie/routers"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(models.Models...)
	if err != nil {
		fmt.Println(err)
	}
	global.DB = db
}

func main() {
	web.Run("127.0.0.1:8080")
}
