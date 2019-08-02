package models

// Bucketlist field collection defined
type Bucketlist struct {
	Name     string
	Created  string
	Modified string
	user     string
}

// BucketlistItem field collectio defined
type BucketlistItem struct {
	name       string
	Created    string
	Modified   string
	Done       string
	Bucketlist string
}
