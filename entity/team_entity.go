package entity

import "time"

type Team struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Departamento_id string `json:"department_id"`
	Created_at      time.Time `json:"created_at"`
}