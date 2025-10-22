package pkg_test

import (
	"testing"

	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pkg"
)

func TestPulumi_NewStack(t *testing.T) {
	mock := pkg.NewMockProvider()
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		got, err := pkg.Pulumi(ctx)
		if err != nil {
			t.Error(err)
		}

		validateStack(t, ctx, got.StackDefault(), "dev")
		validateStack(t, ctx, got.StackStg(), "stg")
		validateStack(t, ctx, got.StackPrd(), "prd")

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

// validateStack スタックが正しく設定されているかを検証する。
func validateStack(t *testing.T, ctx *pulumi.Context, stack *pulumiservice.Stack, wantStackName string) {
	t.Helper()

	pulumi.All(
		stack.URN(),
		stack.StackName,
		stack.OrganizationName,
		stack.ProjectName,
	).ApplyT(func(all []any) error {
		urn := getOutput[pulumi.URN](t, all, 0)
		name := getOutput[string](t, all, 1)
		orgName := getOutput[string](t, all, 2)
		projectName := getOutput[string](t, all, 3)

		// Assert
		t.Run(string(urn), func(t *testing.T) {
			if name != wantStackName {
				t.Errorf("スタック名は %s であるはずが %s を得た", wantStackName, name)
			}

			if orgName != ctx.Organization() {
				t.Errorf("組織名は %s であるはずが %s を得た", ctx.Organization(), orgName)
			}

			if projectName != ctx.Project() {
				t.Errorf("プロジェクト名は %s であるはずが %s を得た", ctx.Project(), projectName)
			}
		})

		return nil
	})
}
