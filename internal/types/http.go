package internal

type ResponseStatus string

type ResponseBase struct {
	Status ResponseStatus `json:"status"`
	Error any `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusError ResponseStatus = "error"
)

var (
	SuccessResponse = ResponseBase{
		Status : ResponseStatusSuccess,
	}
)