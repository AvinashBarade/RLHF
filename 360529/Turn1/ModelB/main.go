
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// Claims for JWT tokens
type Claims struct {
	Tenant string `json:"tenant"`
	jwt.StandardClaims
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize a new router
	r := chi.NewRouter()

	// Register middlewares
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Mount the API handler
	r.Mount("/api/v1", apiHandler())

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func apiHandler() http.Handler {
	// Initialize a new JWT auth handler with a secret key
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	auth := jwtauth.New("HS256", []byte(secretKey), nil)

	// Create a new router for the API endpoints
	r := chi.NewRouter()

	// Register middlewares for the API router
	r.Use(auth.Verifier())
	r.Use(auth.Authenticator)

	// API endpoints
	r.Get("/profile", profileHandler)
	r.Post("/login", loginHandler(auth))

	return r
}

func loginHandler(auth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// In a real application, this would be where you validate user credentials
		// and retrieve the tenant associated with the user.
		tenant := r.FormValue("tenant")
		if tenant == "" {
			render.Render(w, r, ErrInvalidRequest("tenant is required"))
			return
		}

		// Generate a new JWT token with the claims
		claims := Claims{
			Tenant: tenant,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 1609459200, // Set expiration to 2021-01-01
			},
		}
		token, err := auth.Encode(&claims)
		if err != nil {
			render.Render(w, r, ErrInternalServerError(err))
			return
		}

		// Respond with the JWT token
		render.JSON(w, r, map[string]string{"token": token})
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {