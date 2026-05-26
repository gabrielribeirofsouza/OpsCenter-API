package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
	"github.com/gabrielribeirofsouza/OpsCenter-API/service"
)

type IncidentController struct {
	service service.IncidentService
}

func NewIncidentController(ic service.IncidentService) *IncidentController {
	return &IncidentController{
		service: ic,
	}
}
func (ic *IncidentController) CreateIncident(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "METÓDO INVÁLIDO", http.StatusMethodNotAllowed)
	}
	if r.Body == nil {
		http.Error(w, "BODY INVÁLIDO", http.StatusBadRequest)
	}
	var body entity.Incident
	response, err := ic.service.CreateIncident(body)
	if err != nil {
		http.Error(w, "[ERROR] - Ocorreu um erro ao registrar o incidente", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func (ic *IncidentController) AssumeIncident(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "METÓDO INVÁLIDO", http.StatusMethodNotAllowed)
	}
	if r.Body == nil {
		http.Error(w, "BODY INVÁLIDO", http.StatusBadRequest)
	}
	var body entity.User
	query := r.URL.Query()
	idIncident := query.Get("idIncident")
	response, err := ic.service.AssumeIncident(body, idIncident)
	if err != nil {
		http.Error(w, "[ERROR] - Erro ao assumir incidente", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}
