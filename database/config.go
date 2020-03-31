package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	DB *gorm.DB
)

func Connect() *gorm.DB {

	connection, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(127.0.0.1:3306)/%s?charset=utf8&parseTime=true",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_NAME")),
	)
	if err != nil {
		panic(err.Error())
	}

	DB = connection
	return connection
}
