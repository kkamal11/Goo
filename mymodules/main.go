package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Modules")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}


func greeter(){
	fmt.Println("Hello go mod users")
}


func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello there</h1>"))
	databyte, err := fmt.Fprintf(w, "Welcome to the Home Page!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Endpoint Hit: homePage ", databyte, " bytes written")
}