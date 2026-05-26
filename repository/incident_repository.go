package repository

import (
	"database/sql"
	"time"

	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
)

type IncidentRepository interface {
	CreateIncident(incident entity.Incident) (string, error)
	AssumeIncident(responsible entity.User, idIncident string) (string, error)
	FinishedIncident(idIncident string, responsible entity.User) (string, error)
	GetIncident(idIncident string) (*entity.Incident, error)

	CreateHistory(incident entity.Incident) (string, error)
	UpdateHistory(history entity.IncidentHistory, oldVal, newVal string, responsible entity.User) (string, error)
	GetHistory(incident entity.Incident) (*entity.IncidentHistory, error)
}
type incidentRepository struct {
	db *sql.DB
}

func NewIncidentRepository(db *sql.DB) IncidentRepository {
	return &incidentRepository{
		db: db,
	}
}
func (ir *incidentRepository) CreateIncident(incident entity.Incident) (string, error) {
	var id string
	formattedHours := time.Now().Format("2006-01-02 15:04:05")
	err := ir.db.QueryRow(
		`INSERT INTO "incident"
		(title, description, status, priority, severity, category_id, created_by, team_id, sla_due_at, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`, incident.Title, incident.Description, "ABERTO", incident.Priority, incident.Severity, incident.Category_id, incident.CreatedBy, incident.Team_id, incident.SLADue_at, formattedHours).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, err
}
func (ir *incidentRepository) AssumeIncident(responsible entity.User, idIncident string) (string, error) {
	err := ir.db.QueryRow(`
	UPDATE "incident"
	SET assigned_to=$1 AND updated_at=$2
	WHERE id=$3
	`, responsible.ID, time.Now().Format("2006-01-02 15:04:05"), idIncident).Scan()
	if err != nil {
		return "", err
	}
	return responsible.ID, nil
}
func (ir *incidentRepository) GetIncident(idIncident string) (*entity.Incident, error) {
	var incidentResponse entity.Incident
	err := ir.db.QueryRow(`
	SELECT * FROM "incident"
	WHERE id=$1`, idIncident).Scan(
		&incidentResponse.ID,
		&incidentResponse.Title,
		&incidentResponse.Status,
		&incidentResponse.Team_id,
		&incidentResponse.Severity,
		&incidentResponse.Priority,
		&incidentResponse.SLADue_at,
		&incidentResponse.Category_id,
		&incidentResponse.AssignedTo,
		&incidentResponse.CreatedBy,
		&incidentResponse.Closed_at,
		&incidentResponse.Resolved_at,
		&incidentResponse.Updated_at,
		&incidentResponse.Description,
	)
	if err != nil {
		return nil, err
	}
	return &incidentResponse, nil
}
func (ir *incidentRepository) CreateHistory(incident entity.Incident) (string, error) {
	var idHistory string
	err := ir.db.QueryRow(
		`INSERT INTO "incident_history"
		(title, incident_id, old_value, new_value, changed_by, creatd_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id`, incident.Title, incident.ID, "", incident.Status, "ABERTO", time.Now().Format("2006-01-02 15:04:05"),
	).Scan(&idHistory)
	if err != nil {
		return "", err
	}
	return idHistory, nil
}
func (ir *incidentRepository) GetHistory(incident entity.Incident) (*entity.IncidentHistory, error) {
	var history entity.IncidentHistory
	err := ir.db.QueryRow(
		`SELECT * FROM "incident_history"
		 WHERE incident_id= $1`, incident.ID,
	).Scan(
		&history.ID,
		&history.Incident_id,
		&history.Title,
		&history.OldValue,
		&history.NewValue,
		&history.ChangedBy,
		&history.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &history, nil
}
func (ir *incidentRepository) UpdateHistory(history entity.IncidentHistory, oldVal string, newVal string, responsible entity.User) (string, error) {
	var idUpdate string
	err := ir.db.QueryRow(
		`INSERT INTO "incident_history"
		(title, incident_id, old_value, new_value, changed_by, creatd_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id`, history.Title, history.Incident_id, oldVal, newVal, responsible.ID, time.Now().Format("2006-01-02 15:04:05"),
	).Scan(&idUpdate)
	if err != nil {
		return "", err
	}
	return idUpdate, nil
}
func (ir *incidentRepository) FinishedIncident(idIncident string, responsible entity.User) (string, error) {
	var idFinished string
	err := ir.db.QueryRow(
		`UPDATE "incident"
		 status='FINISHED' 
		 WHERE id=$1
		 RETURNING id`, idIncident,
	).Scan(&idFinished)
	if err != nil {
		return "", err
	}
	return idFinished, nil
}
