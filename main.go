package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"gorm-learning/database"
	"gorm-learning/database/migration"
	"gorm-learning/routes"
	"os"
)

var (
	DB *gorm.DB
)

func main() {

	DB = database.Connect()
	migration.Migrate()

	// close database
	defer DB.Close()

	router := routes.InitRoutes()

	router.Run(":" + os.Getenv("PORT"))

}
