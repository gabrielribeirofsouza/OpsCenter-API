package main

import (
	"log"
	"net/http"

	"github.com/gabrielribeirofsouza/OpsCenter-API/config"
	"github.com/gabrielribeirofsouza/OpsCenter-API/controller"
	"github.com/gabrielribeirofsouza/OpsCenter-API/repository"
	"github.com/gabrielribeirofsouza/OpsCenter-API/service"
)

func main() {
	db, err := config.SetupDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	incidentRepo := repository.NewIncidentRepository(db)

	userService := service.NewUserService(userRepo)
	incidentService := service.NewIncidentService(incidentRepo, userRepo)

	userController := controller.NewAuthController(userService)
	incidentController := controller.NewIncidentController(incidentService)

	http.HandleFunc("auth/login", userController.Login)
	http.HandleFunc("auth/register", userController.Register)
	http.HandleFunc("incident/createIncident", incidentController.CreateIncident)
	http.HandleFunc("incident/assumeIncident", incidentController.AssumeIncident)
}
