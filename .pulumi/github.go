package main

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GitHubResource struct{}

func (g *GitHubResource) NewRepository(ctx *pulumi.Context) (*github.Repository, error) {
	owner := "anoriqq"
	repo := "pj-tmpl-go"
	branch := "main"

	repository, err := g.newRepository(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	ctx.Log.Info(
		fmt.Sprintf("new: %s/%s", owner, repo),
		&pulumi.LogArgs{Resource: repository},
	)

	branchDefault, err := g.newBranchDefault(ctx, owner, repo, branch)
	if err != nil {
		return nil, err
	}
	ctx.Log.Info(
		fmt.Sprintf("new: %s/%s %s", owner, repo, branch),
		&pulumi.LogArgs{Resource: branchDefault},
	)

	return repository, nil
}

func (*GitHubResource) newRepository(
	ctx *pulumi.Context,
	owner, repo string,
) (*github.Repository, error) {
	args := &github.RepositoryArgs{
		// General
		Name:       pulumi.String(repo),
		IsTemplate: pulumi.Bool(true),
		//   Features
		HasWiki:        pulumi.Bool(false),
		HasIssues:      pulumi.Bool(true),
		HasDiscussions: pulumi.Bool(false),
		HasProjects:    pulumi.Bool(false),
		HasDownloads:   pulumi.Bool(true),
		//   Pull Requests
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(false),
		SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
		SquashMergeCommitMessage: pulumi.String("PR_BODY"),
		AllowUpdateBranch:        pulumi.Bool(true),
		AllowAutoMerge:           pulumi.Bool(true),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		//   Danger Zone
		Visibility: pulumi.String("public"),

		// Security
		SecurityAndAnalysis: &github.RepositorySecurityAndAnalysisArgs{
			SecretScanning: &github.RepositorySecurityAndAnalysisSecretScanningArgs{
				Status: pulumi.String("enabled"),
			},
			SecretScanningPushProtection: &github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
				Status: pulumi.String("enabled"),
			},
		},
	}
	opts := []pulumi.ResourceOption{
		pulumi.Import(pulumi.ID(repo)),
		pulumi.RetainOnDelete(true),
	}
	enforceDevOnlyChanges(ctx, &opts)
	result, err := github.NewRepository(ctx, owner, args, opts...)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}

func (*GitHubResource) newBranchDefault(
	ctx *pulumi.Context,
	owner, repo string,
	branch string,
) (*github.BranchDefault, error) {
	args := &github.BranchDefaultArgs{
		Repository: pulumi.String(repo),
		Branch:     pulumi.String(branch),
		Rename:     pulumi.Bool(false),
	}
	opts := []pulumi.ResourceOption{}
	enforceDevOnlyChanges(ctx, &opts)
	result, err := github.NewBranchDefault(ctx, owner, args, opts...)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}

var GitHub GitHubResource
