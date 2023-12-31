package controller

import (
	"github.com/gin-gonic/gin"
	"golang-restful-api-framework/internal/service"
	"golang-restful-api-framework/internal/utils"
)

type UserController interface {
	GetUserById(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func InitUserController(userService service.UserService) UserControllerImpl {
	return UserControllerImpl{svc: userService}
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	userVO := u.svc.GetUserById(c)
	utils.Success(c, userVO)
}
