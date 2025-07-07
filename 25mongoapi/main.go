package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ffaiyaz23/golangLco/mongoapi/router"
)

func main() {
	fmt.Println("mongoDB Api")

	r := router.Router()

	fmt.Println("Server is getting is started ...")
	log.Fatal(http.ListenAndServe(":3000", r))

	fmt.Println("Listening to port 3000")
}
