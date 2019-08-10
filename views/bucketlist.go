package views

import (
	"encoding/json"
	"net/http"

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

// GetBucketlistHandler function defined to get a single bucketlist
func GetBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	bucketlist := &models.Bucketlist{}
	models.GetDB().Find(&bucketlist)

	resp := utils.Message(true, "Fetch complete")
	resp["bucketlist"] = bucketlist
	utils.Respond(w, resp)
}
