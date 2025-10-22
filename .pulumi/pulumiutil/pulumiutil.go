/*
Package pulumiutil provides utility functions for Pulumi projects.
*/
package pulumiutil

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// IsDefaultStack checks if the current stack is the default stack.
func IsDefaultStack(ctx *pulumi.Context) bool {
	return ctx.Stack() == GetDefaultStack(ctx)
}

// GetDefaultStack returns the default stack name from the configuration.
func GetDefaultStack(ctx *pulumi.Context) string {
	cfg := config.New(ctx, "")
	result := cfg.Require("defaultStack")

	return result
}

// GetDefaultRepositoryOwner returns the default repository owner from the configuration.
func GetDefaultRepositoryOwner(ctx *pulumi.Context) string {
	cfg := config.New(ctx, "")
	result := cfg.Require("defaultRepositoryOwner")

	return result
}

// GetDefaultBranch returns the default branch name from the configuration.
func GetDefaultBranch(ctx *pulumi.Context) string {
	cfg := config.New(ctx, "")
	result := cfg.Require("defaultBranch")

	return result
}
