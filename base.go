package zsc

import (
	"fmt"
	"strings"
)

type Page struct {
	Page     int64 `query:"page,default=1"`
	PageSize int64 `query:"pageSize,default=10"`
}

type WhereOne struct {
	Key     string
	Value   interface{}
	IsFuzzy bool
	isNot   bool
	isIn    bool
}

type OrderByOne struct {
	Key    string
	IsDESC bool
}

type Where []WhereOne

type SetWhereOptions struct {
	IsFuzzy bool
	IsNot   bool
	IsIn    bool
}

func (w *Where) Set(key string, value interface{}, opts ...*SetWhereOptions) {
	var isFuzzy bool
	var isNot bool
	var isIn bool
	if len(opts) > 0 && opts[0] != nil {
		isFuzzy = opts[0].IsFuzzy
		isNot = opts[0].IsNot
		isIn = opts[0].IsIn
	}

	*w = append(*w, WhereOne{
		Key:     key,
		Value:   value,
		IsFuzzy: isFuzzy,
		isNot:   isNot,
		isIn:    isIn,
	})
}

func (w *Where) Get(key string) (interface{}, bool) {
	for _, v := range *w {
		if v.Key == key {
			return v.Value, true
		}
	}

	return "", false
}

func (w *Where) Length() int {
	return len(*w)
}

func (w *Where) Build() (string, []interface{}) {
	whereClauses := []string{}
	whereValues := []interface{}{}
	for _, w := range *w {
		if w.IsFuzzy {
			whereClauses = append(whereClauses, fmt.Sprintf("%s ILike ?", w.Key))
			whereValues = append(whereValues, fmt.Sprintf("%%%s%%", w.Value))
		} else if w.isNot {
			whereClauses = append(whereClauses, fmt.Sprintf("%s != ?", w.Key))
			whereValues = append(whereValues, w.Value)
		} else if w.isIn {
			whereClauses = append(whereClauses, fmt.Sprintf("%s in (?)", w.Key))
			whereValues = append(whereValues, w.Value)
		} else {
			whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", w.Key))
			whereValues = append(whereValues, w.Value)
		}
	}
	whereClause := strings.Join(whereClauses, " AND ")

	return whereClause, whereValues
}

func (w *Where) Debug() {
	for _, where := range *w {
		var fuzzy string
		if where.IsFuzzy {
			fuzzy = "Fuzzy"
		} else {
			fuzzy = "Extract"
		}

		fmt.Printf("[where] %s %s %s\n", where.Key, where.Value, fuzzy)
	}
}

type OrderBy []OrderByOne

func (w *OrderBy) Set(key string, IsDESC bool) {
	*w = append(*w, OrderByOne{
		Key:    key,
		IsDESC: IsDESC,
	})
}

func (w *OrderBy) Get(key string) (bool, bool) {
	for _, v := range *w {
		if v.Key == key {
			return v.IsDESC, true
		}
	}

	return false, false
}

func (w *OrderBy) Debug() {
	for _, v := range *w {
		var desc string
		if v.IsDESC {
			desc = "DESC"
		} else {
			desc = "ASC"
		}

		fmt.Printf("[order_by] %s %s\n", v.Key, desc)
	}
}

func (w *OrderBy) Length() int {
	return len(*w)
}

func (w *OrderBy) Build() string {
	orders := []string{}
	for _, order := range *w {
		orderMod := "ASC"
		if order.IsDESC {
			orderMod = "DESC"
		}

		orders = append(orders, fmt.Sprintf("%s %s", order.Key, orderMod))
	}

	return strings.Join(orders, ",")
}
