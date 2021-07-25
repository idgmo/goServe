package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/idgmo/goServe/tree/main/internal/db"
)

func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed")
	}

	_, err = database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("Something has been called. (POST)")
	fmt.Println("Something has been called. (POST)")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))

}
