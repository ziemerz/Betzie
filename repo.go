package main

import (
	"database/sql"
	"fmt"
	"github.com/elithrar/simple-scrypt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Give us some seed data
func init() {
	//RepoCreateTodo(Todo{Name: "Write presentation"})
	//RepoCreateTodo(Todo{Name: "Host meetup"})
	db, _ = sql.Open("mysql", "root:@/betzie")
}

func AuthenticateLogin(u User) (User, error) {
	var user User
	var hashedPW []byte

	stmt, err := db.Prepare("SELECT id, username, password, role_id, first_name, last_name FROM users WHERE username=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res := stmt.QueryRow(u.Username)
	err = res.Scan(&user.Id, &user.Username, &hashedPW, &user.Role, &user.FirstName, &user.LastName)
	if err != nil {
		panic(err.Error())
	}

	if scrypt.CompareHashAndPassword([]byte(hashedPW), []byte(u.Password)) != nil {
		return User{}, fmt.Errorf("Something went wrong with the login, try again!")
	} else {
		fmt.Println("The user was found with the following information:")
		fmt.Println("Username: " + user.Username)
		fmt.Printf("Id: %d", user.Id)
		return user, nil
	}
}

func InsertUser(u User) error {
	stmt, err := db.Prepare("INSERT INTO users (username, password, role_id, first_name, last_name) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	hashedPW, err := scrypt.GenerateFromPassword([]byte(u.Password), scrypt.DefaultParams)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(u.Username, hashedPW, u.Role, u.FirstName, u.LastName)
	if err != nil {
		panic(err)
	} else {
		return nil
	}
	return fmt.Errorf("Could not create user")
}
