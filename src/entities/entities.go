package entities

type User struct {
	Id            int
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password" binding:"required"`
	Email         string `json:"email" binding:"required"`
}
