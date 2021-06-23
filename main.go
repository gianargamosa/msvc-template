package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "stellar microservice is up and running!\n")
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router()
	// port := os.Getenv("PORT")

	fmt.Println("server started and listening")

	log.Fatal(http.ListenAndServe(":8080", router))
}
