package models

type Application struct {
	ID         string       `json:"id"`
	IsComplete bool         `json:"is_completed"`
	FormSchema Form         `json:"form_schema"`
	FormData   FormInstance `json:"form_data"`
}
