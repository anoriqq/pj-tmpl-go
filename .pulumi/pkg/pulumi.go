package pkg

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/anoriqq/pj-tmpl-go/.pulumi/pulumiutil"
)

// PulumiResource Pulumiのリソースを管理する構造体。
type PulumiResource struct{}

// NewStack 新しいPulumiスタックを作成する。
func (p *PulumiResource) NewStack(ctx *pulumi.Context, name string) (*pulumiservice.Stack, error) {
	stack, err := p.newStack(ctx, name)
	if err != nil {
		return nil, err
	}

	err = ctx.Log.Info(
		"new: "+name,
		&pulumi.LogArgs{
			Resource:  stack,
			StreamID:  0,
			Ephemeral: false,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return stack, nil
}

func (*PulumiResource) newStack(ctx *pulumi.Context, name string) (*pulumiservice.Stack, error) {
	args := &pulumiservice.StackArgs{
		OrganizationName: pulumi.String(ctx.Organization()),
		ProjectName:      pulumi.String(ctx.Project()),
		StackName:        pulumi.String(name),
		ForceDestroy:     nil,
	}
	opts := []pulumi.ResourceOption{}

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

	return result, nil
}

// Pulumi pulumiリソースを管理する構造体を返す。
func Pulumi() *PulumiResource {
	return &PulumiResource{}
}
