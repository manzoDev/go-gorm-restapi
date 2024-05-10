package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manzoDev/go-gorm-restapi/db"
	"github.com/manzoDev/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	//w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	json.NewEncoder(w).Encode(params)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
	//w.Write([]byte("post"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
