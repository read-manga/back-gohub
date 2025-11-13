package repo

type Stars struct {
	user_id int
}

type Repo struct {
	ID    int
	name  string
	about string
	tag   string
	stars []Stars

	public bool

	storage_s3 string

	user_id      int
	colaborators []int
}
