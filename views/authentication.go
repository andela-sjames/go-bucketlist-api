package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// CreateUserHandler function defined to create new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and fail if any error occur
	if err != nil {
		fmt.Println("Debug user CreateUserHandler:", err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create user
	utils.Respond(w, resp)
}

// AuthenticateHandler function defined to authenticate new users
func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and fail if any error occur
	if err != nil {
		fmt.Println("Debug user AuthenticateHandler:", err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	utils.Respond(w, resp)
}
