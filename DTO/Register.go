package dto
type RequestRegister struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
type ResponseRegister struct{
	ID string `json:"id"`
}