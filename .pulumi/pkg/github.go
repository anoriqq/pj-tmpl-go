package pkg

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pulumiutil"
)

// GithubResource GitHubのリソースを管理する構造体。
type GithubResource struct {
	pulumi.ResourceState

	repository    *github.Repository
	branchDefault *github.BranchDefault
}

func (r *GithubResource) RepositoryName() pulumi.StringOutput {
	return r.repository.Name
}

func (r *GithubResource) newBranchDefault(
	ctx *pulumi.Context,
	owner, repo string,
) (*github.BranchDefault, error) {
	branch := pulumiutil.GetDefaultBranch(ctx)

	args := &github.BranchDefaultArgs{
		Repository: pulumi.String(repo),
		Branch:     pulumi.String(branch),
		Rename:     pulumi.Bool(false),
	}
	opts := []pulumi.ResourceOption{
		pulumi.Parent(r),
	}

	result, err := github.NewBranchDefault(ctx, fmt.Sprintf("%s-%s", repo, branch), args, opts...)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	err = ctx.Log.Info(
		fmt.Sprintf("new: %s/%s", owner, repo),
		&pulumi.LogArgs{
			Resource:  result,
			StreamID:  0,
			Ephemeral: false,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}

func (r *GithubResource) newRepository(
	ctx *pulumi.Context,
	owner, repo string,
) (*github.Repository, error) {
	args := &github.RepositoryArgs{
		// General
		Name:              pulumi.String(repo),
		Description:       pulumi.StringPtr(""),
		HomepageUrl:       pulumi.StringPtr(""),
		Topics:            pulumi.ToStringArray([]string{}),
		IsTemplate:        pulumi.Bool(true),
		AutoInit:          pulumi.Bool(false),
		GitignoreTemplate: pulumi.StringPtr(""),
		LicenseTemplate:   pulumi.StringPtr(""),
		Template:          nil,
		// Features
		HasWiki:        pulumi.Bool(false),
		HasIssues:      pulumi.Bool(true),
		HasDiscussions: pulumi.Bool(false),
		HasProjects:    pulumi.Bool(false),
		HasDownloads:   pulumi.Bool(true),
		// Pull Requests
		AllowSquashMerge:         pulumi.Bool(true),
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(false),
		SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
		SquashMergeCommitMessage: pulumi.String("PR_BODY"),
		MergeCommitTitle:         pulumi.String("PR_TITLE"),
		MergeCommitMessage:       pulumi.String("PR_BODY"),
		AllowUpdateBranch:        pulumi.Bool(true),
		AllowAutoMerge:           pulumi.Bool(true),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		// Danger Zone
		Visibility:       pulumi.String("public"),
		ArchiveOnDestroy: pulumi.Bool(true),
		Archived:         pulumi.Bool(false),
		// Security
		SecurityAndAnalysis: &github.RepositorySecurityAndAnalysisArgs{
			SecretScanning: &github.RepositorySecurityAndAnalysisSecretScanningArgs{
				Status: pulumi.String("enabled"),
			},
			SecretScanningPushProtection: &github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
				Status: pulumi.String("enabled"),
			},
			AdvancedSecurity: nil,
		},
		IgnoreVulnerabilityAlertsDuringRead: pulumi.Bool(false),
		VulnerabilityAlerts:                 pulumi.Bool(true),
		WebCommitSignoffRequired:            pulumi.Bool(false),
		// Pages
		Pages: nil,
		// Deprecated
		DefaultBranch: nil,
		Private:       nil,
	}
	opts := []pulumi.ResourceOption{
		pulumi.Import(pulumi.ID(repo)),
		pulumi.RetainOnDelete(true),
		pulumi.Parent(r),
	}

	result, err := github.NewRepository(ctx, repo, args, opts...)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	err = ctx.Log.Info(
		fmt.Sprintf("new: %s/%s", owner, repo),
		&pulumi.LogArgs{
			Resource:  result,
			StreamID:  0,
			Ephemeral: false,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}

// GitHub githubリソースを管理する。
func GitHub(ctx *pulumi.Context) (*GithubResource, error) {
	comp := &GithubResource{}
	t := fmt.Sprintf("%s:github:Suite", ctx.Organization())
	err := ctx.RegisterComponentResource(t, "github", comp)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	owner := pulumiutil.GetDefaultRepositoryOwner(ctx)
	repo := ctx.Project()

	// リポジトリ
	repository, err := comp.newRepository(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	comp.repository = repository

	// デフォルトブランチ
	branchDefault, err := comp.newBranchDefault(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	comp.branchDefault = branchDefault

	return comp, nil
}
