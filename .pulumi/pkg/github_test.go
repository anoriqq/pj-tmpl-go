package pkg_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pkg"
)

func TestGithub_NewRepository(t *testing.T) {
	t.Parallel()

	mock := pkg.NewMockProvider()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		got, err := pkg.GitHub(ctx)
		if err != nil {
			t.Error(err)
		}

		t.Run("リポジトリ名にはプロジェクト名が使われる", func(t *testing.T) {
			t.Parallel()

			pulumi.All(got.URN(), got.RepositoryName()).ApplyT(func(all []any) error {
				urn := getOutput[pulumi.URN](t, all, 0)
				name := getOutput[string](t, all, 1)

				if name != ctx.Project() {
					t.Errorf("URN=%sのリポジトリ名はプロジェクト名であるはずが %s を得た", urn, name)
				}

				return nil
			})
		})

		return nil
	},
		pulumi.WithMocks("project", "dev", mock),
		func(ri *pulumi.RunInfo) {
			ri.Config = mock.Config
		},
	)
	if err != nil {
		t.Fatalf("pulumi.RunErr failed: %v", err)
	}
}
