package repositories

import (
	"context"
	"fmt"

	// "log"
	"net/http"
	"os"
	"strings"
	"time"

	// "github.com/Massad/gin-boilerplate/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	uuid "github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "gopkg.in/mgo.v2/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateToken ...
func (r *MongoClient) CreateToken(userID string) (*forms.TokenDetails, error) {

	td := &forms.TokenDetails{}
	// td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AtExpires = time.Now().Add(time.Minute * 1).Unix()
	td.AccessUUID = uuid.NewV4().String()

	// td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RtExpires = time.Now().Add(time.Minute * 2).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//CreateAuth ...
func (r *MongoClient) CreateAuth(userid string, td *forms.TokenDetails) error {
	// redisClient :=
	// at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	// rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	accessTokenToDB := forms.AccessRefreshTokenTTL{
		CreatedAt: now,
		Time:      now,
		UserID:    userid,
		UUID:      td.AccessUUID,
		TokenType: "AccessToken",
		Token:     td.AccessToken,
	}
	refreshTokenToDB := forms.AccessRefreshTokenTTL{
		CreatedAt: now,
		Time:      now,
		UserID:    userid,
		UUID:      td.RefreshUUID,
		TokenType: "RefreshToken",
		Token:     td.RefreshToken,
	}
	// fmt.Println("creatauth22::-->>", rt, at, userid, "td.AccessToken", "td.AccessUUID", "td.RefreshUUID")
	// fmt.Println("at.sub.now::-->>", at.Sub(now))
	// fmt.Println("at.sub.now::-->>", rt.Sub(now))
	// log.Println("ABBCBCBCBCB==>>", &user)
	// log.Println("----ABBCBCBCBCB==>>", user.CreatedAt)

	// user.createdAt = abnow
	// user.userID = userid
	// user.AccessUUID = td.AccessUUID

	_, err := r.collection["tokens"].InsertOne(context.TODO(), &accessTokenToDB)
	if err != nil {
		log.Println("Unabelt to save in db for accesstoken ", err)
		return err
	}

	_, err = r.collection["tokens"].InsertOne(context.TODO(), &refreshTokenToDB)
	if err != nil {
		log.Println("Unabelt to save for refreshtoken ", err)
		return err
	}

	// output: {Id:ObjectIdHex("572f3c68e43001d2c1703aa7") Time:2015-07-08 17:29
	// errAccess := db.GetRedis().Set(td.AccessUUID, userid, at.Sub(now)).Err()
	// // errAccess := db.GetRedis().Set("asddsfs", strconv.Itoa(int(25)), at.Sub(now)).Err()
	// if errAccess != nil {
	// 	return errAccess
	// }

	// errRefresh := db.GetRedis().Set(td.RefreshUUID, userid, rt.Sub(now)).Err()
	// // errRefresh := db.GetRedis().Set("sskks", strconv.Itoa(int(30)), rt.Sub(now)).Err()

	// if errRefresh != nil {
	// 	return errRefresh
	// }
	return nil
}

//TokenValidFromRepo ...
func (r *MongoClient) TokenValidFromRepo(req *http.Request) error {
	token, err := r.VerifyToken(req)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//VerifyToken ...
func (r *MongoClient) VerifyToken(req *http.Request) (*jwt.Token, error) {
	tokenString := r.ExtractToken(req)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken ...
func (r *MongoClient) ExtractToken(req *http.Request) string {
	bearToken := req.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//ExtractTokenMetadata ...
func (r *MongoClient) ExtractTokenMetadata(req *http.Request) (*forms.AccessDetails, error) {
	token, err := r.VerifyToken(req)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		return &forms.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

//FetchAuth ...
func (r *MongoClient) FetchAuth(authD *forms.AccessDetails) (string, error) {
	var result forms.AccessDetails
	filter := bson.D{{"uuid", authD.AccessUUID}}
	err := r.collection["people"].FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	// fmt.Printf("Found a single document: %+v\n", result)
	return string(result.UserID), nil
}

//DeleteAuth ...
func (r *MongoClient) DeleteAuth(givenUUID string) (*mongo.DeleteResult, error) {
	deleted, err := r.collection["tokens"].DeleteOne(context.TODO(), bson.D{{"uuid", givenUUID}})
	if err != nil {
		return nil, err
	}
	log.Info("Deleted....", deleted, givenUUID)
	return deleted, nil
}

// //FetchAuth ...
// func (r *MongoClient) FetchAuth(authD *AccessDetails) (string, error) {
// 	userid, err := db.GetRedis().Get(authD.AccessUUID).Result()
// 	if err != nil {
// 		return userid, err
// 	}
// 	// userID, _ := strconv.ParseInt(userid, 10, 64)
// 	// userID, ok := authD.UserID
// 	// if !ok {
// 	// 	return nil, err
// 	// }

// 	return userid, nil
// }

// //DeleteAuth ...
// func (r *MongoClient) DeleteAuth(givenUUID string) (int64, error) {
// 	deleted, err := db.GetRedis().Del(givenUUID).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }
