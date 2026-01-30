package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	Health(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	body := rr.Body.String()
	if body != `{"status":"healthy","service":"order-service"}` {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestListOrders(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/orders", nil)
	rr := httptest.NewRecorder()

	ListOrders(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	body := rr.Body.String()
	if body != `{"orders":[],"total":0}` {
		t.Errorf("unexpected body: %s", body)
	}
}
