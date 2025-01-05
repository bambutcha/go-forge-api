package model

type User struct {
	ID                int    `json:"id"`
	Email 		      string `json:"username"`
	EncryptedPassword string `json:"password"`
}
