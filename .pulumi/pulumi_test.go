package main

import (
	"testing"
)

func Test_Pulumi_ReturnsValidResource(t *testing.T) {
	tests := map[string]struct{}{
		"Pulumi関数が有効なリソースを返すこと": {},
	}

	for name := range tests {
		t.Run(name, func(t *testing.T) {
			result := Pulumi()

			if result == nil {
				t.Error("Pulumi() should not return nil")
			}
		})
	}
}

