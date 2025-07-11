package api

import (
	"Task_04/middlewares"
	"Task_04/routers"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	//创建引擎实例
	engine := gin.Default()
	//创建路由分组

	routerGroup := engine.Group("/v1")
	{
		routerGroup.POST("/user/register", routers.Register)
		routerGroup.POST("/user/login", routers.Login)
	}

	group := engine.Group("/v2")
	group.Use(middlewares.ValidateToken())
	{
		group.POST("/article/create", routers.CreateArticle)
		group.POST("/article/getArticles", routers.GetArticles)
		group.POST("/article/getArticle", routers.GetArticle)
		group.POST("/article/updateArticle", routers.UpdateArticle)
		group.POST("/article/deleteArticle", routers.DeleteArticle)
		group.POST("/article/createComments", routers.CreateComments)
		group.POST("/article/readComment", routers.ReadComment)
	}
	engine.Run(":8600")
}
