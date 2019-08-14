package models

import (
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// User field (Model) defined
type User struct {
	gorm.Model
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	Token      string       `json:"token"`
	Bucketlist []Bucketlist `json:"bucketlist,omitempty"`
}

// Token JWT claims struct
type Token struct {
	UserID uint
	Email  string
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

	fmt.Println(user, "the user object")
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
	expirationTime := time.Now().Add(24 * time.Hour)
	tk := &Token{
		user.ID,
		user.Email,
		jwt.StandardClaims{
			Audience:  "devs",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "gobucketlistapi",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("PASSPHRASE")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := utils.Message(true, "user has been created")
	response["user"] = user
	return response
}

// Login function defined
func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserID: user.ID, Email: email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("PASSPHRASE")))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

// GetUser function defined to retrieve a user by id
func GetUser(u uint) *User {

	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}
