package commit

import "time"

type Commit struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RepoId      string    `json:"repo_id"`
	UserId      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
