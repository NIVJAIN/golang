package models

import (
	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository ...
type UserRepository interface {
	FindByEmailID(EMAILID string) (*forms.UserDetails, error)
	GetUserByEmail(email string) (*forms.UserDetails, error) // During jwt login verifification
	CreateToken(userID string) (*forms.TokenDetails, error)
	RegisterUser(user *forms.UserDetails) (*mongo.InsertOneResult, error)
	Login(user *forms.LoginUserCommand) (*forms.Token, error)
	Logout(c *gin.Context) error
	NotExistsSavePeopleInMongoDB(peope *forms.PersonDetails) (*mongo.InsertOneResult, error)
	UpdateFieldValueInMongoDB(fvn *forms.FieldNameValue) (*mongo.UpdateResult, error)
	UpdateFieldValueInMongoDBUsingMongoID(fvn *forms.MongoidStruct) (*mongo.UpdateResult, error)
	UpdateFieldValueInMongoDBUsingEmailID(fvn *forms.EmailStruct) (*mongo.UpdateResult, error)
	PostUpdateFieldInMongoDBviaEmailID(fvn *forms.EmailStruct) (*mongo.UpdateResult, error)
	QueryDatesMongoDB(fnv *forms.QueryDates) ([]*forms.UserDetails, error)
}

// // UserRepository ..
// type UserRepository interface {
// 	// FindByID(ID int) (*User, error)
// 	// Save(user *User) error
// 	FindByEmailID(EMAILID string) (*UserDetails, error)
// 	SaveInMongoDB(user *UserDetails) (*mongo.InsertOneResult, error)
// }
