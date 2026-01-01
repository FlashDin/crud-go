package dto

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
