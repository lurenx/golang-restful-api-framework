package setup

import (
	"golang-restful-api-framework/internal/config"
	"golang-restful-api-framework/internal/controller"
	"golang-restful-api-framework/internal/dao"
	"golang-restful-api-framework/internal/service"
)

type Initialization struct {
	Config config.Config
	Router Router

	userDao        dao.UserDAO
	userSevice     service.UserService
	UserController controller.UserController

	taskDao        dao.TaskDAO
	taskSevice     service.TaskService
	TaskController controller.TaskController
}

func NewInitialization(
	config config.Config,
	router Router,
	userDao dao.UserDAO,
	userSevice service.UserService,
	userController controller.UserController,

	taskDao dao.TaskDAO,
	taskSevice service.TaskService,
	taskController controller.TaskController,

) *Initialization {
	return &Initialization{
		Config:  config,
		Router:  router,
		userDao: userDao, taskDao: taskDao,
		userSevice: userSevice, taskSevice: taskSevice,
		UserController: userController, TaskController: taskController,
	}
}
