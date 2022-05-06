package main

import (
	"github.com/pepese/golang-sample/gin/app/adapter/controller"
	"github.com/pepese/golang-sample/gin/app/infrastructure/server"
	"github.com/pepese/golang-sample/gin/app/usecase"
)

func main() {
	srv := server.NewGinServer()
	// usecase
	userUC := usecase.NewUser()
	// adapter/controller の設定
	userCtr := controller.NewUserGinRoute(userUC)
	user := srv.Group("/api/v1/users")
	{
		user.GET("", userCtr.ListUsers)
		user.GET("/:user_id", userCtr.GetUser)
		user.POST("", userCtr.CreateUser)
		user.PUT("/:user_id", userCtr.UpdateUser)
		user.DELETE("/:user_id", userCtr.DeleteUser)
	}
	srv.Run()
}
