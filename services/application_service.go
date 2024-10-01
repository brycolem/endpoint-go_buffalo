package services

import (
	"errors"
	"go_buffalo/models"

	"github.com/gobuffalo/pop/v6"
)

type ApplicationService struct {
	DB *pop.Connection
}

func NewApplicationService(db *pop.Connection) *ApplicationService {
	return &ApplicationService{
		DB: db,
	}
}

func (s *ApplicationService) GetAllApplications() ([]models.Application, error) {
	var applications []models.Application
	if err := s.DB.Eager("Notes").All(&applications); err != nil {
		return nil, errors.New("failed to fetch applications")
	}
	return applications, nil
}

func (s *ApplicationService) GetApplicationByID(id string) (*models.Application, error) {
	application := &models.Application{}
	if err := s.DB.Find(application, id); err != nil {
		return nil, errors.New("application not found")
	}
	return application, nil
}
