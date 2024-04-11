package main

import (
	"testing"
)

func TestTestFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestFunc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Error("xiba")
		})
	}
}
