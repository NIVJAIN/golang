package repositories

import (
	"context"
	"errors"

	// "log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/alarm"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
)

// var log *logrus.Logger

// RegisterUser ...
func (r *MongoClient) RegisterUser(user *forms.UserDetails) (*mongo.InsertOneResult, error) {
	findUserEmailExist, _ := r.FindByEmailID(user.Email)
	if findUserEmailExist != nil {
		err1 := errors.New("User Email already exists")
		return nil, err1
	}
	user.CreatedAt = time.Now()
	user.NanoUnix = time.Now().UnixNano()
	user.Password = helpers.GeneratePasswordHash([]byte(user.Password))
	insertResult, err := r.collection["users"].InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

type a struct {
	N string
	A int
}

func (r *MongoClient) loga(msg ...interface{}) {
	r.logcollections["info"].Info(msg)
}

// Login ...
func (r *MongoClient) Login(usersLoginData *forms.LoginUserCommand) (*forms.Token, error) {
	result, err := r.GetUserByEmail(usersLoginData.Email)
	b := a{
		"j", 5,
	}
	r.logcollections["info"].Info(b)
	if err != nil {
		r.logcollections["info"].Error(err.Error())
		err = errors.New("Email or Password is incorrect")
		return nil, err
	}
	if result.Email == "" {
		err = errors.New("Problem logging into your account")
		return nil, err
	}
	// Get the hashed password from the saved document
	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(usersLoginData.Password)
	err = helpers.PasswordCompare(password, hashedPassword)
	if err != nil {
		log.Error(err.Error())
		err = errors.New("Invalid user credentials")
		return nil, err
	}
	//Generate the JWT auth token
	tokenDetails, err := r.CreateToken(usersLoginData.Email)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// log.Println("tokendetails::-->>", tokenDetails)
	err = r.CreateAuth(usersLoginData.Email, tokenDetails)
	if err != nil {
		log.Error(err.Error())
		err = errors.New("Unable to create Auth details")
		return nil, err
	}

	tokens := &forms.Token{
		tokenDetails.AccessToken,
		tokenDetails.RefreshToken,
	}
	return tokens, nil
}

// Logout ...
func (r *MongoClient) Logout(c *gin.Context) error {
	au, err := r.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.Error(err.Error())
		alarm.WeChat(err.Error())
		err = errors.New("User not logged in")

		return err

	}
	deleted, delErr := r.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == nil { //if any goes wrong
		log.Error(delErr)
		err = errors.New("Invalid request")
		return err
	}
	return nil
}

// FindByEmailID ...
func (r *MongoClient) FindByEmailID(EMAILID string) (*forms.UserDetails, error) {
	var userDetails forms.UserDetails
	filter := bson.D{{"email", EMAILID}}
	// collx := r.db.Database("test").Collection("jain")
	err := r.collection["users"].FindOne(context.TODO(), filter).Decode(&userDetails)
	if err != nil {
		// log.Println("Unabelt to get ", err)
		log.Error(err.Error())
		return nil, err
	}
	// fmt.Printf("Found a single document: %+v\n", result)
	return &userDetails, nil

}

// GetUserByEmail handles fetching user by email
func (r *MongoClient) GetUserByEmail(email string) (*forms.UserDetails, error) {
	var userDetails forms.UserDetails
	filter := bson.D{{"email", email}}
	// collx := r.db.Database("test").Collection("jain")
	err := r.collection["users"].FindOne(context.TODO(), filter).Decode(&userDetails)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// fmt.Printf("Found a single document: %+v\n", result)
	return &userDetails, nil
}

//NotExistsSavePeopleInMongoDB ...
func (r *MongoClient) NotExistsSavePeopleInMongoDB(people *forms.PersonDetails) (*mongo.InsertOneResult, error) {
	var peopleDetails forms.PersonDetails
	filter := bson.D{{"name", people.Name}}
	err := r.collection["people"].FindOne(context.TODO(), filter).Decode(&peopleDetails)
	if err == nil { // if err == nil means document exists
		// log.Error(err.Error())
		err1 := errors.New("math: User already exists")
		// log.Println("err1", err1)
		return nil, err1
	}
	insertResult, err := r.collection["people"].InsertOne(context.TODO(), people)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return insertResult, nil
}

// GET {{baseUrl}}/user/query?filtername=name&filtervalue=Jain5&fieldname=city&fieldvalue=Russia

//UpdateFieldValueInMongoDB ...
func (r *MongoClient) UpdateFieldValueInMongoDB(fnv *forms.FieldNameValue) (*mongo.UpdateResult, error) {
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.D{{fnv.FilterName, fnv.FilterValue}}
	// log.Println("UpdateFieldValueInMongoDB::>", fnv.FieldName, fnv.FieldValue, fnv.FilterName, fnv.FilterValue)
	update := bson.D{
		{"$set", bson.D{
			{fnv.FieldName, fnv.FieldValue},
		}},
	}
	collx := r.db.Database("test").Collection("people")
	result, err := collx.UpdateOne(context.TODO(), filter, update)
	// Check for error, else print the UpdateOne() API call results
	if err != nil {
		log.Error(err.Error())
		return nil, err
	} else {
		// log.Println("UpdateOne() result:", result)
		// log.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		// log.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		// log.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		// log.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		// log.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
		return result, nil
	}
}

// QueryDatesMongoDB ...
func (r *MongoClient) QueryDatesMongoDB(fnv *forms.QueryDates) ([]*forms.UserDetails, error) {
	// // Declare an _id filter to get a specific MongoDB document
	px, _ := time.Parse("2006-01-02", fnv.StartDate)
	py, _ := time.Parse("2006-01-02", fnv.ToDate)

	filter := bson.M{
		fnv.FilterName: bson.M{
			"$gt": px.UnixNano(),
			"$lt": py.UnixNano(),
		},
	}
	var results []*forms.UserDetails
	findOptions := options.Find()
	// findOptions.SetLimit(2)
	cur, err := r.collection["users"].Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Error(err.Error())

		return nil, err
	}
	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem forms.UserDetails
		err := cur.Decode(&elem)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// // fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	// for i, s := range results {
	// 	log.Println("MongoResults......=>", i, s.CreatedAt)
	// }
	return results, nil
}

// UpdateFieldValueInMongoDBUsingEmailID ...
func (r *MongoClient) UpdateFieldValueInMongoDBUsingEmailID(fnv *forms.EmailStruct) (*mongo.UpdateResult, error) {
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.D{{"email", fnv.EmailID}}
	update := bson.D{
		{"$set", bson.D{
			{fnv.FieldName, fnv.FieldValue},
		}},
	}
	collx := r.db.Database("test").Collection("people")
	result, err := collx.UpdateOne(context.TODO(), filter, update)
	// Check for error, else print the UpdateOne() API call results
	if err != nil {
		// log.Println("UpdateOne() result ERROR:", err)
		log.Error(err.Error())
		return nil, err
	} else {
		// log.Println("UpdateOne() result:", result)
		// log.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		// log.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		// log.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		// log.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		// log.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
		return result, nil
	}
}

//UpdateFieldValueInMongoDBUsingMongoID ...
func (r *MongoClient) UpdateFieldValueInMongoDBUsingMongoID(fnv *forms.MongoidStruct) (*mongo.UpdateResult, error) {
	// Declare an _id string and create an ObjectID
	docID := fnv.MongoID
	objID, err := primitive.ObjectIDFromHex(docID)
	// Check for MongoDB ID ObjectIDFromHex errors
	if err != nil {
		// log.Println("ObjectIDFromHex ERROR", err)
		log.Error(err.Error())
		return nil, err
	}
	filter := bson.D{{"_id", objID}}
	update := bson.D{
		{"$set", bson.D{
			{fnv.FieldName, fnv.FieldValue},
		}},
	}
	collx := r.db.Database("test").Collection("people")
	result, err := collx.UpdateOne(context.TODO(), filter, update)
	// Check for error, else print the UpdateOne() API call results
	if err != nil {
		log.Error(err.Error())
		return nil, err
	} else {
		// log.Println("UpdateOne() result:", result)
		// log.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		// log.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		// log.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		// log.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		// log.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
		return result, nil
	}
}

//PostUpdateFieldInMongoDBviaEmailID ...
func (r *MongoClient) PostUpdateFieldInMongoDBviaEmailID(fnv *forms.EmailStruct) (*mongo.UpdateResult, error) {
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.D{{"email", fnv.EmailID}}
	update := bson.D{
		{"$set", bson.D{
			{fnv.FieldName, fnv.FieldValue},
		}},
	}
	collx := r.db.Database("test").Collection("people")
	result, err := collx.UpdateOne(context.TODO(), filter, update)
	// Check for error, else print the UpdateOne() API call results
	if err != nil {
		log.Error(err.Error())
		return nil, err
	} else {
		// log.Println("UpdateOne() result:", result)
		// log.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		// log.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		// log.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		// log.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		// log.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
		return result, nil
	}
}

// //SavePeopleInMongoDB ...
// func (r *MongoClient) SavePeopleInMongoDB(people *forms.PersonDetails) (*mongo.InsertOneResult, error) {
// 	var peopleDetails forms.PersonDetails
// 	filter := bson.D{{"name", people.Name}}
// 	collx := r.db.Database("test").Collection("people")
// 	err := collx.FindOne(context.TODO(), filter).Decode(&peopleDetails)
// 	if err != nil {
// 		insertResult, err := collx.InsertOne(context.TODO(), people)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return insertResult, nil
// 	}
// 	err1 := errors.New("math: User already exists")
// 	return nil, err1
// }
