package main

import (
	db2 "Fintech/db"
	"Fintech/handler"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/login", handler.LogIn)
	http.HandleFunc("/signup", handler.SignUp)
	http.HandleFunc("/reset", handler.ResetPassword)
	// initialize our database connection
	db2.GetDB()
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8080", nil))
}

