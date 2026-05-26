package service

import (
	"errors"
	"strings"

	dto "github.com/gabrielribeirofsouza/OpsCenter-API/DTO"
	"github.com/gabrielribeirofsouza/OpsCenter-API/entity"
	"github.com/gabrielribeirofsouza/OpsCenter-API/repository"
)

type IncidentService interface {
	CreateIncident(incident entity.Incident) (*dto.IncidentCreateResponse, error)
	AssumeIncident(responsible entity.User, idIncident string) (*dto.IncidentAssumeResponse, error)
	FinishedIncident(idIncidente string, responsible entity.User)(*dto.IncidentUpdateResponse, error)
}
type incidentService struct {
	repo     repository.IncidentRepository
	repoUser repository.UserRepository
}

func NewIncidentService(repoIncident repository.IncidentRepository, repoUser repository.UserRepository) IncidentService {
	return &incidentService{
		repo:     repoIncident,
		repoUser: repoUser,
	}
}
func (is *incidentService) CreateIncident(incident entity.Incident) (*dto.IncidentCreateResponse, error) {
	tx, err := is.repo.BeginTx()
	if err != nil {
		return nil, err
	}
	user, err := is.repoUser.GetUser(incident.CreatedBy)
	if err != nil {
		return &dto.IncidentCreateResponse{}, err
	}

	isPermissions := strings.ToUpper(user.Permissoes)
	if isPermissions != "USER" && isPermissions != "ADMIN" {
		return &dto.IncidentCreateResponse{}, errors.New("VOCÊ NÃO TEM PERMISSÃO PARA REGISTRAR UM INCIDENTE")
	}
	defer tx.Rollback()
	resp, err := is.repo.CreateIncident(tx, incident)
		if err != nil {
			return &dto.IncidentCreateResponse{}, err
		}
	incident.ID = resp
	_, err = is.repo.CreateHistory(tx, incident)
		if err != nil{
			return &dto.IncidentCreateResponse{},errors.New("ERRO AO CRIAR HISTORICO DE INCIDENTE")
		}
	tx.Commit()
	return &dto.IncidentCreateResponse{
		Message: "Incidente registrado com sucesso",
		ID:      resp,
	}, nil
}
func (is *incidentService) AssumeIncident(responsible entity.User, idIncident string) (*dto.IncidentAssumeResponse, error) {
	tx, err := is.repo.BeginTx()
	if err != nil {
		return nil, err
	}
	incident, err := is.repo.GetIncident(idIncident)
	if err != nil {
		return nil, err
	}
	if incident.AssignedTo != "" {
		return nil, errors.New("INCIDENTE JÁ ASSUMIDO")
	}
	if incident.Status != "ABERTO" {
		return nil, errors.New("não é possivel assumir o incidente com o status atual")
	}
	defer tx.Rollback()
	_, err = is.repo.AssumeIncident(tx, responsible, idIncident)
	if err != nil {
		return nil, err
	}
	history, err := is.repo.GetHistory(*incident)
	_,err = is.repo.UpdateHistory(tx, *history, "ABERTO", "ASSUMIDO", responsible)
		if err != nil{
		return nil, err
		}
	tx.Commit()
	return &dto.IncidentAssumeResponse{
		Message:         "Incidente assumido com sucesso",
		NameResponsable: responsible.ID,
	}, nil
}

func(is *incidentService)FinishedIncident(idIncidente string, responsible entity.User)(*dto.IncidentUpdateResponse, error){
	tx, err := is.repo.BeginTx()
		if err != nil {
			return nil, err
		}
	incidente, err := is.repo.GetIncident(idIncidente)
		if err != nil{
			return nil, err
		}
		if incidente.Status != "ASSUMIDO" {
			return nil, errors.New("não é possivel finalizar o chamado com o status atual")
		}
	defer tx.Rollback()
	resp, err := is.repo.FinishedIncident(tx, idIncidente)
	if err != nil{
		return nil, err
	}
	history, err := is.repo.GetHistory(*incidente)
	_, err = is.repo.UpdateHistory(tx, *history, "ASSUMIDO", "FINALIZADO", responsible)
		if err != nil{
			return nil, err
		}
	tx.Commit()
	return &dto.IncidentUpdateResponse{
		Message: "incidente finalizado com sucesso",
		ID: resp,
	}, nil
}