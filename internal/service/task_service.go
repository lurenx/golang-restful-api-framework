package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang-restful-api-framework/internal/dao"
	"golang-restful-api-framework/internal/utils"
	"strconv"
)

type TaskService interface {
	GetTaskById(c *gin.Context)
}

type TaskServiceImpl struct {
	userDao dao.TaskDAO
}

func InitTaskService(useDao dao.TaskDAO) TaskServiceImpl {
	return TaskServiceImpl{userDao: useDao}
}
func (u TaskServiceImpl) GetTaskById(c *gin.Context) {
	log.Info("start to execute program get user by id")
	taskID, _ := strconv.Atoi(c.Param("id"))

	data, err := u.userDao.FindTaskById(taskID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
	}
	utils.Success(c, data)
}
