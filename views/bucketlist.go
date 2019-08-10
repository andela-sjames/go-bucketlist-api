package views

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/andela-sjames/go-bucketlist-api/auth"
	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// CreateBucketlistHandler function defined to create a new bucketlist
func CreateBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	userObj := r.Context().Value(auth.CtxKey).(map[string]interface{}) //Grab the userObj of the user that send the request

	bucketlist := &models.Bucketlist{}
	err := json.NewDecoder(r.Body).Decode(bucketlist) //decode the request body into struct and fail if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	userID := userObj["userID"].(uint)
	userEmail := userObj["userEmail"].(string)

	bucketlist.UserID = userID
	bucketlist.CreatedBy = userEmail
	resp := bucketlist.Create() //Create user
	utils.Respond(w, resp)
}

// GetAllBucketlistHandler function defined to list all bucketlist for a
// specific user
func GetAllBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	userObj := r.Context().Value(auth.CtxKey).(map[string]interface{})
	data := models.GetAllBucketlist(userObj["userID"].(uint))

	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

// GetBucketByIDlistHandler function defined to get a single bucketlist
func GetBucketByIDlistHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetBucketlist(uint(id))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}