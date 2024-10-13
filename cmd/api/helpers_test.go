package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eynopv/image-scaler/internal/validator"
)

func TestWriteJson(t *testing.T) {
	app := application{}
	data := map[string]any{"key": "value"}

	rr := httptest.NewRecorder()
	status := http.StatusOK

	err := app.writeJson(rr, status, data)
	if err != nil {
		t.Fatalf("writeJson() error = %v", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected content type application/json, got %s", contentType)
	}

	expectedBody, _ := json.Marshal(data)
	if rr.Body.String() != string(expectedBody) {
		t.Errorf("expected body %s, got %s", expectedBody, rr.Body.String())
	}
}

func TestReadInt(t *testing.T) {
	testCases := []struct {
		name         string
		query        url.Values
		key          string
		defaultValue int
		expected     int
		expectedMsg  string
	}{
		{
			name:         "valid",
			query:        url.Values{"key": []string{"20"}},
			key:          "key",
			defaultValue: 10,
			expected:     20,
			expectedMsg:  "",
		},
		{
			name:         "empty",
			query:        url.Values{},
			key:          "key",
			defaultValue: 10,
			expected:     10,
			expectedMsg:  "",
		},
		{
			name:         "invalid",
			query:        url.Values{"key": []string{"invalid"}},
			key:          "key",
			defaultValue: 10,
			expected:     10,
			expectedMsg:  "strconv.Atoi: parsing \"invalid\": invalid syntax",
		},
	}

	app := application{}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			validator := validator.NewValidator()
			result := app.readInt(testCase.query, testCase.key, testCase.defaultValue, validator)

			if result != testCase.expected {
				t.Errorf("expected %d got %d; app.readInt()", testCase.expected, result)
			}

			if validator.Message != testCase.expectedMsg {
				t.Errorf("expected %s got %s; Validator.Message", testCase.expectedMsg, validator.Message)
			}
		})
	}
}
