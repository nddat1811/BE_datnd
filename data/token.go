package data

type Token struct {
	Token string `json:"token" binding:"required"`
}