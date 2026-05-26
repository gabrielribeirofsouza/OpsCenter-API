package entity

import "time"

type User struct {
	ID              string
	Status          string
	Permissoes      string
	Equipe_id       string
	Departamento_id string
	Cargo           string
	Email           string
	Name            string
	Password        string
	Created_at      time.Time
	Deleted_at 		time.Time
	Updated_at 		time.Time
}