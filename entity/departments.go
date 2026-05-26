package entity

import "time"

type Departamentos struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code_team"`
	Created_at time.Time `json:"created_at"`
}