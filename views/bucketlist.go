package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// CreateBucketlistHandler function defined to create new user
func CreateBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	userObj := r.Context().Value("userObj").(map[string]interface{}) //Grab the userObj of the user that send the request

	bucketlist := &models.Bucketlist{}
	err := json.NewDecoder(r.Body).Decode(bucketlist) //decode the request body into struct and fail if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	fmt.Println("userObj:", userObj)

	userID := userObj["userID"].(uint)
	userEmail := userObj["userEmail"].(string)

	bucketlist.UserID = userID
	bucketlist.CreatedBy = userEmail
	resp := bucketlist.Create() //Create user
	utils.Respond(w, resp)
}
