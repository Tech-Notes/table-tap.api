package types

type TableStatus string

const (
	TableStatusAvailable  TableStatus = "available"
	TableStatusReserved   TableStatus = "reserved"
	TableStatusOccupied   TableStatus = "occupied"     // When the table is currently in use
	TableStatusCleaning   TableStatus = "cleaning"     // When the table is being cleaned
	TableStatusOutOfOrder TableStatus = "out_of_order" // When the table is not usable
)

type Table struct {
	ID         int64       `json:"id" db:"id"`
	BusinessID int64       `json:"business_id" db:"business_id"`
	Token      string      `json:"token" db:"token"`
	Status     TableStatus `json:"status" db:"status"`
	QrCodeURL  string      `json:"qr_code_url" db:"qr_code_url"`
}

type TableDetailResponse struct {
	Table *Table `json:"table"`
}
type TableDetailSuccessResponse struct {
	ResponseBase
	Data *TableDetailResponse `json:"data"`
}
