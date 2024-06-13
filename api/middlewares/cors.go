package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

const (
	maxAge = 300
)

func CorsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// Define a custom CORS policy function
		customCORS := func(r *http.Request) cors.Options {
			// Default CORS options for other routes
			return cors.Options{
				AllowedOrigins: []string{"*"},
				MaxAge:         maxAge,
			}
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Determine CORS options based on the request
			options := customCORS(r)
			// Initialize the CORS handler with the determined options
			cors := cors.New(options)
			corsHandler := cors.Handler(next)

			// Serve the request with the CORS handler
			corsHandler.ServeHTTP(w, r)
		})
	}
}
