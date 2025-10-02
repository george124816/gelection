package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	t.Run("should return OK on GET request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()

		HealthCheckHandler(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
		}

		expectedBody := "OK\n"
		if string(body) != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, string(body))
		}
	})

	t.Run("should return Method Not Allowed on POST request", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/health", nil)
		w := httptest.NewRecorder()

		HealthCheckHandler(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}

		expectedBody := "Method Not Allowed\n"
		if string(body) != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, string(body))
		}
	})

	t.Run("should return Method Not Allowed on PUT request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/health", nil)
		w := httptest.NewRecorder()

		HealthCheckHandler(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}

		expectedBody := "Method Not Allowed\n"
		if string(body) != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, string(body))
		}
	})

	t.Run("should return Method Not Allowed on DELETE request", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/health", nil)
		w := httptest.NewRecorder()

		HealthCheckHandler(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}

		expectedBody := "Method Not Allowed\n"
		if string(body) != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, string(body))
		}
	})
}