package zsc

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
)

type Controller interface {
	Name() string
	//
	// Model() container.Container
	Service() container.Container
	//
	Params(ctx *zoox.Context) *Params
	// //
	// Success(ctx *zoox.Context, data interface{})
	// FailWithError(ctx *zoox.Context, err zoox.HTTPError)
	User(ctx *zoox.Context) interface{}
}

type ControllerImpl struct {
}

func RegisterController(name string, m Controller) {
	if controller.Has(name) {
		panic("controller already exists: " + name)
	}

	logger.Info("[cms][controller] register: %s", name)
	controller.Register(name, m)
}

func GetController[T any](id string) T {
	if !controller.Has(id) {
		panic("controller not registered: " + id)
	}

	s, ok := controller.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("controller not valid type(%v): %s", new(T), id))
	}

	return s
}

// func (c *ControllerImpl) Context() *zoox.Context {
// 	return c.ctx
// }

func (c *ControllerImpl) Service() container.Container {
	return service
}

// func (c *ControllerImpl) Model() container.Container {
// 	return model
// }

// func (c *ControllerImpl) ServiceS() S {
// 	v, ok := c.Service().MustGet(c.Name()).(S)
// 	if !ok {
// 		panic(fmt.Sprintf("service(%s) not found", c.Name()))
// 	}

// 	return v
// }

func (c *ControllerImpl) Params(ctx *zoox.Context) *Params {
	return NewParams(ctx)
}

func (c *ControllerImpl) Success(ctx *zoox.Context, data interface{}) {
	ctx.Success(data)
}

func (c *ControllerImpl) FailWithError(ctx *zoox.Context, err zoox.HTTPError) {
	ctx.FailWithError(err)
}

func (c *ControllerImpl) User(ctx *zoox.Context) interface{} {
	return ctx.User().Get()
}
