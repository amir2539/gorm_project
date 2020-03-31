package models

import (
	"github.com/jinzhu/gorm"
	"gorm-learning/database"
	"gorm-learning/utils/encrypting"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/logger"
)

const (
	StatusActive = "active"
)

type User struct {
	ModelID
	Name     string `gorm:"type:varchar(100); not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index; not null" json:"email"`
	Password string `json:"-"`
	Status   string `gorm:"type:varchar(25)" json:"status"`

	// groups relation
	Groups []Group `json:"-"`

	ModelTimeStamps
}

func (u *User) Save() (string /* token */, *errors.RestErr) {

	//email validates in request

	// Check that user with this email exists
	var user User
	if !database.DB.Scopes(user.ScopeEmail(u.Email)).Find(&user).RecordNotFound() {
		return "", errors.NewBadRequestErr("email in use")
	}

	u.Password = encrypting.GetHashedPassword(u.Password)
	u.Status = StatusActive

	if err := database.DB.Create(&u).Error; err != nil {
		logger.Error("saving user", err)
		return "", errors.NewInternalServerError(err.Error())
	}

	token, err := OauthAccessToken{}.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) FindWithId() *errors.RestErr {

	if database.DB.Find(&u, u.ID).RecordNotFound() {
		return errors.NewNotFoundRequestErr("user not found")
	}
	return nil

}

func (u User) Login(email string, password string) (string /* token */, *errors.RestErr) {

	if database.DB.Scopes(u.ScopeEmail(email)).Find(&u).RecordNotFound() {
		return "", errors.NewNotFoundRequestErr("کاربری با این ایمیل وجود ندارد.")
	}

	if u.Status != StatusActive {
		return "", errors.NewBadRequestErr("your register hasn't complete yet")
	}

	if !encrypting.CheckPassword(u.Password, password) {
		return "", errors.NewBadRequestErr("wrong email or password")
	}

	token, err := OauthAccessToken{}.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

//scopes
func (u *User) ScopeEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("Email = ?", email)
	}
}

func (u *User) ScopeStatusActive(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", StatusActive)
}
