package model

import (
	"encoding/json"
	"testing"
)

func TestCandidateModel(t *testing.T) {
	t.Run("should create candidate with all fields", func(t *testing.T) {
		candidate := Candidate{
			Id:         1,
			Name:       "John Doe",
			ElectionId: 100,
		}

		if candidate.Id != 1 {
			t.Errorf("Expected Id 1, got %d", candidate.Id)
		}
		if candidate.Name != "John Doe" {
			t.Errorf("Expected Name 'John Doe', got %s", candidate.Name)
		}
		if candidate.ElectionId != 100 {
			t.Errorf("Expected ElectionId 100, got %d", candidate.ElectionId)
		}
	})

	t.Run("should handle empty candidate", func(t *testing.T) {
		var candidate Candidate

		if candidate.Id != 0 {
			t.Errorf("Expected zero Id, got %d", candidate.Id)
		}
		if candidate.Name != "" {
			t.Errorf("Expected empty Name, got %s", candidate.Name)
		}
		if candidate.ElectionId != 0 {
			t.Errorf("Expected zero ElectionId, got %d", candidate.ElectionId)
		}
	})

	t.Run("should marshal to JSON correctly", func(t *testing.T) {
		candidate := Candidate{
			Id:         42,
			Name:       "Jane Smith",
			ElectionId: 2024,
		}

		jsonData, err := json.Marshal(candidate)
		if err != nil {
			t.Fatalf("Failed to marshal candidate: %v", err)
		}

		expectedJSON := `{"id":42,"name":"Jane Smith","election_id":2024}`
		if string(jsonData) != expectedJSON {
			t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
		}
	})

	t.Run("should unmarshal from JSON correctly", func(t *testing.T) {
		jsonData := `{"id":123,"name":"Test Candidate","election_id":456}`

		var candidate Candidate
		err := json.Unmarshal([]byte(jsonData), &candidate)
		if err != nil {
			t.Fatalf("Failed to unmarshal candidate: %v", err)
		}

		if candidate.Id != 123 {
			t.Errorf("Expected Id 123, got %d", candidate.Id)
		}
		if candidate.Name != "Test Candidate" {
			t.Errorf("Expected Name 'Test Candidate', got %s", candidate.Name)
		}
		if candidate.ElectionId != 456 {
			t.Errorf("Expected ElectionId 456, got %d", candidate.ElectionId)
		}
	})

	t.Run("should handle partial JSON", func(t *testing.T) {
		jsonData := `{"name":"Partial Candidate"}`

		var candidate Candidate
		err := json.Unmarshal([]byte(jsonData), &candidate)
		if err != nil {
			t.Fatalf("Failed to unmarshal candidate: %v", err)
		}

		if candidate.Id != 0 {
			t.Errorf("Expected Id 0, got %d", candidate.Id)
		}
		if candidate.Name != "Partial Candidate" {
			t.Errorf("Expected Name 'Partial Candidate', got %s", candidate.Name)
		}
		if candidate.ElectionId != 0 {
			t.Errorf("Expected ElectionId 0, got %d", candidate.ElectionId)
		}
	})
}