package migration

import (
	"gorm-learning/database"
	"gorm-learning/http/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Group{})
	database.DB.AutoMigrate(&models.UserGroup{})
	database.DB.AutoMigrate(&models.OauthAccessToken{})

	// add foreign keys

	// groups table
	database.DB.Model(&models.Group{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	//user_group table
	database.DB.Model(&models.UserGroup{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	database.DB.Model(&models.UserGroup{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")

	//oath_access_tokens table
	database.DB.Model(&models.OauthAccessToken{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

}
