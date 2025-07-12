package main

import (
	"github.com/go-errors/errors"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		if isDefaultStack(ctx) {
			if err := defaultStackOnly(ctx); err != nil {
				return err
			}
		}

		if err := NewRandomID(ctx); err != nil {
			return err
		}

		return nil
	})
}

func defaultStackOnly(ctx *pulumi.Context) error {
	if _, err := Pulumi().NewStack(ctx, getDefaultStack(ctx)); err != nil {
		return err
	}

	if _, err := Pulumi().NewStack(ctx, "stg"); err != nil {
		return err
	}

	if _, err := Pulumi().NewStack(ctx, "prd"); err != nil {
		return err
	}

	if _, err := GitHub().NewRepository(ctx); err != nil {
		return err
	}

	return nil
}

func NewRandomID(ctx *pulumi.Context) error {
	randomID, err := newRandomID(ctx)
	if err != nil {
		return err
	}
	ctx.Log.Info(
		"new: ",
		&pulumi.LogArgs{Resource: randomID},
	)

	ctx.Export("randomIdHex", randomID.Hex)
	ctx.Export("randomId", randomID.ID())

	return nil
}

func newRandomID(ctx *pulumi.Context) (*random.RandomId, error) {
	name := "random"

	args := &random.RandomIdArgs{
		Keepers:    pulumi.StringMap{},
		ByteLength: pulumi.Int(8),
	}
	result, err := random.NewRandomId(ctx, name, args)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return result, nil
}
