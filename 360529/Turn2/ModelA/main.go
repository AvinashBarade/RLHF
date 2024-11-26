package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// Tenant represents a tenant in the application
type Tenant struct {
	ID     string
	Secret string
	Name   string
}

// GetTenantSecret returns the secret key for the given tenant ID
func GetTenantSecret(tenantID string) (string, error) {
	switch tenantID {
	case "tenant1":
		return os.Getenv("TENANT1_SECRET"), nil
	case "tenant2":
		return os.Getenv("TENANT2_SECRET"), nil
	default:
		return "", fmt.Errorf("tenant ID %s not found", tenantID)
	}
}

// CreateJWT creates a new JWT for a given tenant
func CreateJWT(tenantID string, claims map[string]interface{}) (string, error) {
	secret, err := GetTenantSecret(tenantID)
	if err != nil {
		return "", err
	}

	signingKey := []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyJWT verifies a JWT for a given tenant
func VerifyJWT(tenantID string, tokenString string) (bool, error) {
	secret, err := GetTenantSecret(tenantID)
	if err != nil {
		return false, err
	}

	signingKey := []byte(secret)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

// Middleware to authenticate requests
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := mux.Vars(r)["tenantID"]
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Token required", http.StatusUnauthorized)
			return
		}

		valid, err := VerifyJWT(tenantID, tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Example endpoint to generate a JWT
	r.HandleFunc("/auth/{tenantID}/jwt", func(w http.ResponseWriter, r *http.Request) {
		tenantID := mux.Vars(r)["tenantID"]
		claims := &jwt.StandardClaims{
			Subject:   "user123",
			Issuer:    tenantID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		}

		jwtString, err := CreateJWT(tenantID, claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": jwtString})
	})

	// Example protected endpoint
	r.HandleFunc("/protected/{tenantID}", AuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := mux.Vars(r)["tenantID"]
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Access granted to " + tenantID))
	})))

	http.ListenAndServe(":8080", r)
}
