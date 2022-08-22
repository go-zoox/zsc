package zsc

import "fmt"

type Page struct {
	Page     int64 `query:"page,default=1"`
	PageSize int64 `query:"pageSize,default=10"`
}

type WhereOne struct {
	Key     string
	Value   interface{}
	IsFuzzy bool
	isNot   bool
}

type OrderByOne struct {
	Key    string
	IsDESC bool
}

type Where []WhereOne

type SetWhereOptions struct {
	IsFuzzy bool
	IsNot   bool
}

func (w *Where) Set(key string, value interface{}, opts ...*SetWhereOptions) {
	var isFuzzy bool
	var isNot bool
	if len(opts) > 0 {
		isFuzzy = opts[0].IsFuzzy
		isNot = opts[0].IsNot
	}

	*w = append(*w, WhereOne{
		Key:     key,
		Value:   value,
		IsFuzzy: isFuzzy,
		isNot:   isNot,
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
