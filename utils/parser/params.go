package parser

import (
	"gorm-learning/utils/errors"
	"strconv"
)

func GetIntegerParam(param string) (uint, *errors.RestErr) {

	res, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestErr("invalid param")
	}

	return uint(res), nil
}
