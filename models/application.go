package models

import "encoding/json"

type Application struct {
	ID                 string `json:"id"`
	IsComplete         bool   `json:"is_completed"`
	FormDefinitionJSON []byte `json:"-"`
	FormInstanceJSON   []byte `json:"-"`
}

func (a *Application) Form() Form {
	f := Form{}
	json.Unmarshal(a.FormDefinitionJSON, &f)
	return f
}
