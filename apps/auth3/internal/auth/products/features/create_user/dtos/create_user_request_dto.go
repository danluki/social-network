package dtos

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
