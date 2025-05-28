package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

// Global users slice to hold data temporarily
var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte("Here are the users"))
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("User created"))
	// decode request body to user struct
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}
	// users = append(users, u)
	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}

	for _, user := range users {
		if user.Email == u.Email {
			return errors.New("email already exists")
		}
	}
	users = append(users, u)
	return nil
}

// func main(){
// 	api :=&api{addr:":8081"}
// }
