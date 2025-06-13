package routes

import (
	"github.com/gorilla/mux"
	"github.com/sago-code/user-api-go/controllers"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.PutUsersHandler).Methods("PUT")
	r.HandleFunc("/users/sofd/{id}", controllers.SofDeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")
}
