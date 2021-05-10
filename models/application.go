package models

import (
	"fmt"

	"github.com/signintech/gopdf"
)

type Application struct {
	ID         string       `json:"id"`
	IsComplete bool         `json:"is_completed"`
	FormSchema Form         `json:"form_schema"`
	FormData   FormInstance `json:"form_data"`
}

func (a Application) PDF() (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("Ubuntu-Regular", "Ubuntu-Regular.ttf")
	if err != nil {
		return nil, err
	}

	err = pdf.SetFont("Ubuntu-Regular", "", 18)
	if err != nil {
		return nil, err
	}
	pdf.SetX(10.0)
	pdf.SetY(10.0)
	pdf.Cell(nil, fmt.Sprintf("Application %s", a.ID))

	err = pdf.SetFont("Ubuntu-Regular", "", 14)
	if err != nil {
		return nil, err
	}
	pdf.SetX(10.0)
	pdf.SetY(10.0 + 18.0)
	pdf.Cell(nil, "Completed")

	y := 10.0 + 18.0 + 18.0

	err = pdf.SetFont("Ubuntu-Regular", "", 12)
	if err != nil {
		return nil, err
	}

	for fieldName, value := range a.FormData.Fields() {
		y = y + 18.0
		pdf.SetY(y)
		pdf.SetX(10.0)
		pdf.Cell(nil, fmt.Sprintf("%s: %v", fieldName, value))
	}

	y = y + 18.0

	for fieldsetName, fieldMap := range a.FormData.Fieldsets() {
		y = y + 9.0
		for fieldName, value := range fieldMap {
			y = y + 18.0
			pdf.SetY(y)
			pdf.SetX(10.0)
			pdf.Cell(nil, fmt.Sprintf("%s.%s: %v", fieldsetName, fieldName, value))
		}
	}

	return pdf, nil
}
