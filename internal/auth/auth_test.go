package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		err      error
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			expected: "abc123",
			err:      nil,
		},
		{
			name: "missing header",
			headers: http.Header{
				"Authorization": []string{},
			},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			// Check error
			if tt.err == nil && err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if tt.err != nil && err == nil {
				t.Fatalf("expected error %v, got nil", tt.err)
			}
			if tt.err != nil && err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}

			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
