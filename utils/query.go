package utils

import (
	"server/types"
)

type FindListOption struct {
	types.Pagination
	types.Order
}

func PaginationDefault(page types.Pagination) types.Pagination {
	current := page.Current
	pageSize := page.PageSize
	if current == 0 {
		current = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	return types.Pagination{
		Current:  current,
		PageSize: pageSize,
	}
}
