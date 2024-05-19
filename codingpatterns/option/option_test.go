package option

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// Test case 1: Default server
	server := NewServer()
	if server.host != "localhost" {
		t.Errorf("Expected host to be 'localhost', but got '%s'", server.host)
	}
	if server.port != 8080 {
		t.Errorf("Expected port to be 8080, but got %d", server.port)
	}
	if server.timeout != time.Second*5 {
		t.Errorf("Expected timeout to be 5 seconds, but got %s", server.timeout)
	}

	// Test case 2: Custom server with options
	server = NewServer(
		WitHost("example.com"),
		WithPort(9000),
		WithTimeout(time.Second*10),
	)
	if server.host != "example.com" {
		t.Errorf("Expected host to be 'example.com', but got '%s'", server.host)
	}
	if server.port != 9000 {
		t.Errorf("Expected port to be 9000, but got %d", server.port)
	}
	if server.timeout != time.Second*10 {
		t.Errorf("Expected timeout to be 10 seconds, but got %s", server.timeout)
	}
}
