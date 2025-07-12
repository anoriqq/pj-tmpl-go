package main

import (
	"testing"
)

// 設定に依存しない基本的なテストのみ
func Test_FactoryFunctions_ReturnValidResources(t *testing.T) {
	tests := map[string]struct {
		factory func() any
		name    string
	}{
		"GitHub関数が有効なリソースを返すこと": {
			factory: func() any { return GitHub() },
			name:    "GitHub",
		},
		"Pulumi関数が有効なリソースを返すこと": {
			factory: func() any { return Pulumi() },
			name:    "Pulumi",
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			result := tt.factory()

			if result == nil {
				t.Errorf("%s() should not return nil", tt.name)
			}
		})
	}
}

