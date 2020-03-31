package auth_routes

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/controllers/v1/v1_auth"
	"gorm-learning/http/middleware/auth"
)

func AuthApis(router gin.IRouter) {

	authApi := router.Group("/user")
	{
		authApi.POST("/login", v1_auth.AuthController.Login)
	}

	router.Use(auth.AuthUserMiddleware)
	authWithMiddleware := router.Group("/user")
	{
		authWithMiddleware.GET("/authenticate", v1_auth.AuthController.VerifyToken)
	}

}
