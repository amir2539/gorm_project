package parser

import (
	"encoding/json"
	"gorm-learning/utils/errors"
	"gorm-learning/utils/logger"
)

func SwapJson(out interface{}, in interface{}) *errors.RestErr {

	marshalledJson, err := json.Marshal(in)
	if err != nil {
		logger.Error(err.Error(), err)
		return errors.NewInternalServerError(err.Error())
	}

	json.Unmarshal(marshalledJson, &out)
	return nil
}
