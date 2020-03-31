package parser

import (
	"gorm-learning/utils/errors"
	"strings"
)

func ReadBearerToken(tokenHeader string) (string, *errors.RestErr) {
	if !strings.Contains(tokenHeader, "Bearer ") {
		return "", errors.NewUnauthorizedError()
	}

	token := strings.Replace(tokenHeader, "Bearer", "", 1)
	token = strings.TrimSpace(token)

	return token, nil
}
