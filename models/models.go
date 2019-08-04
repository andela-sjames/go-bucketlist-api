package models

// User field (Model) defined
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastNames string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Bucketlist field (Model) defined
type Bucketlist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Created  string `json:"date_created"`
	Modified string `json:"date_modified"`
	User     *User  `json:"user"`
}

// BucketlistItem field (Model) defined
type BucketlistItem struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Created    string      `json:"date_created"`
	Modified   string      `json:"date_modified"`
	Done       bool        `json:"done"`
	Bucketlist *Bucketlist `json:"bucketlist,omitempty"`
}
