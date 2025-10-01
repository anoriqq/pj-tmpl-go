package pkg

import (
	"maps"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MockProvider Pulumiのモックプロバイダー。
type MockProvider struct {
	config map[string]string
}

// NewMockProvider モックプロバイダーを生成する。
func NewMockProvider() MockProvider {
	return MockProvider{
		config: map[string]string{
			"project:defaultStack":  "dev",
			"project:defaultBranch": "main",
		},
	}
}

// Call モックの呼び出しを処理する。
func (m MockProvider) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

// NewResource モックのリソース生成を処理する。
func (m MockProvider) NewResource(
	args pulumi.MockResourceArgs,
) (string, resource.PropertyMap, error) {
	// モックの戻り値を入力から生成
	outputs := make(resource.PropertyMap)
	maps.Copy(outputs, args.Inputs)

	return args.Name + "_id", outputs, nil
}
