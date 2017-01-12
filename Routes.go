package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"CreateUser",
		"POST",
		"/register",
		CreateUser,
	},
	Route{
		"CreateCoupon",
		"POST",
		"/coupon",
		CreateCoupon,
	},
	Route{
		"GetCoupon",
		"GET",
		"/coupon/{couponId}",
		GetCoupon,
	},
}
