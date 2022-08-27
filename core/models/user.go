package models

type User struct {
	Id       string `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type AuthInput struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}
