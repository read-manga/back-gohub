package repo

type Stars struct {
	user_id int
}

type Repo struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	About        string   `json:"about"`
	Tag          []string `json:"tag"`
	Stars        []Stars  `json:"stars"`
	Public       bool     `json:"public"`
	StorageS3    string   `json:"storage_s3"`
	UserId       string   `json:"user_id"`
	Colaborators []string `json:"colaborators"`
}
