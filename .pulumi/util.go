package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// isDefaultStack checks if the current stack is the default stack.
func isDefaultStack(ctx *pulumi.Context) bool {
	return ctx.Stack() == getDefaultStack(ctx)
}

// getDefaultStack returns the default stack name from the configuration.
func getDefaultStack(ctx *pulumi.Context) string {
	cfg := config.New(ctx, "")
	result := cfg.Require("defaultStack")
	return result
}

// getDefaultBranch returns the default branch name from the configuration.
func getDefaultBranch(ctx *pulumi.Context) string {
	cfg := config.New(ctx, "")
	result := cfg.Require("defaultBranch")
	return result
}
