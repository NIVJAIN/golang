package controllers

import (
	"fmt"

	// "log"
	// "log"
	"net/http"
	"time"

	"github.com/forestgiant/sliceutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/forms"
	"github.com/sirupsen/logrus"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/models"
)

var validate *validator.Validate

// var logger *logrus.Logger

// UserController will hold everything that controller needs
type UserController struct {
	userRepo models.UserRepository
	authRepo models.AuthRepository
	logras   map[string]*logrus.Logger
}

// SetUserController returns a new UserHandler, this stuff is called dependency Injection took a while to understand...
func (h *UserController) SetUserController(userRepo models.UserRepository, authRepo models.AuthRepository, logams map[string]*logrus.Logger) *UserController {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	log = logams["info"]
	return &UserController{
		userRepo: userRepo,
		authRepo: authRepo,
		logras:   logams,
	}
}

//RegisterUser ...
func (h *UserController) RegisterUser(c *gin.Context) {
	var userDetails forms.UserDetails
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepo.RegisterUser(&userDetails)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": user})
}

// Login allows a user to login a user and get access token
func (h *UserController) Login(c *gin.Context) {
	log.Info("Iam from Controller calling...")
	var usersLoginData forms.LoginUserCommand
	// Bind the request body data to var data and check if all details are provided need to figure out how to do validationg using validator library
	err := c.ShouldBindJSON(&usersLoginData)
	// not sure on best practice to put validation in controller or in services
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	user, err := h.userRepo.Login(&usersLoginData)
	if err != nil {
		log.Error("Error", err.Error(), user)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"success": user})
	c.JSON(200, gin.H{"message": "Log in success", "accessToken": user.AccessToken, "refrestToken": user.RefreshToken})
}

// Logout ...
func (h *UserController) Logout(c *gin.Context) {
	err := h.userRepo.Logout(c)
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

//FindUser ...
func (h *UserController) FindUser(c *gin.Context) {
	var user forms.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getUser, err := h.userRepo.FindByEmailID(user.Email)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	// log.Println("GetUser..", getUser, user)
	c.JSON(http.StatusOK, gin.H{"success": getUser})
}

//CreateUserPerson ...
func (h *UserController) CreateUserPerson(c *gin.Context) {
	var peopleDetails forms.PersonDetails
	if err := c.ShouldBindJSON(&peopleDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultStatus, err := h.userRepo.NotExistsSavePeopleInMongoDB(&peopleDetails)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})
}

//QueryStrings ... GET {{baseUrl}}/user/query?filtername=name&filtervalue=Jain5&fieldname=city&fieldvalue=Germany
func (h *UserController) QueryStrings(c *gin.Context) {
	filtername := c.Query("filtername") //filtername=jain
	filtervalue := c.Query("filtervalue")
	fieldname := c.Query("fieldname")
	filedvalue := c.Query("fieldvalue")
	var fvn forms.FieldNameValue
	// log.Println(".......*******...........", filtername, fieldname, filedvalue)
	fvn.FilterName = filtername
	fvn.FilterValue = filtervalue
	fvn.FieldName = fieldname
	fvn.FieldValue = filedvalue

	// log.Println("ddddd", fvn)
	resultStatus, err := h.userRepo.UpdateFieldValueInMongoDB(&fvn)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})
	// c.JSON(200, gin.H{
	// 	"name": name,
	// 	"age":  age,
	// })
}

// user.GET("/querydates", usercontroller.QueryMongoDBDates)
//  {{baseUrl}}/user/querydates?filtername=unixnano&startdate=2016-12-01&todate=2017-12-24&limit=5

// QueryMongoDBDates ...
func (h *UserController) QueryMongoDBDates(c *gin.Context) {
	var fvn forms.QueryDates
	// filtername := c.Query("filtername") //filtername=unixnano or unix or createdat
	// startdate := c.Query("startdate")
	// todate := c.Query("todate")
	// limit := c.Query("limit")
	// log.Println(".......*******...........", filtername, startdate, todate, limit)
	// fvn.FilterName = filtername
	// fvn.StartDate = filtervalue
	// fvn.ToDate = fieldname
	// fvn.Limit = limit

	if err := c.ShouldBindQuery(&fvn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := time.Parse("2006-01-02", fvn.StartDate)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	_, err = time.Parse("2006-01-02", fvn.ToDate)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	resultStatus, err := h.userRepo.QueryDatesMongoDB(&fvn)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})

	// resultStatus, err := h.userRepo.QueryDatesMongoDB(&fvn)
	// if err != nil {
	// 	c.JSON(403, gin.H{"message": err.Error()})
	// 	c.Abort()
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"success": resultStatus})
}

//QueryStringsUpdateFieldViaMongoID ... GET {{baseUrl}}/user/query?filtername=name&fieldname=city&fieldvalue=germany
func (h *UserController) QueryStringsUpdateFieldViaMongoID(c *gin.Context) {
	mongoid := c.Query("id") //filtername=jain
	fieldname := c.Query("fieldname")
	filedvalue := c.Query("fieldvalue")
	var fvn forms.MongoidStruct
	fvn.MongoID = mongoid
	fvn.FieldName = fieldname
	fvn.FieldValue = filedvalue
	log.Println(".......*******...........", mongoid, fieldname, filedvalue)
	resultStatus, err := h.userRepo.UpdateFieldValueInMongoDBUsingMongoID(&fvn)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})
	// c.JSON(200, gin.H{
	// 	"name": name,
	// 	"age":  age,
	// })
}

//QueryStringsUpdateFieldViaEmailID ... GET {{baseUrl}}/user/query?filtername=name&fieldname=city&fieldvalue=germany
func (h *UserController) QueryStringsUpdateFieldViaEmailID(c *gin.Context) {
	email := c.Query("email") //filtername=jain
	fieldname := c.Query("fieldname")
	filedvalue := c.Query("fieldvalue")
	var fvn forms.EmailStruct
	fvn.EmailID = email
	fvn.FieldName = fieldname
	fvn.FieldValue = filedvalue
	log.Println(".......*******...........", email, fieldname, filedvalue)
	resultStatus, err := h.userRepo.UpdateFieldValueInMongoDBUsingEmailID(&fvn)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})
	// c.JSON(200, gin.H{
	// 	"name": name,
	// 	"age":  age,
	// })
}

//PostUpdateUserInMongoDB ...
func (h *UserController) PostUpdateUserInMongoDB(c *gin.Context) {
	var emailDetails forms.EmailStruct
	if err := c.BindJSON(&emailDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// listofValidItems := []string{"name", "age", "city", "email"}
	listofitems := []string{"email", "age"}
	if sliceutil.Contains(listofitems, emailDetails.FieldName) {
		response := fmt.Sprintf("Changing of [%s] is not allowed", emailDetails.FieldName)
		c.JSON(403, gin.H{"message": response})
		c.Abort()
		return
	}
	if emailDetails.FieldName == "email" || emailDetails.FieldName == "age" {
		c.JSON(403, gin.H{"message": "Changing of EmailID or age is not allowed"})
		c.Abort()
		return
	}
	resultStatus, err := h.userRepo.PostUpdateFieldInMongoDBviaEmailID(&emailDetails)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": resultStatus})
}

//PathQueryStrings2 ...
func (h *UserController) PathQueryStrings2(c *gin.Context) {
	name := c.Param("name") //name=jain
	age := c.Param("age")   //age=41
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// //RegisterUser ...
// func (h *UserController) RegisterUser(c *gin.Context) {
// 	var userDetails forms.UserDetails
// 	userDetails.CreatedAt = time.Now()
// 	userDetails.NanoUnix = time.Now().UnixNano()
// 	if err := c.ShouldBindJSON(&userDetails); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	findUserEmailExist, _ := h.userRepo.FindByEmailID(userDetails.Email)
// 	log.Println("SDdddd", findUserEmailExist)
// 	if findUserEmailExist != nil {
// 		c.JSON(403, gin.H{"message": "User Email already exists"})
// 		c.Abort()
// 		return
// 	}

// 	user, err := h.userRepo.RegisterUser(&userDetails)
// 	if err != nil {
// 		fmt.Println("Error", user)
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": user})
// }

// // Login allows a user to login a user and get access token
// func (h *UserController) Login(c *gin.Context) {
// 	var data forms.LoginUserCommand
// 	// Bind the request body data to var data and check if all details are provided
// 	err := c.ShouldBindJSON(&data)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		c.Abort()
// 		return
// 	}
// 	err = validate.Struct(data)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"validator": err.Error()})
// 		c.Abort()
// 		return
// 	}
// 	// if c.ShouldBindJSON(&data) != nil {
// 	// 	c.JSON(406, gin.H{"message": "Provide required details email and password", "error": })
// 	// 	c.Abort()
// 	// 	return
// 	// }
// 	result, err := h.userRepo.GetUserByEmail(data.Email)
// 	fmt.Println("result---", result, err)
// 	if err != nil {
// 		c.JSON(400, gin.H{"message": "Problem logging into your account"})
// 		c.Abort()
// 		return
// 	}
// 	if result.Email == "" {
// 		c.JSON(404, gin.H{"message": "User account was not found"})
// 		c.Abort()
// 		return
// 	}
// 	// Get the hashed password from the saved document
// 	// Get the hashed password from the saved document
// 	hashedPassword := []byte(result.Password)
// 	// Get the password provided in the request.body
// 	password := []byte(data.Password)

// 	err = helpers.PasswordCompare(password, hashedPassword)

// 	if err != nil {
// 		c.JSON(403, gin.H{"message": "Invalid user credentials"})
// 		c.Abort()
// 		return
// 	}
// 	//Generate the JWT auth token
// 	tokenDetails, err := h.authRepo.CreateToken(data.Email)
// 	fmt.Println("tokendetails::-->>", tokenDetails)

// 	err = h.authRepo.CreateAuth(data.Email, tokenDetails)
// 	if err != nil {
// 		c.JSON(403, gin.H{"message": "Invalid user credentials", "err": err.Error()})
// 		c.Abort()
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Log in success", "accessToken": tokenDetails.AccessToken, "refrestToken": tokenDetails.RefreshToken})
// 	c.Abort()
// 	return
// }

// // Logout ...
// func (h *UserController) Logout(c *gin.Context) {
// 	au, err := h.authRepo.ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
// 		return
// 	}

// 	deleted, delErr := h.authRepo.DeleteAuth(au.AccessUUID)
// 	if delErr != nil || deleted == nil { //if any goes wrong
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
// }
// //CreateUser ...
// func (h *UserHandler) CreateUser(c *gin.Context) {
// 	var userDetails models.UserDetails
// 	if err := c.ShouldBindJSON(&userDetails); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := h.userRepo.SaveInMongoDB(&userDetails)
// 	if err != nil {
// 		fmt.Println("Error", user)
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": user})
// }
