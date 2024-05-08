package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}
