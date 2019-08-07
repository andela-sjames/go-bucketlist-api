package models

import (
	"github.com/jinzhu/gorm"
)

// BucketlistItem field (Model) defined
type BucketlistItem struct {
	gorm.Model
	Name         string `json:"name"`
	Done         bool   `json:"done"`
	BucketlistID uint   `json:"bucketlist_id,omitempty"`
}
