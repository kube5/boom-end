package base

import (
	"sync"

	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

var lock = new(sync.Mutex)
var constructors []interface{}

func RegisterComponents(componentConstructors ...interface{}) {
	lock.Lock()
	constructors = append(constructors, componentConstructors...)
	lock.Unlock()
}

func BuildComponents(cfg repo.CommonComponents, fxInvokeFunc interface{}) (*fx.App, error) {
	app := fx.New(
		fx.Logger(cfg.Logger),
		fx.Supply(cfg),
		fx.Provide(
			constructors...,
		),
		fx.Invoke(fxInvokeFunc),
	)

	if app.Err() != nil {
		return nil, errors.Wrap(app.Err(), "app setup error")
	}
	return app, nil
}
