package main

import (
	"main/service"
	"net/http"
	"fmt"
)


// func hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Hello World")
// 	fmt.Fprint(w,"Hello World")
// }
var UserSvc = service.NewUserService()

// type UserRegister struct {
// 	name 	string		`json:"name"`
// }

func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":	
			UserSvc.GetUser(w)	
		case "POST":
			UserSvc.Register(r, &w)
			// UserSvc.Register(&service.User{Nama: name}, r, &w)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	
}

const (
	port = "8090"
)

func main() {

	// userSvc.Register(&service.User{Nama: "budi"})
	// userSvc.Register(&service.User{Nama: "ke"})
	// userSvc.Register(&service.User{Nama: "vin"})
	// userSvc.Register(&service.User{Nama: "hu"})
	// userSvc.Register(&service.User{Nama: "go"})
	// userSvc.GetUser()

	// http.HandleFunc("/hello", hello)

	http.HandleFunc("/user", user)

	fmt.Println("=================Running at port " + port + "=================")
	http.ListenAndServe(":" + port, nil)
}