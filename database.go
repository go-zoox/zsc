package zsc

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func DBList[T any](page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error) {
	offset := int((page - 1) * pageSize)
	limit := int(pageSize)

	// whereClauses := []string{}
	// whereValues := []interface{}{}
	// for _, w := range *where {
	// 	if w.IsFuzzy {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s ILike ?", w.Key))
	// 		whereValues = append(whereValues, fmt.Sprintf("%%%s%%", w.Value))
	// 	} else if w.isNot {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	} else if w.isIn {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	} else {
	// 		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", w.Key))
	// 		whereValues = append(whereValues, w.Value)
	// 	}
	// }
	// whereClause := strings.Join(whereClauses, " AND ")

	whereClause, whereValues := where.Build()

	countTx := GetDB().Model(new(T))
	dataTx := GetDB()

	if orderBy != nil {
		for _, order := range *orderBy {
			// fmt.Println("order by:", order.Key, order.IsDESC)
			orderMod := "ASC"
			if order.IsDESC {
				orderMod = "DESC"
			}

			orderStr := fmt.Sprintf("%s %s", order.Key, orderMod)
			countTx = countTx.Order(orderStr)
			dataTx = dataTx.Order(orderStr)
		}
	}
	if whereClause != "" {
		countTx = countTx.Where(whereClause, whereValues...)
		dataTx = dataTx.Where(whereClause, whereValues...)
	}

	err = countTx.
		Count(&total).
		Error
	if err != nil {
		return
	}

	err = dataTx.
		Offset(offset).
		Limit(limit).
		Find(&data).
		Error

	return
}

func DBCreate[T any](one *T) (*T, error) {
	err := GetDB().
		Create(one).Error

	return one, err
}

func DBRetrieve[T any](id uint) (*T, error) {
	var f T
	if err := GetDB().First(&f, id).Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func DBUpdate[T any](id uint, uc func(*T)) (err error) {
	var f T
	err = GetDB().First(&f, id).Error
	if err != nil {
		return
	}

	uc(&f)

	err = GetDB().Save(&f).Error
	return
}

func DBDelete[T any](id uint) (err error) {
	var f T
	err = GetDB().First(&f, id).Error
	if err != nil {
		return
	}

	err = GetDB().Delete(&f).Error
	return
}

func DBSave[T any](one *T) error {
	return GetDB().Save(one).Error
}

func DBHas[T any](id uint) bool {
	var f T
	if err := GetDB().First(&f, id).Error; err != nil {
		return false
	}

	return true
}

func DBGetMany[T any](ids []uint) (data []*T, err error) {
	err = GetDB().
		Where("id IN (?)", ids).
		Find(&data).Error
	return
}

func DBFind[T any](where map[any]any) (data []*T, err error) {
	err = GetDB().Find(&data, where).Error
	return
}

func DBFindOne[T any](where map[any]any) (*T, error) {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func DBExists[T any](where map[any]any) (bool, error) {
	_, err := DBFindOne[T](where)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func DBFindByID[T any](id uint) (*T, error) {
	var f T
	if err := GetDB().First(&f, id).Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func DBFindOneOrCreate[T any](where map[any]any, callback func(*T)) (*T, error) {
	f, err := DBFindOne[T](where)
	if err != nil {
		var tmp T
		callback(&tmp)

		if f, err = DBCreate(&tmp); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func DBFindOneByIDOrCreate[T any](id uint, callback func(*T)) (*T, error) {
	f, err := DBFindByID[T](id)
	if err != nil {
		var tmp T
		callback(&tmp)

		if f, err = DBCreate(&tmp); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func DBFindOneAndUpdate[T any](where map[any]any, callback func(*T)) (*T, error) {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return nil, err
	}

	callback(&f)

	return &f, nil
}

func DBFindOneByIDAndUpdate[T any](id uint, callback func(*T)) (*T, error) {
	f, err := DBFindByID[T](id)
	if err != nil {
		return nil, err
	}

	callback(f)

	return f, nil
}

func DBFindOneAndDelete[T any](where map[any]any) (*T, error) {
	f, err := DBFindOne[T](where)
	if err != nil {
		return nil, err
	}

	err = GetDB().Delete(f).Error
	return f, err
}

func DBFindOneByIDAndDelete[T any](id uint) (*T, error) {
	f, err := DBFindByID[T](id)
	if err != nil {
		return nil, err
	}

	err = GetDB().Delete(f).Error
	return f, err
}

func DBGetOrCreate[T any](where map[any]any, callback func(*T)) (*T, error) {
	return DBFindOneOrCreate(where, callback)
}

func DBListAll[T any](where *Where, orderBy *OrderBy) (data []*T, err error) {
	whereClause, whereValues := where.Build()

	dataTx := GetDB()

	if orderBy != nil {
		for _, order := range *orderBy {
			// fmt.Println("order by:", order.Key, order.IsDESC)
			orderMod := "ASC"
			if order.IsDESC {
				orderMod = "DESC"
			}

			orderStr := fmt.Sprintf("%s %s", order.Key, orderMod)
			dataTx = dataTx.Order(orderStr)
		}
	}
	if whereClause != "" {
		dataTx = dataTx.Where(whereClause, whereValues...)
	}

	err = dataTx.
		Find(&data).
		Error

	return
}
