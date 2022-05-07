package main

import (
	"fmt"
	"time"

	"github.com/pepese/golang-sample/gorm/app/infrastructure/datastore"
)

func main() {
	db := datastore.Gorm()
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
