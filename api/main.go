package main

import (
	"api/login"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	const PORT = "8080"
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/log")
	r.HandleFunc("/signup", loginPost).Methods("POST")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
func loginPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	ans := login.Login(email, pass)
	if ans == true {
		fmt.Println("TRUE")
	}
}
