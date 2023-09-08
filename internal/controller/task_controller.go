package controller

import (
	"github.com/gin-gonic/gin"
	"golang-restful-api-framework/internal/service"
)

type TaskController interface {
	GetTaskById(c *gin.Context)
}

type TaskControllerImpl struct {
	svc service.TaskService
}

func InitTaskController(taskService service.TaskService) TaskControllerImpl {
	return TaskControllerImpl{svc: taskService}
}

func (u TaskControllerImpl) GetTaskById(c *gin.Context) {
	u.svc.GetTaskById(c)
}
