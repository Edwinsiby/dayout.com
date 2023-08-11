package models

type User struct {
	Role     string
	Name     string
	Email    string
	Password string
}
type Invalid struct {
	Errpass string
	Errmail string `json:"email"`
}
