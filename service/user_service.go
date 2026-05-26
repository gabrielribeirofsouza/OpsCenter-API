package service

import (
	"errors"

	dto "github.com/gabrielribeirofsouza/OpsCenter-API/DTO"
	"github.com/gabrielribeirofsouza/OpsCenter-API/config"
	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
	"github.com/gabrielribeirofsouza/OpsCenter-API/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	AuthUser(r dto.AuthRequestLogin) (dto.AuthResponse, error)
	Register(user entity.User)(dto.ResponseRegister, error)
}
type userService struct{
	repo repository.UserRepository
}
func NewUserService(repoUser repository.UserRepository)UserService{
	return &userService{
		repo: repoUser,
	}
}
func(us *userService)AuthUser(r dto.AuthRequestLogin)(dto.AuthResponse, error){
	user, err := us.repo.AuthUser(r.Email, r.Password)
	if err != nil{
		return dto.AuthResponse{}, err
	}
	token, err := config.GenerateToken(user.Email, user.ID)
	if err != nil{
		return dto.AuthResponse{}, err
	}
	return token, nil
}
func(us *userService)Register(user entity.User)(dto.ResponseRegister, error){
	err := ValidateInfoUser(user)
		if err != nil {
			return dto.ResponseRegister{}, err
		}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	resp, err := us.Register(user)
	if err != nil{
		return dto.ResponseRegister{}, err
	}
	return dto.ResponseRegister{
		ID: resp.ID,
	}, nil
}
func ValidateInfoUser(user entity.User)(error){

	if user.Name == "" {
		return  errors.New("Há informações invalidas")
	}
	
	if user.Permissoes == "" {
		return  errors.New("Há informações invalidas")
	}
	
	if user.Departamento_id == "" {
		return  errors.New("Há informações invalidas")
	}
	
	if user.Email == "" {
		return  errors.New("Há informações invalidas")
	}
	
	if user.Password == "" {
		return  errors.New("Há informações invalidas")
	}
	
	if user.Cargo == "" {
		return  errors.New("Há informações invalidas")
	}
	if user.Equipe_id == "" {
		return  errors.New("Há informações invalidas")
	}
	
	return nil
}
