package controller

import (
	"encoding/json"
	"net/http"

	dto "github.com/gabrielribeirofsouza/OpsCenter-API/DTO"
	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
	"github.com/gabrielribeirofsouza/OpsCenter-API/service"
)

type AuthController struct {
	service service.UserService
}
func NewAuthController(us service.UserService) *AuthController{
return &AuthController{
	service: us,
}}
func(ac *AuthController)Login(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "[ERROR] - Metodo inválido", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, "[ERROR] - Body inválido", http.StatusBadRequest)
		return
	}
	var body dto.AuthRequestLogin
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "[ERROR]", http.StatusInternalServerError)
	}
	response, err := ac.service.AuthUser(body)
	if err != nil {
		http.Error(w, "[ERROR] - Erro ao efetuar login", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func(ac *AuthController)Register(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "[ERROR] - Metodo inválido", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, "[ERROR] - Body inválido", http.StatusBadRequest)
		return
	}
	
	var body entity.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "[ERROR]", http.StatusInternalServerError)
	}
	response, err := ac.service.Register(body)
	if err != nil{
		http.Error(w, "[ERROR] - Erro ao criar usuário", http.StatusNotAcceptable)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}