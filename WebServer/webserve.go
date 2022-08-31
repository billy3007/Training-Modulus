package main

import (
	"fmt"
	"net/http"
	"training/Training-Modulus/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)

	http.HandleFunc("/register", userSvc.RegisterHandler)
	http.HandleFunc("/user", userSvc.GetUserHandler)
	fmt.Println("Starting localhost:3030")
	http.ListenAndServe("localhost:3030", nil)
}
