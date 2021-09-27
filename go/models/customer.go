package models

type Customer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lname    string `json:"lname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Img      string `json:"img"`
}
