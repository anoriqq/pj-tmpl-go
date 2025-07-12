package main

import (
	"maps"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type mockProvider struct {
	config map[string]string
}

func (m mockProvider) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	// モックの戻り値を入力から生成
	outputs := make(resource.PropertyMap)
	maps.Copy(outputs, args.Inputs)
	return args.Name + "_id", outputs, nil
}

func (m mockProvider) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	switch args.Token {
	case "pulumi:pulumi:getConfiguration":
		// Return mock configuration values
		result := resource.PropertyMap{}
		for key, value := range m.config {
			result[resource.PropertyKey(key)] = resource.NewStringProperty(value)
		}
		return result, nil
	default:
		return args.Args, nil
	}
}

func newMockProvider() mockProvider {
	return mockProvider{
		config: map[string]string{
			"defaultStack":  "dev",
			"defaultBranch": "main",
		},
	}
}

