package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pkg"
	"github.com/anoriqq/pj-tmpl-go/.pulumi/pulumiutil"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		if pulumiutil.IsDefaultStack(ctx) {
			err := defaultStackOnly(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func defaultStackOnly(ctx *pulumi.Context) error {
	_, err := pkg.Pulumi().NewStack(ctx, pulumiutil.GetDefaultStack(ctx))
	if err != nil {
		return err
	}

	_, err = pkg.Pulumi().NewStack(ctx, "stg")
	if err != nil {
		return err
	}

	_, err = pkg.Pulumi().NewStack(ctx, "prd")
	if err != nil {
		return err
	}

	_, err = pkg.GitHub().NewRepository(ctx)
	if err != nil {
		return err
	}

	return nil
}
