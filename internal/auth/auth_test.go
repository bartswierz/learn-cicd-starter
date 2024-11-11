package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey my-secret-key"},
	}

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if apiKey != "my-secret-key" {
		t.Errorf("expected API key to be 'my-secret-key', got '%s'", apiKey)
	}
}
