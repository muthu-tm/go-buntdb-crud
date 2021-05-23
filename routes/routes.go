package routes

import (
	"github.com/gorilla/mux"

	c "go-inmem-crud/controllers"
)

// InitializeRoutes - Defines the mapping for all routes
func InitializeRoutes(router *mux.Router) {

	//Employees routes
	router.HandleFunc("/emp", SetMiddlewareJSON(c.CreateEmp)).Methods("POST")
	router.HandleFunc("/emp", SetMiddlewareJSON(c.GetAllEmp)).Methods("GET")
	router.HandleFunc("/emp/{id}", SetMiddlewareJSON(c.GetEmp)).Methods("GET")
	router.HandleFunc("/emp/{id}", SetMiddlewareJSON(c.UpdateEmp)).Methods("PUT")
	router.HandleFunc("/emp/{id}", SetMiddlewareJSON(c.DeleteEmp)).Methods("DELETE")
}
