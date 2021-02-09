package models

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dgrijalva/jwt-go"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
)

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"

// AuthRepository ...
type AuthRepository interface {
	CreateToken(userID string) (*forms.TokenDetails, error)
	CreateAuth(userid string, td *forms.TokenDetails) error
	TokenValidFromRepo(r *http.Request) error
	FetchAuth(authD *forms.AccessDetails) (string, error)
	DeleteAuth(givenUUID string) (*mongo.DeleteResult, error)
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, error)
	ExtractTokenMetadata(r *http.Request) (*forms.AccessDetails, error)
	// FetchAuth(authD *AccessDetails) (string, error)
	// DeleteAuth(givenUUID string) (int64, error)
}
