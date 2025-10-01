package pkg_test

import (
	"testing"

	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pkg"
)

func TestPulumi_NewStack(t *testing.T) {
	t.Parallel()

	t.Run("通常のスタック作成", func(t *testing.T) {
		t.Parallel()
		testStackCreation(t, "test-stack", "test-stack")
	})

	t.Run("デフォルトスタック作成", func(t *testing.T) {
		t.Parallel()
		testStackCreation(t, "dev", "dev")
	})
}

// testStackCreation は指定されたスタック名でスタックを作成し、基本的な検証を行う。
func testStackCreation(t *testing.T, stackName, expectedName string) {
	t.Helper()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		got, err := pkg.Pulumi().NewStack(ctx, stackName)
		if err != nil {
			t.Error(err)
		}

		// スタック名の検証
		validateStackName(t, got, expectedName)

		// 組織名とプロジェクト名の検証
		validateOrgAndProject(t, got, ctx)

		return nil
	},
		pulumi.WithMocks("test-org", stackName, pkg.NewMockProvider()),
		func(ri *pulumi.RunInfo) {
			ri.Config = map[string]string{
				"test-org:defaultStack":  "dev",
				"test-org:defaultBranch": "main",
			}
		},
	)
	if err != nil {
		t.Fatalf("pulumi.RunErr failed: %v", err)
	}
}

// validateStackName はスタック名が正しく設定されているかを検証する。
func validateStackName(t *testing.T, stack *pulumiservice.Stack, expectedName string) {
	t.Helper()

	pulumi.All(stack.URN(), stack.StackName).ApplyT(func(all []any) error {
		urn, ok := all[0].(pulumi.URN)
		if !ok {
			t.Fatal("URNの型変換に失敗")
		}

		name, ok := all[1].(string)
		if !ok {
			t.Fatal("StackNameの型変換に失敗")
		}

		if name != expectedName {
			t.Errorf("URN=%sのスタック名は %s であるはずが %s を得た", urn, expectedName, name)
		}

		return nil
	})
}

// validateOrgAndProject は組織名とプロジェクト名が正しく設定されているかを検証する。
func validateOrgAndProject(t *testing.T, stack *pulumiservice.Stack, ctx *pulumi.Context) {
	t.Helper()

	pulumi.All(stack.OrganizationName, stack.ProjectName).ApplyT(func(all []any) error {
		orgName, ok := all[0].(string)
		if !ok {
			t.Fatal("OrganizationNameの型変換に失敗")
		}

		projectName, ok := all[1].(string)
		if !ok {
			t.Fatal("ProjectNameの型変換に失敗")
		}

		if orgName != ctx.Organization() {
			t.Errorf("組織名は %s であるはずが %s を得た", ctx.Organization(), orgName)
		}

		if projectName != ctx.Project() {
			t.Errorf("プロジェクト名は %s であるはずが %s を得た", ctx.Project(), projectName)
		}

		return nil
	})
}
