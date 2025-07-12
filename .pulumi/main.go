package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		if isDefaultStack(ctx) {
			if err := defaultStackOnly(ctx); err != nil {
				return err
			}
		}

		return nil
	})
}

func defaultStackOnly(ctx *pulumi.Context) error {
	if _, err := Pulumi.NewStack(ctx, getDefaultStack(ctx)); err != nil {
		return err
	}

	if _, err := Pulumi.NewStack(ctx, "stg"); err != nil {
		return err
	}

	if _, err := Pulumi.NewStack(ctx, "prd"); err != nil {
		return err
	}

	if _, err := GitHub.NewRepository(ctx); err != nil {
		return err
	}

	return nil
}
