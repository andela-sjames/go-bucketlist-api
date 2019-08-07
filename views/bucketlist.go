package views

import (
	"encoding/json"
	"net/http"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// CreateBucketlistHandler function defined to create new user
func CreateBucketlistHandler(w http.ResponseWriter, r *http.Request) {

	bucketlist := &models.Bucketlist{}
	err := json.NewDecoder(r.Body).Decode(bucketlist) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := bucketlist.Create() //Create user
	utils.Respond(w, resp)
}
