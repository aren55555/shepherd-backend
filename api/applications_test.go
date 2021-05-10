package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aren55555/shepherd-backend/data"
	"github.com/aren55555/shepherd-backend/data/mock"
	"github.com/aren55555/shepherd-backend/models"
)

func Test_GetApplication(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		mock           data.Store
		expectedStatus int
	}{
		{
			desc: "application_not_found",
			mock: &mock.Client{
				Application: nil,
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			desc: "success",
			mock: &mock.Client{
				Application: &models.Application{},
			},
			expectedStatus: http.StatusOK,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/applications/1", nil)
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
		})
	}
}

func Test_CreateApplication(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		mock           data.Store
		requestData    createApplicationRequest
		expectedStatus int
	}{
		{
			desc:           "missing_source_form_id",
			requestData:    createApplicationRequest{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			desc: "could_not_create",
			requestData: createApplicationRequest{
				SourceFormID: "1",
			},
			mock: &mock.Client{
				Application: nil,
				Error:       errors.New("failure"),
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			desc: "success",
			requestData: createApplicationRequest{
				SourceFormID: "1",
			},
			mock: &mock.Client{
				Application: &models.Application{},
				Error:       nil,
			},
			expectedStatus: http.StatusCreated,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			requestData, _ := json.Marshal(tc.requestData)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/applications", bytes.NewReader(requestData))
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
		})
	}
}

func Test_UpdateApplication(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		mock           data.Store
		expectedStatus int
	}{
		{
			desc: "error_updating",
			mock: &mock.Client{
				Application: nil,
				Error:       errors.New("invalid"),
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			desc: "success",
			mock: &mock.Client{
				Application: &models.Application{},
				Error:       nil,
			},
			expectedStatus: http.StatusOK,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			req := httptest.NewRequest(http.MethodPut, "/api/v1/applications/1", nil)
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
		})
	}
}

func Test_CompleteApplication(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		mock           data.Store
		expectedStatus int
	}{
		{
			desc: "invalid",
			mock: &mock.Client{
				Error: errors.New("invalid"),
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			desc: "success",
			mock: &mock.Client{
				Error: nil,
			},
			expectedStatus: http.StatusNoContent,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/applications/1/complete", nil)
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
		})
	}
}

// func Test_GetApplicationPDF(t *testing.T) {
// 	for _, tc := range []struct {
// 		desc           string
// 		mock           data.Store
// 		expectedStatus int
// 	}{
// 		{
// 			desc: "invalid",
// 			mock: &mock.Client{
// 				Error: errors.New("invalid"),
// 			},
// 			expectedStatus: http.StatusInternalServerError,
// 		},
// 		{
// 			desc: "success",
// 			mock: &mock.Client{
// 				Error: nil,
// 			},
// 			expectedStatus: http.StatusNoContent,
// 		},
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			h := New(tc.mock)

// 			req := httptest.NewRequest(http.MethodGet, "/api/v1/applications/1/pdf", nil)
// 			rr := httptest.NewRecorder()

// 			h.getApplicationPDF(rr, req)

// 			// Verify HTTP Status Code
// 			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
// 				t.Fatalf("http status got %v, want %v", got, want)
// 			}
// 		})
// 	}
// }
