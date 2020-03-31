package user

import (
	"gorm-learning/http/models"
	"gorm-learning/http/resources"
)

type UserShowResource struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *UserShowResource) Get(user models.User) *UserShowResource {
	r.Name = user.Name
	r.Email = user.Email

	return r
}

func (r *UserShowResource) Append(success bool, message string) interface{} {
	return resources.BaseResource{
		Success: success,
		Message: message,
		Data:    r,
	}
}
