package types

type Order struct {
	ID         int64       `json:"id" db:"id"`
	BusinessID int64       `json:"business_id" db:"business_id"`
	TableID    int64       `json:"table_id" db:"table_id"`
	TableNo    int64       `json:"table_no" db:"table_no"`
	Status     OrderStatus `json:"status" db:"status"`
	Total      float64     `json:"total" db:"total"`
	CreatedAt  string      `json:"created_at" db:"created_at"`
	UpdatedAt  string      `json:"updated_at" db:"updated_at"`
}

type OrderItem struct {
	ID        int64   `json:"id" db:"id"`
	ItemID    int64   `json:"item_id" db:"item_id"`
	OrderID   int64   `json:"order_id" db:"order_id"`
	Quantity  int64   `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
	CreatedAt string  `json:"created_at" db:"created_at"`
}

type OrderDetail struct {
	*Order
	OrderItems []*OrderItem `json:"order_items"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusPaid      OrderStatus = "paid"
)

func (status OrderStatus) IsValid() bool {
	switch status {
	case OrderStatusPending,
		OrderStatusPreparing,
		OrderStatusReady,
		OrderStatusPaid:
		return true
	default:
		return false
	}
}

type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}

type GetOrdersSuccessResponse struct {
	ResponseBase
	Data *GetOrdersResponse `json:"data"`
}
