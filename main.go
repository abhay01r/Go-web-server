package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my server")
}

func main() {
	http.HandleFunc("/", home)                             // this tells the server whenever someone visits / (the homepage), run the home function / means root url (http://localhost:8080/).
	fmt.Println("Server running on http://localhost:8080") // this will be printed in terminal for us to know that server is running
	http.ListenAndServe(":8080", nil)                      // starts the server on port 8080
}
