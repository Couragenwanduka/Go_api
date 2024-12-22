package userController 

import (
	"api/service"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

func CreateUser(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user Userservice.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || user.Address == "" || user.Phone == "" {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	userExists, err := Userservice.FindUserByEmail(user.Email, client)
	if err != nil {
		// Handle other errors
		http.Error(w, "Error checking email", http.StatusInternalServerError)
		fmt.Println("Error finding user:", err)
		return
	}
	if userExists != nil {
		// User already exists
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	if _, err := Userservice.SaveUser(user, client); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

    w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

