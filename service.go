package zsc

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

type Service interface {
	Name() string
	Model() container.Container
	Service() container.Container
}

func RegisterService(name string, m Service) {
	if service.Has(name) {
		panic("service already exists: " + name)
	}

	logger.Info("[cms][service] register: %s", name)
	service.Register(name, m)
}

func GetService[T any](id string) T {
	if !service.Has(id) {
		panic("service not registered: " + id)
	}

	s, ok := service.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("service not valid type(%v): %s", new(T), id))
	}

	return s
}

type ServiceImpl struct {
}

func (s *ServiceImpl) Name() string {
	panic("service.Name() not implemented")
}

func (s *ServiceImpl) Model() container.Container {
	return model
}

func (s *ServiceImpl) Service() container.Container {
	return service
}

// type ServiceInf[T any] interface {
// 	Name() string
// 	//
// 	Model() ModelInf[T]
// 	SetModel(modelName string)
// 	//
// 	Service() container.Container
// 	// Model() container.Container
// 	//
// 	List(page, pageSize int, where *Where, orderBy *OrderBy) ([]*T, int64, error)
// 	Create(one *T) (*T, error)
// 	Retrieve(id uint) (*T, error)
// 	Update(id uint, one *T) error
// 	Delete(id uint) error
// 	// Save(one *T) error
// }

// type Service[T any] struct {
// 	m ModelInf[T]
// 	//
// 	modelName string
// }

// func (s *Service[T]) Model() ModelInf[T] {
// 	if s.m == nil {
// 		var ok bool
// 		if s.m, ok = model.MustGet(s.modelName).(ModelInf[T]); !ok {
// 			panic(fmt.Sprintf("cannot get model(%s)", s.modelName))
// 		}
// 	}

// 	return s.m
// }

// func (s *Service[T]) Service() container.Container {
// 	return service
// }

// func (s *Service[T]) SetModel(modelName string) {
// 	if s.modelName != "" {
// 		panic(fmt.Sprintf("model already set %s, but got %s", s.modelName, modelName))
// 	}

// 	s.modelName = modelName
// }

// func (s *Service[T]) List(page, pageSize int, where *Where, orderBy *OrderBy) ([]*T, int64, error) {
// 	orderBy.Set("updated_at", true)

// 	data, total, err := s.Model().List(page, pageSize, where, orderBy)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return data, total, nil
// }

// func (s *Service[T]) Create(one *T) (*T, error) {
// 	return s.Model().Create(one)
// }

// func (s *Service[T]) Retrieve(id uint) (*T, error) {
// 	return s.Model().Get(id)
// }

// func (s *Service[T]) Update(id uint, update func(*T)) error {
// 	return s.Model().Update(id, update)
// }

// func (s *Service[T]) Delete(id uint) error {
// 	return s.Model().Delete(id)
// }
