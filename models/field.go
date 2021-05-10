package models

import (
	"errors"
)

type FieldType string

const (
	FieldTypeText       FieldType = "text"
	FieldTypeEmail                = "email"
	FieldTypeURL                  = "url"
	FieldTypeNumber               = "number"
	FieldTypeDate                 = "date"
	FieldTypeSelectOne            = "select_one"
	FieldTypeSelectMany           = "select_many"
)

type Field struct {
	Label    string    `json:"label"` // the string that goes in the <label> for this field
	Type     FieldType `json:"type"`
	Required bool      `json:"required,omitempty"`
	// TODO: metadata to describe additional validations
}

func (f Field) validate(value interface{}) error {
	// TODO: use the Type field to type assert the value before validating it
	if value == nil && f.Required {
		return errors.New("a value is required")
	}
	return nil
}
