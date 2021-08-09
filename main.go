package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndPoint)).Methods("GET")

	log.Println("Listen on port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signed up success"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked")
}
func protectedEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProtectedEndPoint invoked")
}
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked")
	return nil
}
