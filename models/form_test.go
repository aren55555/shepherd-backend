package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var (
	f1 = Form{
		ID: "form1",
		Fieldsets: map[string]Fieldset{
			"company": {
				Label: "Company Info",
				Fields: map[string]Field{
					"name": {
						Label:    "Name",
						Type:     FieldTypeText,
						Required: true,
					},
					"website": {
						Label: "Website",
						Type:  FieldTypeURL,
					},
				},
			},
			"project": {
				Label: "Project Info",
				Fields: map[string]Field{
					"name": {
						Label: "Name",
						Type:  FieldTypeText,
					},
					"address": {
						Label: "Address",
						Type:  FieldTypeText,
					},
				},
			},
		},
		Fields: map[string]Field{
			"taxable": {
				Label: "Taxable?",
				Type:  FieldTypeSelectOne,
			},
		},
	}
)

func TestForm_EmptyInstance(t *testing.T) {
	emptyInstance := f1.EmptyFormInstance()
	data, _ := json.Marshal(emptyInstance)
	fmt.Println(string(data))
	// TODO: compare against expected
}

func TestForm_ValidateFormInstance_valid(t *testing.T) {
	jsonFile, err := os.Open("submission.json")
	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Fatal(err)
	}

	var fi FormInstance
	if err := json.Unmarshal(jsonData, &fi); err != nil {
		t.Fatal(err)
	}

	if err := f1.ValidateFormInstance(fi); err != nil {
		t.Fatal(err)
	}
}

func TestForm_ValidateFormInstance_invalid(t *testing.T) {
	fi := f1.EmptyFormInstance()

	if err := f1.ValidateFormInstance(fi); err == nil {
		t.Fatal("was expecting an error")
	}
}
