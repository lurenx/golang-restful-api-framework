package setup

import "github.com/gin-gonic/gin"

type Router struct {
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func InitRouter() Router {
	engine := gin.Default()
	routergroup := engine.Group("/api")
	return Router{Engine: engine, RouterGroup: routergroup}
}
