package main

import (
	"testing"
	"time"
)

func TestTXTRecordManagement(t *testing.T) {
	// Test configuration
	bunnyCfg := &bunnyClientConfig{
		apiKey: "test-api-key", // Use a test API key
	}
	domain := "_acme-challenge.example.com."
	txtValue := "937g1873bef3bf032"

	// Test adding TXT record
	t.Run("Add TXT Record", func(t *testing.T) {
		err := addTxtRecord(bunnyCfg, domain, txtValue)
		if err != nil {
			t.Errorf("Failed to add TXT record: %v", err)
		}
	})

	// Give DNS some time to propagate
	time.Sleep(1 * time.Second)

	// Test deleting TXT record
	t.Run("Delete TXT Record", func(t *testing.T) {
		err := deleteTxtRecord(bunnyCfg, domain, txtValue)
		if err != nil {
			t.Errorf("Failed to delete TXT record: %v", err)
		}
	})
}
