package types

// Pagination 分页查询
type Pagination struct {
	Current  int `json:"current" form:"current" binding:"omitempty,min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"omitempty,min=0"`
}

// Order 排序
type Order struct {
	OrderKey  string `json:"order_key" form:"order_key"`
	OrderType string `json:"order_type" form:"order_type"`
}
