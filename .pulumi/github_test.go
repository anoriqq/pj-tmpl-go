package main

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Test_githubResource_newRepository_CreatesRepository(t *testing.T) {
	tests := map[string]struct {
		owner string
		repo  string
	}{
		"標準的なリポジトリ作成": {
			owner: "anoriqq",
			repo:  "test-repo",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := pulumi.RunErr(func(ctx *pulumi.Context) error {
				githubRes := &githubResource{}
				
				repository, err := githubRes.newRepository(ctx, tt.owner, tt.repo)
				if err != nil {
					t.Errorf("newRepository() failed: %v", err)
					return err
				}

				// リポジトリが作成されたことを確認
				if repository == nil {
					t.Error("Repository should not be nil")
				}

				return nil
			}, pulumi.WithMocks("project", "stack", newMockProvider()))

			if err != nil {
				t.Fatalf("pulumi.RunErr failed: %v", err)
			}
		})
	}
}

func Test_GitHub_ReturnsValidResource(t *testing.T) {
	tests := map[string]struct{}{
		"GitHub関数が有効なリソースを返すこと": {},
	}

	for name := range tests {
		t.Run(name, func(t *testing.T) {
			result := GitHub()

			if result == nil {
				t.Error("GitHub() should not return nil")
			}
		})
	}
}