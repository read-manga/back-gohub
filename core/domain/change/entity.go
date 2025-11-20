package change

type ChangeType string

const (
	Added    ChangeType = "ADDED"
	Modified ChangeType = "MODIFIED"
	Removed  ChangeType = "REMOVED"
)

type Change struct {
	ID           string     `json:"id"`
	CommitId     string     `json:"commit_id"`
	FilePath     string     `json:"file_path"`
	ChangeType   ChangeType `json:"change_type"`
	PreviousHash string     `json:"previous_hash"`
	NewHash      string     `json:"new_hash"`
}

type NewFilesBody struct {
	Path    string
	Hash    string
	Content []byte
}
