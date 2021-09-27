package mocks

import "github.com/garuda/go/crud/models"

var Customers = []models.Customer{
	{
		Id:       1,
		Name:     "thitipan",
		Lname:    "kijcharounpol",
		Password: "password",
		Email:    "thitipan991@gmail.com",
		Img:      "cat3.jpg",
	},
	{
		Id:       2,
		Name:     "kkddna",
		Lname:    "aeiou",
		Password: "123456",
		Email:    "ddkkiikk@gmail.com",
		Img:      "cat2.jpg",
	},
	{
		Id:       3,
		Name:     "ddeeff",
		Lname:    "aabbcc",
		Password: "98765",
		Email:    "bbbaaa@gmail.com",
		Img:      "cat.jpg",
	},
}
