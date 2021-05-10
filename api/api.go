package api

import (
	"net/http"

	"github.com/aren55555/shepherd-backend/data"
	"github.com/julienschmidt/httprouter"
)

const (
	paramID = "id"

	headerContentType = "Content-Type"
	mimeJSON          = "application/json"
	mimePDF           = "application/pdf"
)

func New(ds data.Store) *Handler {
	r := httprouter.New()

	h := &Handler{
		dataStore: ds,
		Router:    r,
	}

	r.HandlerFunc(http.MethodGet, "/api/v1/forms", h.listAllForms)
	r.HandlerFunc(http.MethodGet, "/api/v1/forms/:id", h.getForm)

	r.HandlerFunc(http.MethodGet, "/api/v1/application/:id", h.getApplication)
	r.HandlerFunc(http.MethodPost, "/api/v1/application", h.createApplication)
	r.HandlerFunc(http.MethodPut, "/api/v1/application/:id", h.updateApplication)
	r.HandlerFunc(http.MethodPost, "/api/v1/application/:id/complete", h.completeApplication)
	r.HandlerFunc(http.MethodGet, "/api/v1/application/:id/pdf", h.getApplicationPDF)

	return h
}

var _ http.Handler = &Handler{}

type Handler struct {
	dataStore data.Store
	*httprouter.Router
}
