package forms

import "time"

//TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     string
}

// //AccessDetails ...
// type AccessDetails struct {
// 	UUID   string
// 	UserID string
// }

//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AccessRefreshTokenTTL ...
type AccessRefreshTokenTTL struct {
	CreatedAt time.Time
	Time      time.Time
	UserID    string
	UUID      string
	TokenType string
	Token     string
}
