package resources

type Authentication struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}
