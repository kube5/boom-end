package basic

import (
	"context"

	"go.uber.org/fx"
)

type Component interface {
	Start(sd fx.Shutdowner) error
	Stop(sd fx.Shutdowner) error
}

func RegisterHook(lc fx.Lifecycle, sd fx.Shutdowner, c Component) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return c.Start(sd)
		},
		OnStop: func(ctx context.Context) error {
			return c.Stop(sd)
		},
	})
}
