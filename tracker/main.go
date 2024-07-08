package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request")
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Printf(string(body))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a GET request")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", postHandler).Methods("POST")
	router.HandleFunc("/", getHandler).Methods("GET")
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
