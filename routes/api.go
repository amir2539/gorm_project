package routes

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/routes/v1/auth_routes"
	"gorm-learning/routes/v1/group_routes"
	"gorm-learning/routes/v1/user_routes"
	"net/http"
)

func InitRoutes() *gin.Engine {

	//version 1 apo

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello index api")
	})

	v1 := router.Group("/v1")
	{
		auth_routes.AuthApis(v1)
		user_routes.UserApis(v1)
		group_routes.GroupApis(v1)

	}

	//load users api

	return router

}
