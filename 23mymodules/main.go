package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// go mod init github.com/ffaiyaz23/golangLco
// go get -u github.com/gorilla/mux   ->  to download certain modules from web
// go mod verify ->   to verify all modules
// go list -m all   ->  my code is dependednt on which modules/files
// go list -m -versions githubcom/gorilla/mux   ->   to check all the versions of the modules (maybe latest is not working)

func main() {
	fmt.Println("Hello mod in Golang")

	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcomme to golang series</h1>"))
}
