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

	r.HandlerFunc(http.MethodPost, "/api/v1/applications", h.createApplication)
	r.HandlerFunc(http.MethodGet, "/api/v1/applications/:id", h.getApplication)
	r.HandlerFunc(http.MethodPut, "/api/v1/applications/:id", h.updateApplication)
	r.HandlerFunc(http.MethodPost, "/api/v1/applications/:id/complete", h.completeApplication)
	r.HandlerFunc(http.MethodGet, "/api/v1/applications/:id/pdf", h.getApplicationPDF)

	return h
}

var _ http.Handler = &Handler{}

type Handler struct {
	dataStore data.Store
	*httprouter.Router
}
