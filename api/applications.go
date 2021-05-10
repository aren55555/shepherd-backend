package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type createApplicationRequest struct {
	SourceFormID string `json:"source_form_id"`
}

func (r createApplicationRequest) validate() error {
	if r.SourceFormID == "" {
		return errors.New("request was missing form id")
	}
	return nil
}

func (h *Handler) getApplication(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	applicationID := params.ByName(paramID)
	application := h.dataStore.GetApplicationBy(applicationID)
	if application == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set(headerContentType, mimeJSON)
	json.NewEncoder(w).Encode(application)
}

func (h *Handler) createApplication(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	requestData := createApplicationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "could not decode request JSON", http.StatusBadRequest)
		return
	}
	if err := requestData.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	application, err := h.dataStore.CreateApplicationFrom(requestData.SourceFormID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(headerContentType, mimeJSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(application)
}

func (h *Handler) updateApplication(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	applicationID := params.ByName(paramID)

	defer r.Body.Close()
	jsonData, _ := ioutil.ReadAll(r.Body)

	application, err := h.dataStore.UpdateApplicationFormData(applicationID, jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(headerContentType, mimeJSON)
	json.NewEncoder(w).Encode(application)
}

func (h *Handler) completeApplication(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	applicationID := params.ByName(paramID)

	if err := h.dataStore.CompleteApplication(applicationID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// OK!
}

func (h *Handler) getApplicationPDF(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	applicationID := params.ByName(paramID)

	application := h.dataStore.GetApplicationBy(applicationID)
	if application == nil || !application.IsComplete {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set(headerContentType, mimePDF)
	pdf, err := application.PDF()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pdf.Write(w)
}
