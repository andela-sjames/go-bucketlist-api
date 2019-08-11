package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
	"github.com/gorilla/mux"
)

// CreateItemHandler function defined to create a new bucketlist
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	id, err := strconv.Atoi(requestParams["id"])

	if err != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	bucketlistItem := &models.BucketlistItem{}
	decodeErr := json.NewDecoder(r.Body).Decode(bucketlistItem)

	if decodeErr != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request body"))
		return
	}

	bucketlist := models.GetBucketlist(uint(id))

	if bucketlist == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("bucket list with id: %d was not found", id)))
	}

	bucketlistItem.BucketlistID = bucketlist.ID
	resp := bucketlistItem.Create()

	utils.Respond(w, resp)
}
