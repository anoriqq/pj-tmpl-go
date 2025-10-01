package pkg_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pkg"
)

func TestGithub_NewRepository(t *testing.T) {
	t.Parallel()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		got, err := pkg.GitHub().NewRepository(ctx)
		if err != nil {
			t.Error(err)
		}

		t.Run("リポジトリ名にはプロジェクト名が使われる", func(t *testing.T) {
			t.Parallel()

			pulumi.All(got.URN(), got.Name).ApplyT(func(all []any) error {
				urn, ok := all[0].(pulumi.URN)
				if !ok {
					t.Fatal()
				}

				name, ok := all[1].(string)
				if !ok {
					t.Fatal()
				}

				if name != ctx.Project() {
					t.Errorf("URN=%sのリポジトリ名はプロジェクト名であるはずが %s を得た", urn, name)
				}

				return nil
			})
		})

		return nil
	},
		pulumi.WithMocks("project", "dev", pkg.NewMockProvider()),
		func(ri *pulumi.RunInfo) {
			ri.Config = map[string]string{
				"project:defaultStack":  "dev",
				"project:defaultBranch": "main",
			}
		},
	)
	if err != nil {
		t.Fatalf("pulumi.RunErr failed: %v", err)
	}
}
