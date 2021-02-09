package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/controllers"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/repositories"
	"github.com/sirupsen/logrus"

	// "github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares"
	// "github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/db"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/middlewares"

	"go.mongodb.org/mongo-driver/mongo"
)

// "github.com/nivjain/7-gininterfaceMongoDB/controllers"

//UserRouter ...
type UserRouter struct{}

var uc = new(controllers.UserController)
var ac = new(controllers.AuthController)

//HomePage ...
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi Sripal Jain greetings from userroutes homepage",
	})
}

// type Person struct {
// 	propName  string `json:"propName"`
// 	propValue string `json:"propValue"`
// }

// type Personses struct {
// 	Persons []Person
// }

type CreateParams struct {
	Username     string `json:"username"`
	Guests       Guests `json:"guests"`
	RoomType     string `json:"roomType"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
}

type Guests struct {
	Person []Person `json:"person"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Guests2 struct {
	Person []Person `json:"person"`
}
type Guests3 struct {
	Person []Person
}

type Person2 struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// func Submit4(c *gin.Context) {
// 	var data Personses
// 	// data := new(List)
// 	err := c.BindJSON(&data)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"err0r": err.Error()})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, data)
// }
func Submit5(c *gin.Context) {
	var f CreateParams
	if err := c.BindJSON(&f); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, f)

}

func Submit(c *gin.Context) {
	var f Guests3
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for a, b := range f.Person {
		fmt.Println(a, b.Firstname, b.Lastname)
	}
	c.IndentedJSON(http.StatusOK, f)

}

func Submitworking(c *gin.Context) {
	var f Guests2
	if err := c.BindJSON(&f); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, f)

}

// func Submit(c *gin.Context) {

// 	var f interface{}
// 	var arr []Person
// 	if err := c.BindJSON(&f); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println(f)
// 	myMap := f.([]interface{})
// 	// elementMap := make(map[string]string)
// 	for _, v := range myMap {
// 		// fmt.Println("k", k)
// 		// fmt.Println("v", reflect.TypeOf(v), v)
// 		for gkey, mvalue := range v.(map[string]interface{}) {
// 			fmt.Println("gkey:", gkey, "mvalue:", mvalue)
// 			// b := Person{propName: mvalue}
// 			// data = append(data, b)
// 			var b Person
// 			b.propName = gkey
// 			b.propValue = fmt.Sprintf("%v", mvalue)
// 			arr = append(arr, b)
// 		}
// 		fmt.Println(arr)
// 	}

// 	c.IndentedJSON(http.StatusOK, arr)

// }

// func Submit(c *gin.Context) {
// 	var p []Person
// 	// var f interface{}

// 	if err := c.BindJSON(&p); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println(p)
// 	// for k, v := range f.([]interface{}) {
// 	// 	p = append(p, v)
// 	// 	fmt.Println("ddddd", k, v)
// 	// }
// 	c.IndentedJSON(http.StatusOK, p)
// 	// c.JSON(200, gin.H{
// 	// 	"message": "Hi Sripal Jain greetings from userroutes homepage",
// 	// })
// }

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware(ac *controllers.AuthController) gin.HandlerFunc {
	return func(c *gin.Context) {
		ac.TokenValid(c)
		c.Next()
	}
}

// Routes for user
func (usr UserRouter) Routes(dba *mongo.Client, collxpool map[string]*mongo.Collection, logpool map[string]*logrus.Logger, mClient *repositories.MongoClient, router *gin.Engine) {

	// authRepo := repositories.SetMongoClient(db, collxpool, nil)
	// userRepo := repositories.SetMongoClient(dba, collxpool, nil, logpool)
	// usercontroller := uc.SetUserController(userRepo, userRepo, logpool)
	// authCntrler := ac.SetAuthController(userRepo, logpool)
	usercontroller := uc.SetUserController(mClient, mClient, logpool)
	authCntrler := ac.SetAuthController(mClient, logpool)

	TokenMW := TokenAuthMiddleware(authCntrler)
	user := router.Group("/user")

	// user.GET("/gethome", usercontroller.GetHomePage)
	// user.GET("/gethome", TokenMW, HomePage)
	// rateLimitter := middlewares.NewRateLimiterMiddleware(redisClientabc, "key", 1, 60*time.Second)
	rateLimitter := middlewares.NewRateLimiterMiddleware("key", 1, 60*time.Second)

	user.GET("/gethome", TokenMW, rateLimitter, HomePage)
	user.POST("/signup", usercontroller.RegisterUser)
	user.POST("/submit", Submit)
	user.POST("/login", usercontroller.Login)
	user.GET("/logout", usercontroller.Logout)
	user.POST("/refresh", authCntrler.Refresh)
	// user.POST("/getuser", TokenMW, usercontroller.FindUser)
	user.POST("/getuser", usercontroller.FindUser)
	user.POST("/createpeople", usercontroller.CreateUserPerson)
	user.GET("/query", usercontroller.QueryStrings)                            //GET {{baseUrl}}/user/query?filtername=name&filtervalue=Jain5&fieldname=city&fieldvalue=Germany
	user.GET("/updatefield", usercontroller.QueryStringsUpdateFieldViaMongoID) //GET {{baseUrl}}/user/updatefield?id=5f8d5ea901b19460e735f888&fieldname=name&fieldvalue=Sripal
	user.GET("/emailfield", usercontroller.QueryStringsUpdateFieldViaEmailID)  //GET {{baseUrl}}/user/emailfield?email=sripal.jain@gmail.com&fieldname=name&fieldvalue=Sripal
	user.POST("/update", usercontroller.PostUpdateUserInMongoDB)

	user.GET("/querydates", usercontroller.QueryMongoDBDates)
}

//GET {{baseUrl}}/user/query?filtername=name&filtervalue=Jain5&fieldname=city&fieldvalue=Germany
// {
// 	"_id" : ObjectId("5f8d5ea901b19460e735f888"),
// 	"name" : "Jain5",
// 	"age" : 41,
// 	"city" : "Russia"
// }

// {
// 	"_id" : ObjectId("5f8d5ea901b19460e735f888"),
// 	"name" : "Jain5",
// 	"age" : 41,
// 	"city" : "Germany"
// }

// var usercontroller = new(controllers.UserController)

// // Routes for user
// func (usr UserRouter) Routes(router *gin.Engine) {
// 	user := router.Group("/user")
// 	user.GET("/gethome", usercontroller.GetHomePage)
// 	user.POST("/posthome", usercontroller.PostHomePage)
// 	user.GET("/query", usercontroller.QueryStrings)
// 	user.GET("/path/:name/:age", usercontroller.PathQueryStrings)
// 	user.POST("/signup", usercontroller.CreateUser)
// 	// user.POST("/login", usercontroller.Login)
// }
