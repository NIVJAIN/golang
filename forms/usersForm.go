package forms

import (
	"time"
)

//UserDetails ..
type UserDetails struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	NanoUnix  int64
}

// LoginUserCommand defines user login form struct
type LoginUserCommand struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6" validate="is-cool"`
}

// User ...
type User struct {
	Email string `json:"email" binding:"required"`
}

//PersonDetails ..
type PersonDetails struct {
	Name string `json:"name" binding:"required"`
	// Age   int    `json:"Age" binding:"required, gte=1 lte=130"`
	Age   int    `json:"Age" binding:"required"`
	City  string `json:"city" binding:"required"`
	Email string `json:"email" binding:"required"`
}

//FieldNameValue ...
type FieldNameValue struct {
	FilterName  string `json:"filtername" binding:"required"`
	FilterValue string `json:"filtevalue" binding:"required"`
	FieldName   string `json:"fieldname" binding:"required"`
	FieldValue  string `json:"fieldvalue" binding:"required"`
}

//MongoidStruct ...
type MongoidStruct struct {
	MongoID    string `json:"id" binding:"required"`
	FieldName  string `json:"fieldname" binding:"required"`
	FieldValue string `json:"fieldvalue" binding:"required"`
}

//EmailStruct ...
type EmailStruct struct {
	EmailID    string `json:"email" binding:"required"`
	FieldName  string `json:"fieldname" binding:"required"`
	FieldValue string `json:"fieldvalue" binding:"required"`
}

// QueryDates ...
type QueryDates struct {
	FilterName string `form:"filtername" json:"filtername" binding:"required"`
	StartDate  string `form:"startdate" json:"startdate" binding:"required"`
	ToDate     string `form:"todate" json:"todate" binding:"required"`
	Limit      int    `form:"limit" json:"limit" binding:"required"`
}
