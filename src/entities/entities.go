package entities

type User struct {
	Id            int
	Username      string `json:"username"`
	Password_hash string `json:"password_hash"`
	Email         string `json:"email"`
}
