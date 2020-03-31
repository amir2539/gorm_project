package v1_group

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/models"
	"gorm-learning/http/requests/group"
	"gorm-learning/utils/auth"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/resource"
	"net/http"
)

type groupControllerInterface interface {
	Create(*gin.Context)
}

type groupController struct{}

var GroupController groupControllerInterface = &groupController{}

func (groupController) Create(c *gin.Context) {
	var request group.GroupCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		err := errors.NewBadRequestErr("invalid json body")

		c.JSON(err.Status, err)
		return
	}

	userId, err := auth.GetAuthUserId(c)
	if err != nil {
		restErr := errors.NewUnauthorizedError()
		c.JSON(restErr.Status, restErr)
		return
	}

	var group models.Group
	group.UserID = userId
	group.Name = request.Name

	groupToken, err := group.Create()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, resource.ActionResource(true, "group created", gin.H{
		"group_token": groupToken,
	}))

}
