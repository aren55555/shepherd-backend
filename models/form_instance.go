package models

import (
	"errors"
)

// FormInstance models a JSON instance of [Form].
type FormInstance map[string]interface{}

func (fi FormInstance) ID() (string, error) {
	idI, ok := fi[_keyID]
	if !ok {
		return "", errors.New("missing ID")
	}
	idS, ok := idI.(string)
	if !ok {
		return "", errors.New("invalid ID")
	}
	return idS, nil
}

func (fi FormInstance) Fieldsets() map[string]map[string]interface{} {
	fieldsetsI, ok := fi[_keyFieldsets]
	if !ok {
		return nil
	}

	fieldsets, ok := fieldsetsI.(map[string]interface{})
	if !ok {
		return nil
	}

	data := map[string]map[string]interface{}{}
	for k, v := range fieldsets {
		vm, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		data[k] = vm
	}
	return data
}

func (fi FormInstance) Fields() map[string]interface{} {
	fieldI, ok := fi[_keyFields]
	if !ok {
		return nil
	}
	fields, ok := fieldI.(map[string]interface{})
	if !ok {
		return nil
	}
	return fields
}
