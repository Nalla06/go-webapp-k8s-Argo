package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServer(t *testing.T) {
	// Create a test server using the same file server configuration
	fs := http.FileServer(http.Dir("./static"))
	
	// Create a test server with the file server handler
	testServer := httptest.NewServer(fs)
	defer testServer.Close()

	// Test case 1: Ensure server starts without errors
	if testServer == nil {
		t.Fatal("Test server could not be created")
	}

	// Test case 2: Check if server responds with 200 OK for existing files
	// Note: You'll need to have a file like index.html or test.txt in your ./static directory
	resp, err := http.Get(testServer.URL + "/index.html")
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test case 3: Check response for non-existent file
	resp, err = http.Get(testServer.URL + "/nonexistent.txt")
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent file, got %d", resp.StatusCode)
	}
}

func TestServerPort(t *testing.T) {
	// Create a mock listener to check port binding
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// Extract the actual port used
	_, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatalf("Failed to get port: %v", err)
	}

	// Verify port is available
	if port == "" {
		t.Error("No port was assigned")
	}
}