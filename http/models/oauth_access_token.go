package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gorm-learning/database"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/logger"
	"gorm-learning/utils/parser"
	"os"
	"time"
)

const (
	ExpirationTimeDays = 90
)

// for generating token
type Claims struct {
	//Username string
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type OauthAccessToken struct {
	ModelID

	User        User `gorm:"foreignkey:user_id;association_foreignkey:id"` // use UserRefer as foreign key
	UserID      uint
	AccessToken string `gorm:"type:varchar(255);unique_index;not null" json:"access_token"`
	ModelTimeStamps
}

func (o OauthAccessToken) GenerateToken(userID uint) (string, *errors.RestErr) {
	claims := &Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * ExpirationTimeDays).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		logger.Error("make jwt error", err)
		return "", errors.NewInternalServerError("internal server error")
	}

	o.UserID = userID
	o.AccessToken = tokenString

	//UpdateOrCreate
	err = database.DB.Scopes(o.scopeUser(o.UserID)).Assign(OauthAccessToken{AccessToken: o.AccessToken}).FirstOrCreate(&o).Error
	if err != nil {
		logger.Error("access_token update user's token error", err)
		return "", errors.NewInternalServerError("database error")
	}

	return tokenString, nil

}

func (o OauthAccessToken) VerifyToken(token string) (uint, *errors.RestErr) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, errors.NewUnauthorizedError()
	}

	if !tkn.Valid {
		return 0, errors.NewUnauthorizedError()
	}

	//check user token in db
	if database.DB.Scopes(o.scopeToken(token)).Find(&o).RecordNotFound() {
		return 0, errors.NewUnauthorizedError()
	}

	return claims.UserId, nil
}

func (o OauthAccessToken) GetAuthUserId(c *gin.Context) (uint, *errors.RestErr) {
	tokenHeader := c.GetHeader("Authorization")
	token, err := parser.ReadBearerToken(tokenHeader)
	if err != nil {
		return 0, errors.NewUnauthorizedError()
	}

	return o.VerifyToken(token)

}

func (o OauthAccessToken) GetAuthUser(c *gin.Context) (*User, *errors.RestErr) {
	userId, err := o.GetAuthUserId(c)
	if err != nil {
		return nil, err
	}

	var user User
	user.ID = userId
	err = user.FindWithId()
	if err != nil {
		return nil, err
	}

	return &user, nil

}

//scopes
func (o *OauthAccessToken) scopeUser(userId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func (o *OauthAccessToken) scopeToken(token string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("access_token = ?", token)
	}
}
