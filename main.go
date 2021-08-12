package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shishanksingh2015/GoJWT/driver"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	db = driver.ConnectDb()
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
