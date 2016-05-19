package main

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Users []User
