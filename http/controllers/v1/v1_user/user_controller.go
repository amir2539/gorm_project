package v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm-learning/http/models"
	"gorm-learning/http/requests/v1/user"
	user2 "gorm-learning/http/resources/v1/user"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/parser"
	"net/http"
)

type userControllerInterface interface {
	Show(*gin.Context)
	SignUp(*gin.Context)
}

type userController struct{}

var UserController userControllerInterface = &userController{}

func (userController) Show(c *gin.Context) {

	userId, err := parser.GetIntegerParam(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user models.User
	user.ID = userId
	if err := user.FindWithId(); err != nil {
		c.JSON(err.Status, err)
		return
	}

	var resource user2.UserShowResource
	result := resource.Get(user).Append(true, "")

	c.JSON(http.StatusOK, result)
}

func (userController) SignUp(c *gin.Context) {
	var request user.UserSaveRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		err := errors.NewBadRequestErr("invalid json body")

		c.JSON(err.Status, err)
		return
	}

	var user models.User
	copier.Copy(&user, &request)

	token, err := user.Save()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(200, user2.UserSaveResource{
		Success: true,
		Message: "signed up successfully",
		Data: gin.H{
			"token": token,
		},
	})
}
