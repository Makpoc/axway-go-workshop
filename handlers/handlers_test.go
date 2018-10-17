package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostParser(t *testing.T) {
	// stubs, spies, preparation

	// build the request
	rBody := strings.NewReader(`{"url" : "localhost"}`)
	req, err := http.NewRequest(http.MethodPost, "/postParser", rBody)
	if err != nil {
		t.Fatalf("Failed to build request: %v", err)
	}

	// a response recorder implements the ResponseWriter interface and acts as a spy we can later inspect
	respRecorder := httptest.NewRecorder()

	// actions
	PostParser(respRecorder, req)

	// verifications
	if respRecorder.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, respRecorder.Code)
	}
}

func TestPostParser_badMethod(t *testing.T) {
	// stubs, spies, preparation

	// build the request
	rBody := strings.NewReader(`{"url" : "localhost"}`)
	req, err := http.NewRequest(http.MethodGet, "/postParser", rBody)
	if err != nil {
		t.Fatalf("Failed to build request: %v", err)
	}

	// a response recorder implements the ResponseWriter interface and acts as a spy we can later inspect
	respRecorder := httptest.NewRecorder()

	// actions
	PostParser(respRecorder, req)

	// verifications
	if respRecorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status code %d, got %d", http.StatusMethodNotAllowed, respRecorder.Code)
	}
}
func TestPostParser_badBody(t *testing.T) {
	// stubs, spies, preparation

	// build the request
	rBody := strings.NewReader(`invalidJson`)
	req, err := http.NewRequest(http.MethodPost, "/postParser", rBody)
	if err != nil {
		t.Fatalf("Failed to build request: %v", err)
	}

	// a response recorder implements the ResponseWriter interface and acts as a spy we can later inspect
	respRecorder := httptest.NewRecorder()

	// actions
	PostParser(respRecorder, req)

	// verifications
	if respRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, respRecorder.Code)
	}
}
