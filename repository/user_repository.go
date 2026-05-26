package repository

import (
	"database/sql"

	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
)

type UserRepository interface {
	AuthUser(email string, password string) (entity.User, error)
	Register(user entity.User) (string, error)
	GetUser(id string) (entity.User, error)
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (ur *userRepository) AuthUser(email string, password string) (entity.User, error) {
	var user entity.User
	err := ur.db.QueryRow(`SELECT  id, name, status, equipe, departamento, cargo, permissoes, email
						WHERE email=$1 AND password=$2`, email, password).Scan(
		&user.ID,
		&user.Name,
		&user.Status,
		&user.Equipe_id,
		&user.Departamento_id,
		&user.Cargo,
		&user.Permissoes,
		&user.Email,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
func (ur *userRepository) Register(user entity.User) (string, error) {
	var id string
	err := ur.db.QueryRow(`INSERT INTO "user" (name, status, equipe, departamento, cargo, permissoes, email, password)
							VALUES($1, $2, $3, $4, $5, $6, $7, $8)
							ON CONFLICT(email)
							DO NOTHING
							RETURNING id`, user.Name, user.Status, user.Equipe_id, user.Departamento_id, user.Cargo, user.Permissoes, user.Email, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
func (ur *userRepository) GetUser(id string) (entity.User, error) {
	var userResponse entity.User
	err := ur.db.QueryRow(
		`SELECT id, name, status, permissions, equipe_id, department_id, cargo, email, name 
		WHERE id=$1`, id,
	).Scan(
		&userResponse.ID,
		&userResponse.Name,
		&userResponse.Status,
		&userResponse.Permissoes,
		&userResponse.Equipe_id,
		&userResponse.Departamento_id,
		&userResponse.Cargo,
		&userResponse.Email,
		&userResponse.Name,
	)
	if err != nil {
		return userResponse, err
	}
	return userResponse, err
}
