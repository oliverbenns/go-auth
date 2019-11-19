package main

import (
	"testing"
)

func TestLoadTemplate(t *testing.T) {
	template := LoadTemplate("index.html")

	if template == nil {
		t.Error("Failed to load template")
	}
}
