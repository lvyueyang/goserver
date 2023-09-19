package utils

import (
	"server/types"
)

type FindListOption struct {
	types.Pagination
	types.Order
}

func PageTrans(page types.Pagination) (offset int, limit int) {
	current := page.Current
	pageSize := page.PageSize
	if current == 0 {
		current = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	return (current - 1) * pageSize, pageSize
}

type ListResult[T any] struct {
	Total int64 `json:"total"`
	List  T     `json:"list"`
}
