package models

import (
	"github.com/andela-sjames/go-bucketlist-api/utils"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

// Bucketlist field (Model) defined
type Bucketlist struct {
	gorm.Model
	Name      string           `json:"name"`
	CreatedBy string           `json:"created_by"`
	UserID    uint             `json:"user_id"`
	Item      []BucketlistItem `json:"item"`
}

// Create a bucketlist function defined
func (bucketlist *Bucketlist) Create() map[string]interface{} {

	if bucketlist.Name == "" {
		return utils.Message(false, "No input data provided")
	}

	GetDB().Create(bucketlist)
	response := utils.Message(true, "bucketlist created")
	response["bucketlist"] = bucketlist
	return response
}

// BeforeCreate hook defined for bucketlist
func (bucketlist *Bucketlist) BeforeCreate(scope *gorm.Scope) error {
	// update CreatedBy here with user name from the request context
	// https://golang.org/pkg/context/#example_WithValue
	scope.SetColumn("ID", uuid.New())
	return nil
}
