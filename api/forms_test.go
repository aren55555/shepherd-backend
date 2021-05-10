package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aren55555/shepherd-backend/data"
	"github.com/aren55555/shepherd-backend/data/mock"
	"github.com/aren55555/shepherd-backend/models"
)

type responseValidator func(*http.Response) error

func formsListResponseValidationFactory(expectedLength int) responseValidator {
	return responseValidator(func(resp *http.Response) error {
		defer resp.Body.Close()
		forms := []models.Form{}
		if err := json.NewDecoder(resp.Body).Decode(&forms); err != nil {
			return err
		}
		if got, want := len(forms), expectedLength; got != want {
			return fmt.Errorf("forms JSON length: got %v, want %v", got, want)
		}
		return nil
	})
}

func Test_ListForms(t *testing.T) {
	for _, tc := range []struct {
		desc                 string
		mock                 data.Store
		expectedStatus       int
		responseVerification func(*http.Response) error
	}{
		{
			desc: "no_forms",
			mock: &mock.Client{
				Forms: []*models.Form{},
			},
			expectedStatus:       http.StatusOK,
			responseVerification: formsListResponseValidationFactory(0),
		},
		{
			desc: "has_forms",
			mock: &mock.Client{
				Forms: []*models.Form{
					{
						ID: "1",
					},
					{
						ID: "2",
					},
				},
			},
			expectedStatus:       http.StatusOK,
			responseVerification: formsListResponseValidationFactory(2),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/forms", nil)
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
			if tc.responseVerification != nil {
				if err := tc.responseVerification(rr.Result()); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func Test_GetForms(t *testing.T) {
	for _, tc := range []struct {
		desc           string
		mock           data.Store
		expectedStatus int
	}{
		{
			desc: "form_not_found",
			mock: &mock.Client{
				Form: nil,
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			desc: "success",
			mock: &mock.Client{
				Form: &models.Form{},
			},
			expectedStatus: http.StatusOK,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			h := New(tc.mock)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/forms/1", nil)
			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			// Verify HTTP Status Code
			if got, want := rr.Result().StatusCode, tc.expectedStatus; got != want {
				t.Fatalf("http status got %v, want %v", got, want)
			}
		})
	}
}
