package models

import (
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// User field (Model) defined
type User struct {
	gorm.Model
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastNames string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

// Bucketlist field (Model) defined
type Bucketlist struct {
	gorm.Model
	Name string `json:"name"`
	User *User  `json:"user"`
}

// BucketlistItem field (Model) defined
type BucketlistItem struct {
	gorm.Model
	Name       string      `json:"name"`
	Done       bool        `json:"done"`
	Bucketlist *Bucketlist `json:"bucketlist,omitempty"`
}

// Token JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Validate incoming user details...
func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

// Create a user object
func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("PASSPHRASE")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := utils.Message(true, "user has been created")
	response["user"] = user
	return response
}
