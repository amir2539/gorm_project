package group

type GroupCreateRequest struct {
	Name string `json:"name" binding:"required,min=3,max=20"`
}
