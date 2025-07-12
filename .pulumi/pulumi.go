package main

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PulumiResource struct{}

func (p *PulumiResource) NewStack(ctx *pulumi.Context, name string) (*pulumiservice.Stack, error) {
	stack, err := p.newStack(ctx, name)
	if err != nil {
		return nil, err
	}
	ctx.Log.Info(
		fmt.Sprintf("new: %s", name),
		&pulumi.LogArgs{Resource: stack},
	)

	return stack, nil
}

func (*PulumiResource) newStack(ctx *pulumi.Context, name string) (*pulumiservice.Stack, error) {
	args := &pulumiservice.StackArgs{
		OrganizationName: pulumi.String(ctx.Organization()),
		ProjectName:      pulumi.String(ctx.Project()),
		StackName:        pulumi.String(name),
	}
	opts := []pulumi.ResourceOption{}
	if name == getDefaultStack(ctx) {
		stackID := fmt.Sprintf("%s/%s/%s", ctx.Organization(), ctx.Project(), name)
		opts = append(
			opts,
			pulumi.Import(pulumi.ID(stackID)),
			pulumi.RetainOnDelete(true),
		)
	}
	result, err := pulumiservice.NewStack(ctx, name, args, opts...)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}

var Pulumi = &PulumiResource{}
