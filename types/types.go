package types

import "server/consts"

// PaginationQuery 分页查询
type PaginationQuery struct {
	Current  uint `json:"current"`
	PageSize uint `json:"page_size"`
}

// OrderQuery 排序
type OrderQuery struct {
	OrderKey  string           `json:"order_key"`
	OrderType consts.OrderType `json:"order_type"`
}
