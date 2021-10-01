package entities

type RegisterPackage struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
