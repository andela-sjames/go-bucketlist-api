package models

import (
	"github.com/andela-sjames/go-bucketlist-api/utils"
	"github.com/jinzhu/gorm"
)

// BucketlistItem field (Model) defined
type BucketlistItem struct {
	gorm.Model
	Name         string `json:"name"`
	Done         bool   `json:"done"`
	BucketlistID uint   `json:"bucketlist_id,omitempty"`
}

// Create method for Bucketlist defined
func (bucketlistItem *BucketlistItem) Create() map[string]interface{} {

	if bucketlistItem.Name == "" {
		return utils.Message(false, "No input data provided")
	}

	GetDB().Create(bucketlistItem)
	response := utils.Message(true, "bucketlistitem created")
	response["bucketlistItem"] = bucketlistItem
	return response
}

// GetBucketItem by ID function defined
func GetBucketItem(id uint) *BucketlistItem {

	bucketlistItem := &BucketlistItem{}
	err := GetDB().Table("bucketlist_items").Where("id = ?", id).First(bucketlistItem).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil
	}

	return bucketlistItem
}

// UpdateBucketItem by ID functio  defined
func UpdateBucketItem(id uint, name string, done bool) *BucketlistItem {
	bucketlistItem := &BucketlistItem{}

	err := GetDB().Table("bucketlist_items").Where("id = ?", id).First(bucketlistItem).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil
	}

	GetDB().Model(&bucketlistItem).Update(BucketlistItem{Name: name, Done: done})
	return bucketlistItem
}

// DeleteBucketItem function defined
func DeleteBucketItem() {

}
