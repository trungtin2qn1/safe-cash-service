package user

//AccessTokenRequest ...
type AccessTokenRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

//AccessTokenResponse ...
type AccessTokenResponse struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}