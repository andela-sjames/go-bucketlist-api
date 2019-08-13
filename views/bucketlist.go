package views

import (
	"encoding/json"
	"fmt"
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
	resp := bucketlist.Create() //Create bucketlist
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

	requestParams := mux.Vars(r)
	id, err := strconv.Atoi(requestParams["id"])

	if err != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetBucketlist(uint(id))

	if data == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("bucketlist with id: %d was not found", id)))
		return
	}
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

// UpdateBucketByIDlistHandler function defined to get a single bucketlist
func UpdateBucketByIDlistHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	id, err := strconv.Atoi(requestParams["id"])

	if err != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	bucketlist := &models.Bucketlist{}
	json.NewDecoder(r.Body).Decode(bucketlist)

	data := models.UpdateBucketlist(uint(id), bucketlist.Name)
	if data == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("bucketlist with id: %d was not found", id)))
		return
	}
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
