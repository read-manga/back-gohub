package change

type Changes struct {
	ID           string `json:"id"`
	CommitId     string `json:"commit_id"`
	FilePath     string `json:"file_path"`
	ChangeType   string `json:"change_type"`
	PreviousHash string `json:"previous_hash"`
	NewHash      string `json:"new_hash"`
}
