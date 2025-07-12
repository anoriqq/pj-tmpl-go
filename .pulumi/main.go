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

		if err := NewRandomID(ctx, "random_10"); err != nil {
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

func NewRandomID(ctx *pulumi.Context, name string) error {
	randomID, err := newRandomID(ctx, name)
	if err != nil {
		return err
	}
	ctx.Log.Info(
		"new: ",
		&pulumi.LogArgs{Resource: randomID},
	)

	ctx.Export(name+"_randomIdHex", randomID.Hex)
	ctx.Export(name+"_randomIdB64Url", randomID.B64Url)

	return nil
}

func newRandomID(ctx *pulumi.Context, name string) (*random.RandomId, error) {
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
