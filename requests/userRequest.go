package requests


type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
}
