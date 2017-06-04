package main

import "testing"

func TestGreeting(t *testing.T) {
	if greeting := Greeting(); greeting != "Hello world" {
		t.Errorf("Expected 'Hello world', got %d", greeting)
	}
}