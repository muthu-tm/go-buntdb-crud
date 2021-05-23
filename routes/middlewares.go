package routes

import (
	"log"
	"net/http"
)

// SetMiddlewareJSON - a simple middleware
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// log the requested URL
		log.Println(r.RequestURI)
		w.Header().Set("Content-Type", "application/json")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next(w, r)
	}
}
