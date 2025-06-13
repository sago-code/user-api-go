package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sago-code/user-api-go/db"
	"github.com/sago-code/user-api-go/models"
	"github.com/sago-code/user-api-go/strategies"
)

// Maneja la solicitud GET para obtener todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

// Maneja la solicitud GET para obtener un usuario por su ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	//manejar el caso de que el usuario no exista
	if err := db.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// Maneja la solicitud POST para crear un nuevo usuario
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

// Maneja la solicitud PUT para actualizar un usuario
func PutUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User

	//manejar el caso de que el usuario no exista
	if err := db.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	//maneja el caso de que el body sea nulo o vacio o invalido
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//actualiza los campos del usuario segun lo que se envie en el body
	if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// Maneja la solicitud DELETE para eliminar un usuario en soft delete
func SofDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	//manejar el caso de que el id sea invalido
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	//manejar el caso de que el usuario no exista
	if err := db.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	//llama al contexto de eliminacion con la estrategia de soft delete
	deleteContext := strategies.NewDeleteContext(db.DB, &strategies.SoftDeleteStrategy{})

	//llama al metodo de eliminacion del contexto y maneja el error
	if err := deleteContext.DeleteUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	//manejar el caso de que el id sea invalido
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	//manejar el caso de que el usuario no exista
	if err := db.DB.Unscoped().First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	//llama al contexto de eliminacion con la estrategia de hard delete
	deleteContext := strategies.NewDeleteContext(db.DB, &strategies.HardDeleteStrategy{})

	//llama al metodo de eliminacion del contexto y maneja el error
	if err := deleteContext.DeleteUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
