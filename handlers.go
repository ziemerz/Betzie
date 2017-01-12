package main

import (
	. "Betzie/ctypes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var rep Repo

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	HttpResponse(w, r, http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	/*if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}*/
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1049576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	JsonResponse(w, r, body, &todo)

	HttpResponse(w, r, http.StatusOK)
	/*if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}*/
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1049576))
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	user, err = rep.AuthenticateLogin(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hej" + user.Username)

	HttpResponse(w, r, http.StatusOK)
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1049576))
	defer r.Body.Close()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	err = rep.InsertUser(user)
	if err != nil {
		panic(err)
	}
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon Coupon
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 104976))
	defer r.Body.Close()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &coupon)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted with success")
	/*err = InsertCoupon(coupon)
	if err != nil {
		panic(err)
	}*/
	HttpResponse(w, r, http.StatusCreated)

	JsonResponse(w, r, body, &coupon)

	if err := json.NewEncoder(w).Encode(coupon); err != nil {
		panic(err)
	}
}

func GetCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon Coupon
	vars := mux.Vars(r)
	couponId, err := strconv.Atoi(vars["couponId"])
	if err != nil {
		panic(err)
	}

	fmt.Println("Getting coupon")
	coupon, err = rep.GetCoupon(couponId)
	if err != nil {
		panic(err)
	}
	jcoupon, err := WrapJSON(coupon, true)
	if err != nil {
		panic(err)
	}
	HttpResponse(w, r, http.StatusOK)
	//if err := json.NewEncoder(w).Encode(jcoupon); err != nil {
	//	panic(err)
	//}
	w.Write(jcoupon)
}
