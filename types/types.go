package types

import "server/consts"

// PaginationQuery 分页查询
type Pagination struct {
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}

// OrderQuery 排序
type Order struct {
	OrderKey  string           `json:"order_key"`
	OrderType consts.OrderType `json:"order_type"`
}
