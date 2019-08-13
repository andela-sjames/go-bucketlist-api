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

// CreateItemHandler function defined to create a new bucketlist item
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

// UpdateDeleteItemHandler function defined to update an item
func UpdateDeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	id, errID := strconv.Atoi(requestParams["id"])
	itemID, erritemID := strconv.Atoi(requestParams["itemID"])

	if errID != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request, bucketlist id missing"))
		return
	}

	if erritemID != nil {
		// The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request, item id missing"))
		return
	}

	bucketlist := models.GetBucketlist(uint(id))

	if bucketlist == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("bucket list with id: %d was not found", id)))
	}

	// get item by id
	item := models.GetBucketItem(uint(itemID))

	if item == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("item with id: %d was not found", itemID)))
		return
	}

	if item.BucketlistID == bucketlist.ID {
		switch r.Method {
		case "PUT":
			bucketlistItem := &models.BucketlistItem{}
			decodeErr := json.NewDecoder(r.Body).Decode(bucketlistItem)

			if decodeErr != nil {
				// The passed path parameter is not an integer
				utils.Respond(w, utils.Message(false, "There was an error in your request body"))
				return
			}

			data := models.UpdateBucketItem(uint(itemID), bucketlistItem.Name, bucketlistItem.Done)
			resp := utils.Message(true, "success")
			resp["data"] = data
			utils.Respond(w, resp)
		case "DELETE":
			models.GetDB().Unscoped().Delete(&item)
			resp := utils.Message(true, "success")
			utils.Respond(w, resp)
		}
	} else {
		utils.Respond(w, utils.Message(false, "bucketlist-items mistatch error"))
		return
	}
}
