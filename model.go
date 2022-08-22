package zsc

import (
	"fmt"
	"time"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
	"gorm.io/gorm"
)

type Model interface {
	ModelName() string
	Model() container.Container
}

func RegisterModel(name string, m Model) {
	if model.Has(name) {
		panic("model already exists: " + name)
	}

	logger.Info("[cms][model] register: %s", name)
	model.Register(name, m)
}

func GetModel[T any](id string) T {
	if !model.Has(id) {
		panic("model not registered: " + id)
	}

	s, ok := model.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("model not valid type(%v): %s", new(T), id))
	}

	return s
}

type ModelImpl struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//
	Creator uint `json:"creator"`
}

func (m *ModelImpl) ModelName() string {
	panic("model.ModelName() not implemented")
}

func (m *ModelImpl) Model() container.Container {
	return model
}

// type ModelInf[T any] interface {
// 	ModelName() string
// 	//
// 	List(page, pageSize int, where *Where, orderBy *OrderBy) (data []*T, total int64, err error)
// 	Create(one *T) (*T, error)
// 	Get(id uint) (*T, error)
// 	Update(id uint, uc func(*T)) error
// 	Delete(id uint) error
// 	Save(one *T) error
// }

// type Model[T any] struct {
// 	ID uint `gorm:"primarykey" json:"id"`
// 	// CreatedAt time.Time      `json:"created_at"`
// 	// UpdatedAt time.Time      `json:"updated_at"`
// 	CreatedAt time.Time      `json:"createdAt"`
// 	UpdatedAt time.Time      `json:"updatedAt"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
// 	//
// 	Creator int `json:"creator"`
// }

// func New[T any]() *Model[T] {
// 	return &Model[T]{}
// }

// func (m *Model[T]) List(page, pageSize int, where *Where, orderBy *OrderBy) (data []*T, total int64, err error) {
// 	offset := (page - 1) * pageSize
// 	limit := pageSize

// 	whereClauses := []string{}
// 	whereValues := []interface{}{}
// 	for _, w := range *where {
// 		if w.IsFuzzy {
// 			whereClauses = append(whereClauses, fmt.Sprintf("%s Like ?", w.Key))
// 			whereValues = append(whereValues, fmt.Sprintf("%%%s%%", w.Value))
// 		} else {
// 			whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", w.Key))
// 			whereValues = append(whereValues, w.Value)
// 		}
// 	}
// 	whereClause := strings.Join(whereClauses, " AND ")

// 	countTx := GetDB().Model(new(T))
// 	dataTx := GetDB()

// 	if orderBy != nil {
// 		for _, order := range *orderBy {
// 			countTx = countTx.Order(order)
// 			dataTx = dataTx.Order(order)
// 		}
// 	}
// 	if whereClause != "" {
// 		countTx = countTx.Where(whereClause, whereValues...)
// 		dataTx = dataTx.Where(whereClause, whereValues...)
// 	}

// 	err = countTx.
// 		Count(&total).
// 		Error
// 	if err != nil {
// 		return
// 	}

// 	err = dataTx.
// 		Offset(offset).
// 		Limit(limit).
// 		Find(&data).
// 		Error

// 	return
// }

// func (m *Model[T]) Create(one *T) (*T, error) {
// 	err := GetDB().
// 		Create(one).Error

// 	return one, err
// }

// func (m *Model[T]) Update(id uint, uc func(*T)) (err error) {
// 	var f T
// 	err = GetDB().First(&f, id).Error
// 	if err != nil {
// 		return
// 	}

// 	uc(&f)

// 	err = GetDB().Save(&f).Error
// 	return
// }

// func (m *Model[T]) Delete(id uint) (err error) {
// 	var f T
// 	err = GetDB().First(&f, id).Error
// 	if err != nil {
// 		return
// 	}

// 	err = GetDB().Delete(&f).Error
// 	return
// }

// func (m *Model[T]) Get(id uint) (*T, error) {
// 	var f T
// 	if err := GetDB().First(&f, id).Error; err != nil {
// 		return nil, err
// 	}

// 	return &f, nil
// }

// func (m *Model[T]) Save(one *T) error {
// 	return GetDB().Save(one).Error
// }
