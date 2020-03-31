package models

import (
	"github.com/jinzhu/gorm"
	"gorm-learning/database"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/logger"
)

type UserGroup struct {
	ModelID

	User   User `gorm:"foreignkey:UserId"` // use UserRefer as foreign key
	UserID uint

	Group   Group `gorm:"foreignkey:GroupID"` // use UserRefer as foreign key
	GroupID uint

	ModelTimeStamps
}

func (UserGroup) TableName() string {
	return "user_group"
}

func (userGroup *UserGroup) AddUser() *errors.RestErr {
	//check user exists in group
	if database.DB.
		Scopes(userGroup.ScopeGroupId(userGroup.GroupID), userGroup.ScopeUserId(userGroup.UserID)).
		RecordNotFound() {
		return errors.NewInternalServerError("user already exists in group")
	}

	if err := database.DB.Create(&userGroup).Error; err != nil {
		logger.Error("saving user_group", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

//scopes
func (userGroup *UserGroup) ScopeGroupId(groupId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("group_id = ?", groupId)
	}
}

func (userGroup *UserGroup) ScopeUserId(userId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}
