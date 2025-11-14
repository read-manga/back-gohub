package user

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	Bio         string `json:"bio"`
	Profile_url string `json:"profile_url"`
	Status      string `json:"status"`
}
