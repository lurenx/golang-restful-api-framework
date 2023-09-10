package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"golang-restful-api-framework/internal/dao"
	"golang-restful-api-framework/internal/model/vo"
	"strconv"
)

type UserService interface {
	GetUserById(c *gin.Context) vo.UserVO
}

type UserServiceImpl struct {
	userDao dao.UserDAO
}

func InitUserService(useDao dao.UserDAO) UserServiceImpl {
	return UserServiceImpl{userDao: useDao}
}
func (u UserServiceImpl) GetUserById(c *gin.Context) vo.UserVO {
	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userDao.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
	}
	userVO := vo.UserVO{}
	copier.Copy(data, &userVO)
	return userVO
}
