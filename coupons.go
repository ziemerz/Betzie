package main

import ()

type Coupon struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	BetAmount      float64 `json:"betamount"`
	ExpectedReturn float64 `json:"expectedreturn"`
	ActualReturn   float64 `json:"actualreturn"`
	FailSuccess    int     `json:"failsuccess"`
	Bets           Bets    `json:"bets"`
}

type Bet struct {
	SportId    int     `json:"sportid"`
	Sport      string  `json:"sport"`
	Team1      string  `json:team1"`
	Team2      string  `json:team2"`
	Type       string  `json:"type"`
	Multiplier float64 `json:"multiplier"`
}

type Bets []Bet
