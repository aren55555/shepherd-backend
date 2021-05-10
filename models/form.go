package models

import "fmt"

const (
	_keyID        = "id"
	_keyFieldsets = "fieldsets"
	_keyFields    = "fields"
)

// From models the structure of an HTML Form
type Form struct {
	ID        string              `json:"id"`
	Fieldsets map[string]Fieldset `json:"fieldsets"`
	Fields    map[string]Field    `json:"fields"`
}

// EmptyFormInstance returns a FormInstance where all values are set to nil/null.
func (f Form) EmptyFormInstance() FormInstance {
	data := map[string]interface{}{}

	fieldSets := map[string]interface{}{}
	for k, v := range f.Fieldsets {
		fieldSets[k] = v.emptyInstance()
	}
	if len(fieldSets) > 0 {
		data[_keyFieldsets] = fieldSets
	}

	fields := map[string]interface{}{}
	for k, _ := range f.Fields {
		fields[k] = nil
	}
	if len(fields) > 0 {
		data[_keyFields] = fields
	}

	return data
}

func (f Form) ValidateFormInstance(jsonFormInstance FormInstance) error {
	// Handle Fieldsets
	for k, v := range jsonFormInstance.Fieldsets() {
		formFieldset, ok := f.Fieldsets[k]
		if !ok {
			return fmt.Errorf("fieldset %q not present in form definition", k)
		}

		for fk, fv := range v {
			formField, ok := formFieldset.Fields[fk]
			if !ok {
				return fmt.Errorf("field %q not present in form's fieldset %q definition", fk, k)
			}
			if err := formField.validate(fv); err != nil {
				return fmt.Errorf("%s.%s.%s: %s", _keyFieldsets, k, fk, err.Error())
			}
		}
	}

	// Handle Fields
	for k, v := range jsonFormInstance.Fields() {
		formField, ok := f.Fields[k]
		if !ok {
			return fmt.Errorf("field %q not present in form definition", k)
		}
		if err := formField.validate(v); err != nil {
			return fmt.Errorf("%s.%s: %s", _keyFields, k, err.Error())
		}
	}

	return nil
}
