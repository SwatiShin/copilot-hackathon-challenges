package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchScrollContent(t *testing.T) {
	// Create a test server with a mock response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a test scroll with secrets {*secret1*} and {*secret2*}."))
	}))
	defer server.Close()

	content, err := fetchScrollContent(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedContent := "This is a test scroll with secrets {*secret1*} and {*secret2*}."
	if content != expectedContent {
		t.Errorf("Expected content %q, got %q", expectedContent, content)
	}
}

func TestExtractSecrets(t *testing.T) {
	content := "This is a test scroll with secrets {*secret1*} and {*secret2*}."
	expectedSecrets := []string{"secret1", "secret2"}

	secrets := extractSecrets(content)
	if !reflect.DeepEqual(secrets, expectedSecrets) {
		t.Errorf("Expected secrets %v, got %v", expectedSecrets, secrets)
	}
}
