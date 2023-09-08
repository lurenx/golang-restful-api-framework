//go:build wireinject
// +build wireinject

// go:build wireinject
package setup

import (
	"github.com/google/wire"
	"golang-restful-api-framework/internal/controller"
	"golang-restful-api-framework/internal/dao"
	"golang-restful-api-framework/internal/service"
)

var cfg = wire.NewSet(InitConfigFile, InitConfig)

var db = wire.NewSet(InitSqlite)

// DAO set
var userDAOSet = wire.NewSet(
	dao.InitUserDAO,
	wire.Bind(new(dao.UserDAO), new(dao.UserDAOImpl)),
)
var taskDAOSet = wire.NewSet(
	dao.InitTaskDAO,
	wire.Bind(new(dao.TaskDAO), new(dao.TaskDAOImpl)),
)

// service set
var userServiceSet = wire.NewSet(
	service.InitUserService,
	wire.Bind(new(service.UserService), new(service.UserServiceImpl)),
)
var taskServiceSet = wire.NewSet(
	service.InitTaskService,
	wire.Bind(new(service.TaskService), new(service.TaskServiceImpl)),
)

// controller set
var userControllerSet = wire.NewSet(
	controller.InitUserController,
	wire.Bind(new(controller.UserController), new(controller.UserControllerImpl)),
)

var taskControllerSet = wire.NewSet(
	controller.InitTaskController,
	wire.Bind(new(controller.TaskController), new(controller.TaskControllerImpl)),
)
var router = wire.NewSet(InitRouter)

func Init(configFile string) *Initialization {
	wire.Build(
		NewInitialization,
		router,
		userControllerSet, taskControllerSet,
		userServiceSet, taskServiceSet,
		userDAOSet, taskDAOSet,
		db,
		cfg,
	)
	return nil
}
