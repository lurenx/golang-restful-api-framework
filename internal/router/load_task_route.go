package router

import "golang-restful-api-framework/internal/setup"

func LoadTask(i *setup.Initialization) {
	rg := i.Router.RouterGroup.Group("/v1")
	{
		rg.GET("/task/:id", i.TaskController.GetTaskById)
	}
}
