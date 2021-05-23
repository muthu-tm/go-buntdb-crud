package main

import (
	"fmt"
	"go-inmem-crud/db"
	"go-inmem-crud/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("unable to load .env file. WARNING!")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting environment values, %v", err)
	} else {
		fmt.Println("We are getting the environment values")
	}

	port := os.Getenv("PORT")

	db.InitBuntDB()

	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	fmt.Println("Listening to port : ", port)
	err = http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:PORT
	if err != nil {
		fmt.Print(err)
	}

	defer db.Close()
}
