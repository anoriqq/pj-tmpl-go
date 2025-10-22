package pkg

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pulumiutil"
)

// PulumiResource Pulumiのリソースを管理する構造体。
type PulumiResource struct {
	pulumi.ResourceState

	stackDefault *pulumiservice.Stack
	stackStg     *pulumiservice.Stack
	stackPrd     *pulumiservice.Stack
}

func (r *PulumiResource) StackDefault() *pulumiservice.Stack {
	return r.stackDefault
}

func (r *PulumiResource) StackStg() *pulumiservice.Stack {
	return r.stackStg
}

func (r *PulumiResource) StackPrd() *pulumiservice.Stack {
	return r.stackPrd
}

func (r *PulumiResource) newStack(ctx *pulumi.Context, name string) (*pulumiservice.Stack, error) {
	args := &pulumiservice.StackArgs{
		OrganizationName: pulumi.String(ctx.Organization()),
		ProjectName:      pulumi.String(ctx.Project()),
		StackName:        pulumi.String(name),
		ForceDestroy:     nil,
	}
	opts := []pulumi.ResourceOption{
		pulumi.Parent(r),
	}

	if name == pulumiutil.GetDefaultStack(ctx) {
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

	err = ctx.Log.Info(
		"new: "+name,
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

// Pulumi pulumiリソースを管理する構造体を返す。
func Pulumi(ctx *pulumi.Context) (*PulumiResource, error) {
	comp := &PulumiResource{}
	t := fmt.Sprintf("%s:pulumi:Suite", ctx.Organization())
	err := ctx.RegisterComponentResource(t, "pulumi", comp)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	// Stack (default)
	stackDefault, err := comp.newStack(ctx, pulumiutil.GetDefaultStack(ctx))
	if err != nil {
		return nil, err
	}

	comp.stackDefault = stackDefault

	// Stack (stg)
	stackStg, err := comp.newStack(ctx, "stg")
	if err != nil {
		return nil, err
	}

	comp.stackStg = stackStg

	// Stack (prd)
	stackPrd, err := comp.newStack(ctx, "prd")
	if err != nil {
		return nil, err
	}

	comp.stackPrd = stackPrd

	return comp, err
}
