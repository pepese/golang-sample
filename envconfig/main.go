package main

import (
	"fmt"

	"github.com/pepese/golang-sample/envconfig/app"
)

func main() {
	conf := app.Config()
	fmt.Println(conf.AppName)
}
