package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := strings.TrimRight(response.Body.String(), "\n"); body != "[]" {
		t.Errorf("Expected an empty table. Got %s\n", body)
	}
}



func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}


func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}