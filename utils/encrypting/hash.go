package encrypting

import (
	"golang.org/x/crypto/bcrypt"
	"gorm-learning/utils/logger"
)

func GetHashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		logger.Error("error in encrypting", err)
		panic(err.Error())
	}

	return string(hash)

}

func CheckPassword(hashed string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}
