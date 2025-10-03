package model

import (
	"encoding/json"
	"testing"
)

func TestElectionModel(t *testing.T) {
	t.Run("should create election with all fields", func(t *testing.T) {
		election := Election{
			Id:   1,
			Name: "Presidential Election 2024",
		}

		if election.Id != 1 {
			t.Errorf("Expected Id 1, got %d", election.Id)
		}
		if election.Name != "Presidential Election 2024" {
			t.Errorf("Expected Name 'Presidential Election 2024', got %s", election.Name)
		}
	})

	t.Run("should handle empty election", func(t *testing.T) {
		var election Election

		if election.Id != 0 {
			t.Errorf("Expected zero Id, got %d", election.Id)
		}
		if election.Name != "" {
			t.Errorf("Expected empty Name, got %s", election.Name)
		}
	})

	t.Run("should marshal to JSON correctly", func(t *testing.T) {
		election := Election{
			Id:   42,
			Name: "Senate Election",
		}

		jsonData, err := json.Marshal(election)
		if err != nil {
			t.Fatalf("Failed to marshal election: %v", err)
		}

		expectedJSON := `{"id":42,"name":"Senate Election"}`
		if string(jsonData) != expectedJSON {
			t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
		}
	})

	t.Run("should unmarshal from JSON correctly", func(t *testing.T) {
		jsonData := `{"id":123,"name":"Local Election"}`

		var election Election
		err := json.Unmarshal([]byte(jsonData), &election)
		if err != nil {
			t.Fatalf("Failed to unmarshal election: %v", err)
		}

		if election.Id != 123 {
			t.Errorf("Expected Id 123, got %d", election.Id)
		}
		if election.Name != "Local Election" {
			t.Errorf("Expected Name 'Local Election', got %s", election.Name)
		}
	})

	t.Run("should handle partial JSON", func(t *testing.T) {
		jsonData := `{"name":"Partial Election"}`

		var election Election
		err := json.Unmarshal([]byte(jsonData), &election)
		if err != nil {
			t.Fatalf("Failed to unmarshal election: %v", err)
		}

		if election.Id != 0 {
			t.Errorf("Expected Id 0, got %d", election.Id)
		}
		if election.Name != "Partial Election" {
			t.Errorf("Expected Name 'Partial Election', got %s", election.Name)
		}
	})

	t.Run("should handle invalid JSON gracefully", func(t *testing.T) {
		jsonData := `{"id":"not-a-number","name":"Invalid Election"}`

		var election Election
		err := json.Unmarshal([]byte(jsonData), &election)
		if err == nil {
			t.Error("Expected error for invalid JSON, got nil")
		}
	})
}
