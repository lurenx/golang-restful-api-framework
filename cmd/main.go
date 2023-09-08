package main

import (
	"golang-restful-api-framework/internal/router"
	"golang-restful-api-framework/internal/setup"
	"strconv"
)

func main() {
	i := setup.Init("/tmp/hello.yml")
	router.LoadUser(i)
	router.LoadTask(i)
	i.Router.Engine.Run(":" + strconv.Itoa(i.Config.App.Port))
}
