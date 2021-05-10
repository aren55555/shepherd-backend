package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) listAllForms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, mimeJSON)
	forms := h.dataStore.GetForms()
	json.NewEncoder(w).Encode(forms)
}

func (h *Handler) getForm(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	formID := params.ByName(paramID)
	form := h.dataStore.GetFormBy(formID)
	if form == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set(headerContentType, mimeJSON)
	json.NewEncoder(w).Encode(form)
}
