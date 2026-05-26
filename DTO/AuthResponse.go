package dto
type AuthResponse struct {
	AccessToken string `json:"acess_token"`
	TokenType  string `json:"token_type"`
	ExpiresIn  int64 `json:"expires_in"`
	Email      string `json:"email"`
	ID         string `json:"id"`
}
type AuthRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
