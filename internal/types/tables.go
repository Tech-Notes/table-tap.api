package types

import (
	"gopkg.in/guregu/null.v4"
)

type TableStatus string

const (
	TableStatusAvailable  TableStatus = "available"
	TableStatusReserved   TableStatus = "reserved"
	TableStatusOccupied   TableStatus = "occupied"     // When the table is currently in use
	TableStatusCleaning   TableStatus = "cleaning"     // When the table is being cleaned
	TableStatusOutOfOrder TableStatus = "out_of_order" // When the table is not usable
)

type Table struct {
	ID          int64       `json:"id" db:"id"`
	BusinessID  int64       `json:"business_id" db:"business_id"`
	Token       string      `json:"token" db:"token"`
	Status      TableStatus `json:"status" db:"status"`
	QrCodeURL   string      `json:"qr_code_url" db:"qr_code_url"`
	TableNo     int64       `json:"table_no" db:"table_no"`
	Description null.String `json:"description" db:"description"`
	CreatedAt   ISO8601Time `json:"created_at" db:"created_at"`
}

type TableDetailResponse struct {
	Table  *Table   `json:"table"`
	Orders []*Order `json:"orders"`
}
type TableDetailSuccessResponse struct {
	ResponseBase
	Data *TableDetailResponse `json:"data"`
}

type CreateTableRequest struct {
	TableNo     int64       `json:"table_no"`
	Description null.String `json:"description"`
}
