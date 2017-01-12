package main

import (
	"database/sql"
	"fmt"
	"github.com/elithrar/simple-scrypt"
	_ "github.com/go-sql-driver/mysql"
)

var repo Repo

type Repo struct {
	db *sql.DB
}

func init() {
	//RepoCreateTodo(Todo{Name: "Write presentation"})
	//RepoCreateTodo(Todo{Name: "Host meetup"})
	//db, _ = sql.Open("mysql", "root:@/betzie")
	repo.db, _ = sql.Open("mysql", "root:@/betzie")
}

func (rep Repo) AuthenticateLogin(u User) (User, error) {
	var user User
	var hashedPW []byte

	stmt, err := repo.db.Prepare("SELECT id, username, password, role_id, first_name, last_name FROM users WHERE username=?")
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

func (rep Repo) InsertUser(u User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (username, password, role_id, first_name, last_name) VALUES (?,?,?,?,?)")
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

func (rep Repo) GetCoupon(id int) (Coupon, error) {
	var coupon Coupon
	bets := make([]Bet, 0, 20)
	var sql string
	sql = "SELECT coupon_id, custom_name, bet_amount, expected_return, actual_return, fail_success FROM coupons WHERE coupon_id = ?"
	res, err := repo.db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}
	if !res.Next() {
		return Coupon{}, fmt.Errorf("Could not find coupon")
	}
	err = res.Scan(&coupon.Id, &coupon.Name, &coupon.BetAmount, &coupon.ExpectedReturn, &coupon.ActualReturn, &coupon.FailSuccess)

	if err != nil {
		panic(err.Error())
	}

	sql = "SELECT bets.team1, bets.team2, bets.type, sports.sport_name FROM coupons_to_bets INNER JOIN bets ON coupons_to_bets.bet_id_fk = bets.bet_id INNER JOIN sports ON sports.sport_id = bets.sport_id_fk WHERE coupons_to_bets.coupons_id_fk = ?"
	res, err = repo.db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}
	for res.Next() {
		var bet Bet
		res.Scan(&bet.Team1, &bet.Team2, &bet.Type, &bet.Sport)
		bets = append(bets, bet)
	}
	coupon.Bets = bets
	return coupon, nil
}
