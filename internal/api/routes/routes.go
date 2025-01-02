package routes

import (
	"github.com/gorilla/mux"
	"urlshortener/internal/api/handlers"
)

// SetupRoutes initializes the API routes for URL handling
func SetupRoutes(r *mux.Router, urlHandler *handlers.URLHandler) {
	// Route for creating a new short URL
	r.HandleFunc("/shorten", urlHandler.CreateShortURL).Methods("POST")

	// Route for retrieving a URL by its short code
	r.HandleFunc("/shorten/{shortCode}", urlHandler.GetURL).Methods("GET")

	// Route for updating an existing short URL
	r.HandleFunc("/shorten/{shortCode}", urlHandler.UpdateURL).Methods("PUT")

	// Route for deleting a URL by its short code
	r.HandleFunc("/shorten/{shortCode}", urlHandler.DeleteURL).Methods("DELETE")

	// Route for retrieving statistics for a URL by its short code
	r.HandleFunc("/shorten/{shortCode}/stats", urlHandler.GetStats).Methods("GET")
}
