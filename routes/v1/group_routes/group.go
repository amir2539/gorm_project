package group_routes

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/controllers/v1/v1_group"
	"gorm-learning/http/middleware/auth"
)

func GroupApis(router gin.IRouter) {

	api := router.Group("/group")
	{
		api.Use(auth.AuthUserMiddleware)
		api.POST("/create", v1_group.GroupController.Create)
	}
}
