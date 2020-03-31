package auth

import (
	"github.com/gin-gonic/gin"
	"gorm-learning/http/models"
	"gorm-learning/utils/logger"
	"gorm-learning/utils/parser"
)

func AuthUserMiddleware(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")

	token, err := parser.ReadBearerToken(tokenHeader)
	if err != nil {
		logger.Info("read error")
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	_, validErr := models.OauthAccessToken{}.VerifyToken(token)
	if validErr != nil {
		logger.Info("validate error")
		c.AbortWithStatusJSON(validErr.Status, validErr)
		return
	}
	c.Next()

}
