package data

import "github.com/aren55555/shepherd-backend/models"

type Store interface {
	GetForms() []*models.Form
	GetFormBy(string) *models.Form

	GetApplicationBy(string) *models.Application
	CreateApplicationFrom(string) (*models.Application, error)
	UpdateApplicationFormData(string, []byte) (*models.Application, error)
	CompleteApplication(string) error
}
