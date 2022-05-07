package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectTemplate := "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	user := "testuser"
	passwd := "testpass"
	database := "testdb"
	dsn := fmt.Sprintf(connectTemplate, user, passwd, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	// create
	insert := &User{}
	if result := db.Create(insert); result.Error != nil {
		fmt.Println("create")
		fmt.Println(result.Error)
	}

	// get
	searched := &User{
		FirstName: "first1",
		LastName:  "last1",
	}
	if result := db.First(searched, &User{}); result.Error != nil {
		fmt.Println("get")
		fmt.Println(result.Error)
	}
	fmt.Println(searched)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
