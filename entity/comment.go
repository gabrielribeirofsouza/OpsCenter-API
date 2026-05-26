package entity

import "time"

type Comment struct {
	ID string `json:"id"`

	Incident_id string `json:"incident_id"`

	User_id string `json:"user_id"`

	Content string `json:"content"`

	Visibility string `json:"visibility"`

	Created_at time.Time `json:"created_at"`
}