package mock

import (
	"github.com/aren55555/shepherd-backend/data"
	"github.com/aren55555/shepherd-backend/models"
)

var _ data.Store = &Client{}

// Client implements the data.Store interface, should be used in testing/stubbing.
type Client struct {
	Forms       []*models.Form
	Form        *models.Form
	Application *models.Application
	Error       error
}

func (m *Client) GetForms() []*models.Form {
	return m.Forms
}

func (m *Client) GetFormBy(string) *models.Form {
	return m.Form
}

func (m *Client) GetApplicationBy(string) *models.Application {
	return m.Application
}

func (m *Client) CreateApplicationFrom(string) (*models.Application, error) {
	return m.Application, m.Error
}

func (m *Client) UpdateApplicationFormData(string, []byte) (*models.Application, error) {
	return m.Application, m.Error
}

func (m *Client) CompleteApplication(string) error {
	return m.Error
}
