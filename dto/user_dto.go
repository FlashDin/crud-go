package dto

type UserRequest struct {
	Name string `json:"name" binding:"required,min=2"`
	Age  int    `json:"age" binding:"required,min=1"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
