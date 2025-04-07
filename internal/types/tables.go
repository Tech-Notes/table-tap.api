package internal

type Table struct {
	ID          int64  `json:"id" db:"id"`
	BusinessID  int64  `json:"business_id" db:"business_id"`
	Status 	string `json:"status" db:"status"`
	QrCodeURL string `json:"qr_code_url" db:"qr_code_url"`
}