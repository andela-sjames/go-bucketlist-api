package models

import (
	"fmt"

	"github.com/andela-sjames/go-bucketlist-api/utils"
	"github.com/jinzhu/gorm"
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

// GetAllBucketlist function defined
func GetAllBucketlist(user uint) []*Bucketlist {

	bucketlists := make([]*Bucketlist, 0)
	err := GetDB().Table("bucketlists").Where("user_id = ?", user).Find(&bucketlists).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bucketlists
}

// GetBucketlist by ID function defined
func GetBucketlist(id uint) *Bucketlist {

	bucketlist := &Bucketlist{}
	err := GetDB().Table("bucketlists").Where("id = ?", id).First(bucketlist).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil
	}

	GetDB().Model(&bucketlist).Association("Item").Find(&bucketlist.Item)
	return bucketlist
}

// UpdateBucketlist by ID functio  defined
func UpdateBucketlist(id uint, name string) *Bucketlist {
	bucketlist := &Bucketlist{}

	err := GetDB().Table("bucketlists").Where("id = ?", id).First(bucketlist).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil
	}

	GetDB().Model(&bucketlist).Update("name", name)
	return bucketlist
}
