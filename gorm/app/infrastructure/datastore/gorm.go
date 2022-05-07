package datastore

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	gm "gorm.io/gorm"
)

var (
	gorm  *gm.DB
	gOnce sync.Once
)

func Gorm() *gm.DB {
	initGorm()
	return gorm
}

func initGorm() {
	gOnce.Do(func() {
		connectTemplate := "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		user := "testuser"
		passwd := "testpass"
		database := "testdb"
		dsn := fmt.Sprintf(connectTemplate, user, passwd, database)
		db, err := gm.Open(mysql.Open(dsn), &gm.Config{})
		if err != nil {
			panic(err.Error())
		}
		gorm = db
	})
}
