package params

import (
	"context"
)

type Params struct {
}

func (params *Params) GetInt(name string) int {
	return 0
}

func (params *Params) GetString(name string) string {
	return ""
}

func (params *Params) Get(name string) string {
	return params.GetString(name)
}

func ParamsFromContext(ctx *context.Context) Params {
	return Params{}
}
