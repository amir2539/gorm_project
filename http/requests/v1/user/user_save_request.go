package user

type UserSaveRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" binding:"required,min=6"`
}
