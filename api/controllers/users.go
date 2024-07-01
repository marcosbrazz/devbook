package controllers

import (
	"api/dao"
	"api/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := dao.Connect()
	if err != nil {
		log.Fatal(err)
	}

	dao := dao.CreateUserDao(db)

	id, err := dao.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(id))

}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
