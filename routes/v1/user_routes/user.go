package user_routes

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/controllers/v1/v1_user"
	"gorm-learning/http/middleware/auth"
)

func UserApis(router gin.IRouter) {

	router.POST("/user/signup", v1_user.UserController.SignUp)

	router.Use(auth.AuthUserMiddleware)
	api := router.Group("/user")
	{
		api.GET("/get_user/:id", v1_user.UserController.Show)
	}

}
