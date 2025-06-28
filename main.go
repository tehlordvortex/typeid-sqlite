package main

import (
	"go.jetify.com/typeid"
	"go.riyazali.net/sqlite"
)

type Generate struct{}

func (g *Generate) Args() int           { return 1 }
func (g *Generate) Deterministic() bool { return false }
func (g *Generate) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	id, err := typeid.WithPrefix(values[0].Text())
	if err != nil {
		ctx.ResultError(err)
		return
	}

	ctx.ResultText(id.String())
}

type Check struct{}

func (c *Check) Args() int           { return 2 }
func (c *Check) Deterministic() bool { return true }
func (c *Check) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	id, err := typeid.FromString(values[1].Text())
	if err != nil {
		ctx.ResultInt(0)
		return
	}

	if id.Prefix() != values[0].Text() {
		ctx.ResultInt(0)
		return
	}

	ctx.ResultInt(1)
}

func init() {
	sqlite.Register(func(ea *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := ea.CreateFunction("typeid_generate_text", &Generate{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := ea.CreateFunction("typeid_check_text", &Check{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}

func main() {}
