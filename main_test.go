package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	server := httptest.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "warn" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	client := server.Client()

	if code, _, err := ping(client, "https://pass"); code != http.StatusOK || err != nil {
		t.Fatalf("Expected 200, got status code %d instead (%v)", code, err)
	}

	if code, _, err := ping(client, "https://warn"); code != http.StatusInternalServerError || err != nil {
		t.Fatalf("Expected 500, got status code %d instead (%v)", code, err)
	}

	if code, _, err := ping(client, "://fail"); err == nil {
		t.Fatalf("Expected error, got status code %d instead", code)
	}
}
