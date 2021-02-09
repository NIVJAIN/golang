package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file, please create one in the root directory")
// 	}
// }

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userID string) (string, error) {
	// Define token expiration time
	fmt.Println("jwtket::==>>", jwtKey)
	// expirationTime := time.Now().Add(1440 * time.Minute)
	expirationTime := time.Now().Add(1 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func GenerateRefreshToken(userID string) (string, error) {
	// Define token expiration time
	// expirationTime := time.Now().Add(1440 * time.Minute)
	expirationTime := time.Now().Add(1 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("niltokenerror", err)
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		fmt.Println("tkn.valud", err)
		return "", err
	}
	fmt.Println("clams.valud", claims.UserID)
	return claims.UserID, nil
}
