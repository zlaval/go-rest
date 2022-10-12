package main

import (
	"encoding/json"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users
	json.NewDecoder(r.Body).Decode(&user)
	Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []Users
	Instance.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
