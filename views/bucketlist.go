package views

import (
	"encoding/json"
	"net/http"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// CreateBucketlistHandler function defined to create new user
func CreateBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	UserID := r.Context().Value("UserID").(uint)         //Grab the id of the user that send the request
	userEmail := r.Context().Value("userEmail").(string) //Grab the email of the user that send the request

	bucketlist := &models.Bucketlist{}
	err := json.NewDecoder(r.Body).Decode(bucketlist) //decode the request body into struct and fail if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	bucketlist.UserID = UserID
	bucketlist.CreatedBy = userEmail
	resp := bucketlist.Create() //Create user
	utils.Respond(w, resp)
}
