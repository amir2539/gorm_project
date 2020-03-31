package models

import (
	"gorm-learning/database"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/helper"
	"gorm-learning/utils/logger"
	"strconv"
)

type Group struct {
	ModelID
	Name  string `gorm:"type:varchar(100); not null"`
	Token string `gorm:"type:varchar(25); not null"`

	//group creator
	User   User `gorm:"foreignkey:user_id;association_foreignkey:id"` // use UserRefer as foreign key
	UserID uint

	TotalCost int `gorm:"default:0"`
	// many 2 many relation
	Users []User

	ModelTimeStamps
}

func (group *Group) Create() (string, *errors.RestErr) {
	group.Token = strconv.Itoa(helper.RandNumber(100000, 999999))

	if err := database.DB.Create(&group).Error; err != nil {
		logger.Error("create group error", err)
		return "", errors.NewInternalServerError("database error")
	}

	var userGroup UserGroup
	userGroup.UserID = group.UserID
	userGroup.GroupID = group.ID

	if err := userGroup.AddUser(); err != nil {
		return "", err
	}

	return group.Token, nil
}
