package main

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()

	user := &User{}
	if err := validate.Struct(user); err != nil {
		fmt.Println("Validation Error !")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Print("err.Namespace(): ")
			fmt.Println(err.Namespace())
			fmt.Print("err.Field(): ")
			fmt.Println(err.Field())
			fmt.Print("err.StructNamespace(): ")
			fmt.Println(err.StructNamespace())
			fmt.Print("err.StructField(): ")
			fmt.Println(err.StructField())
			fmt.Print("err.Tag(): ")
			fmt.Println(err.Tag())
			fmt.Print("err.ActualTag(): ")
			fmt.Println(err.ActualTag())
			fmt.Print("err.Kind(): ")
			fmt.Println(err.Kind())
			fmt.Print("err.Type(): ")
			fmt.Println(err.Type())
			fmt.Print("err.Value(): ")
			fmt.Println(err.Value())
			fmt.Print("err.Param(): ")
			fmt.Println(err.Value())
			fmt.Println()
		}
	} else {
		fmt.Println("No Validation Error.")
	}

	user = &User{
		Name: "1234",
	}
	if err := validate.Struct(user); err != nil {
		fmt.Println("Validation Error !")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Print("err.Namespace(): ")
			fmt.Println(err.Namespace())
			fmt.Print("err.Field(): ")
			fmt.Println(err.Field())
			fmt.Print("err.StructNamespace(): ")
			fmt.Println(err.StructNamespace())
			fmt.Print("err.StructField(): ")
			fmt.Println(err.StructField())
			fmt.Print("err.Tag(): ")
			fmt.Println(err.Tag())
			fmt.Print("err.ActualTag(): ")
			fmt.Println(err.ActualTag())
			fmt.Print("err.Kind(): ")
			fmt.Println(err.Kind())
			fmt.Print("err.Type(): ")
			fmt.Println(err.Type())
			fmt.Print("err.Value(): ")
			fmt.Println(err.Value())
			fmt.Print("err.Param(): ")
			fmt.Println(err.Value())
		}
	} else {
		fmt.Println("No Validation Error.")
	}
}

type User struct {
	Name string `validate:"required,alpha,gte=0,lte=10"`
}
