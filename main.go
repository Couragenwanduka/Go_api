package main

import (
	"api/db"
	"api/controller"
	"net/http"
	"fmt"
)

func main() {
	client := db.Connect()
	
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userController.CreateUser(w, r, client)
	})

	port:= ":8000"

	fmt.Printf("Server running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	
}