package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sago-code/user-api-go/controllers"
	"github.com/sago-code/user-api-go/db"
	"github.com/sago-code/user-api-go/models"
	"github.com/sago-code/user-api-go/routes"
)

func main() {

	//Database
	db.DBConnection()

	//Migrate the schema
	db.DB.AutoMigrate(models.User{})

	//Gorilla Mux
	r := mux.NewRouter()

	//Routes
	r.HandleFunc("/", controllers.HomeHandler)
	routes.UserRoutes(r)

	//Server
	http.ListenAndServe(":3000", r)
}
