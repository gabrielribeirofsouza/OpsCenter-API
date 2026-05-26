package entity

import "time"

type IncidentHistory struct {
	ID string `json:"id"`

	Incident_id string `json:"incident_id"`

	Title string `json:"title"`

	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`

	ChangedBy string `json:"changed_by"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}