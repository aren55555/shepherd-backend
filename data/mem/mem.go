package mem

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/aren55555/shepherd-backend/data"
	"github.com/aren55555/shepherd-backend/models"
	"github.com/google/uuid"
)

var _ data.Store = &Client{}

type Client struct {
	*client
}

type client struct {
	lock             sync.RWMutex
	formsByID        map[string]*models.Form
	applicationsByID map[string]*models.Application
}

func New() *Client {
	return &Client{
		&client{
			formsByID:        map[string]*models.Form{},
			applicationsByID: map[string]*models.Application{},
		},
	}
}

func (c *Client) Seed(fileLocation string) error {
	if fileLocation == "" {
		return errors.New("no file location specified")
	}

	jsonData, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return err
	}

	seedForms := []*models.Form{}
	if err := json.Unmarshal(jsonData, &seedForms); err != nil {
		return err
	}

	c.seed(seedForms)
	return nil
}

func (c *client) GetForms() []*models.Form {
	c.lock.RLock()
	defer c.lock.RUnlock()

	forms := []*models.Form{}
	for _, v := range c.formsByID {
		forms = append(forms, v)
	}
	return forms
}

func (c *client) GetFormBy(formID string) *models.Form {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.getFormBy(formID)
}

func (c *client) getFormBy(formID string) *models.Form {
	form, ok := c.formsByID[formID]
	if !ok {
		return nil
	}
	return form
}

func (c *client) GetApplicationBy(applicationID string) *models.Application {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.getApplicationBy(applicationID)
}

func (c *client) getApplicationBy(applicationID string) *models.Application {
	application, ok := c.applicationsByID[applicationID]
	if !ok {
		return nil
	}
	return application
}

func (c *client) CreateApplicationFrom(formID string) (*models.Application, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	form := c.getFormBy(formID)
	if form == nil {
		return nil, fmt.Errorf("form %q not found", formID)
	}

	application := &models.Application{
		ID:         uuid.NewString(),
		FormSchema: *form,
	}
	c.applicationsByID[application.ID] = application

	return application, nil
}

func (c *client) UpdateApplicationFormData(applicationID string, formData []byte) (*models.Application, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	application := c.getApplicationBy(applicationID)
	if application == nil {
		return nil, fmt.Errorf("application %q not found", applicationID)
	}
	if application.IsComplete {
		return nil, fmt.Errorf("application already complete")
	}

	var fi models.FormInstance
	if err := json.Unmarshal(formData, &fi); err != nil {
		return nil, err
	}
	if err := application.FormSchema.ValidateFormInstance(fi); err != nil {
		return nil, err
	}

	// OK to update!
	application.FormData = fi

	return application, nil
}

func (c *client) CompleteApplication(applicationID string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	application := c.getApplicationBy(applicationID)
	if application == nil {
		return fmt.Errorf("application %q not found", applicationID)
	}

	application.IsComplete = true
	return nil
}

func (c *client) seed(forms []*models.Form) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, f := range forms {
		c.formsByID[f.ID] = f
	}
}
