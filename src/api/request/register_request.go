package request

type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
