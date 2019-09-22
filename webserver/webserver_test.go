package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var errorString = "Expected %v, got %v"

func TestPrint(t *testing.T) {
	w := httptest.NewRecorder()

	Print(w, http.StatusOK, []byte("OK"))

	if w.Code != http.StatusOK {
		t.Errorf(errorString, http.StatusOK, w.Code)
	}
	if w.Body.String() != "OK" {
		t.Errorf(errorString, "OK", w.Body.String())
	}

	w = httptest.NewRecorder()

	Print(w, http.StatusInternalServerError, []byte("InternalServerError"))

	if w.Code != http.StatusInternalServerError {
		t.Errorf(errorString, http.StatusInternalServerError, w.Code)
	}
	if w.Body.String() != "InternalServerError" {
		t.Errorf(errorString, "InternalServerError", w.Body.String())
	}
}

func TestCheckSettings(t *testing.T) {
	expected := Settings{
		Host: "localhost",
		Port: 65432,
	}

	checkSettings()

	if settings != expected {
		t.Errorf(errorString, expected, settings)
	}

	settings.Host = "    "
	settings.Port = -1

	checkSettings()
	if settings != expected {
		t.Errorf(errorString, expected, settings)
	}
}
