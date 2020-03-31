package auth

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/models"
	"gorm-learning/utils/errors"
)

func GetAuthUserId(c *gin.Context) (uint, *errors.RestErr) {
	return models.OauthAccessToken{}.GetAuthUserId(c)
}

func GetAuthUser(c *gin.Context) (*models.User, *errors.RestErr) {
	return models.OauthAccessToken{}.GetAuthUser(c)
}
