package entity

import "time"

type Incident struct {
	ID string `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Status   string `json:"status"`
	Priority string `json:"priority"`
	Severity string `json:"severity"`

	Category_id string `json:"category_id"`

	CreatedBy string `json:"created_by"`
	AssignedTo string `json:"assigned_to"`

	Team_id string `json:"team_id"`

	SLADue_at *time.Time `json:"sla_due_at"`

	Resolved_at *time.Time `json:"resolved_at"`
	Closed_at   *time.Time `json:"closed_at"`

	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}