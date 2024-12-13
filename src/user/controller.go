package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Reusable components => structs
//userservice => process
//user => data
//I want structs for both processes and data
// processes are used through references
// data are used by values
// contoller is a process by itself, it implement a logic
// input => controller => service

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service}
}

func (controller *UserController) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/users", controller.Get).Methods("GET")
	r.HandleFunc("/api/users/{id}", controller.GetByID).Methods("GET")
}

func (controller *UserController) Get(w http.ResponseWriter, req *http.Request) {

	users, err := controller.service.Read()
	fmt.Println(users)
	if err != nil {
		http.Error(w, "Error fetching all the users! ", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		http.Error(w, "Error marshalling all users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (controller *UserController) GetByID(w http.ResponseWriter, req *http.Request) {
	// Extract the user ID from the query parameter
	vars := mux.Vars(req)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	idi, err := strconv.Atoi(id) // Convert string to integer

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	// Call the service to read the user by ID
	user, err := controller.service.ReadById(idi)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
