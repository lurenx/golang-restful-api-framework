package router

import "golang-restful-api-framework/internal/setup"

func LoadUser(i *setup.Initialization) {
	rg := i.Router.RouterGroup.Group("/v1")
	{
		rg.GET("/user/:id", i.UserController.GetUserById)
	}
}
