package zsc

import (
	"fmt"

	"github.com/go-zoox/logger"
)

var creates = make(map[string]bool)

func Register(namespace string, model Model, service Service, controller Controller) {
	if _, ok := creates[namespace]; ok {
		panic(fmt.Sprintf("[cms] app(%s) already registered", namespace))
	}
	creates[namespace] = true

	logger.Info("[cms][app] register: %s", namespace)

	if model != nil {
		RegisterModel(model.ModelName(), model)
	}

	if service != nil {
		RegisterService(service.Name(), service)
	}

	if controller != nil {
		RegisterController(controller.Name(), controller)
	}
}
