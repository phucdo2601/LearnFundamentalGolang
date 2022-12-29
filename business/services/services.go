package services

import (
	"encoding/json"
	entities "entities"
	"fmt"
	"io/ioutil"
	"net/http"
	repositories "repositories"
)

func CreateDatabaseSchema(w http.ResponseWriter, r *http.Request) {
	fmt.Println("service CreateDatabaseSchema")
	repositories.SchemaMigrations()
	w.WriteHeader(http.StatusOK)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("service CreateNewUser")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user entities.User
	json.Unmarshal(reqBody, &user)
	user = repositories.CreateNewUser(user, user.Password, false)
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("services update user")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user entities.User
	json.Unmarshal(reqBody, &user)
	repositories.UpdateUser(user)
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("services get all users")
	var users []entities.User
	users = repositories.GetAllUsers()
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
