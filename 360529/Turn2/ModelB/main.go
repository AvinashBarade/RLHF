package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/golang-jwt/jwt/v4"
)

// Tenant represents a tenant in the application
type Tenant struct {
	ID      string
	Secret  string
	Name    string
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
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

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return false, err
	}

	// Validate the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the tenant ID matches the claim
		if claims["tenantID"] != tenantID {
			return false, fmt.Errorf("invalid tenant ID in claim")
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid token claims")
}

func main() {
	r := mux.NewRouter()

	// Example endpoint to generate a JWT
	r.HandleFunc("/auth/{tenantID}/jwt", func(w http.ResponseWriter, r *http.Request) {
		tenantID := mux.Vars(r)["tenantID"]
		claims := map[string]interface{}{
			"tenantID": tenantID,
			"userID":   "user123",
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		}

		jwtString, err := CreateJWT(tenantID, claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return