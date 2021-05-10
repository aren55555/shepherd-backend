package models

// Fieldset models a named collection of HTML input fields.
type Fieldset struct {
	Label  string           `json:"label"`
	Fields map[string]Field `json:"fields"`
}

func (f Fieldset) emptyInstance() map[string]interface{} {
	data := map[string]interface{}{}
	for k := range f.Fields {
		data[k] = nil
	}
	return data
}
