package v1_auth

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/models"
	"gorm-learning/http/requests/v1/auth"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/resource"
	"net/http"
)

type authControllerInterface interface {
	Login(*gin.Context)
	VerifyToken(*gin.Context)
}

type authController struct{}

var AuthController authControllerInterface = &authController{}

func (*authController) Login(c *gin.Context) {

	var request auth.AuthLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		err := errors.NewBadRequestErr("invalid body")

		c.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := models.User{}.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, resource.ActionResource(true, "successfully logged in", gin.H{
		"token": token,
	}))
}

func (*authController) VerifyToken(c *gin.Context) {

	//it will pass middleware

	c.JSON(http.StatusOK, resource.ActionResource(true, "successfully logged in"))
}
