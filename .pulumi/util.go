package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// enforceDevOnlyChanges ensures that changes are only applied in the "dev" stack.
func enforceDevOnlyChanges(ctx *pulumi.Context, args *[]pulumi.ResourceOption) {
	if ctx.Stack() != "dev" {
		*args = append(*args, pulumi.IgnoreChanges([]string{"*"}))
	}
}
