package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-sample/gin/app/usecase"
)

/*
UserGinRouter interface.
*/
type UserGinRouter interface {
	ListUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUsers(c *gin.Context)
}

type userGinRoute struct {
	uc usecase.User
}

func NewUserGinRoute(uc usecase.User) *userGinRoute {
	return &userGinRoute{
		uc: uc,
	}
}

func (ctr *userGinRoute) ListUsers(c *gin.Context) {
	uc := ctr.uc
	ctx := c.Request.Context()
	result, err := uc.List(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (ctr *userGinRoute) GetUser(c *gin.Context) {
	uc := ctr.uc
	ctx := c.Request.Context()
	result, err := uc.Get(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (ctr *userGinRoute) CreateUser(c *gin.Context) {
	uc := ctr.uc
	ctx := c.Request.Context()
	result, err := uc.Create(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (ctr *userGinRoute) UpdateUser(c *gin.Context) {
	uc := ctr.uc
	ctx := c.Request.Context()
	result, err := uc.Update(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (ctr *userGinRoute) DeleteUser(c *gin.Context) {
	uc := ctr.uc
	ctx := c.Request.Context()
	err := uc.Delete(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}
