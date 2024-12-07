package types

type ApiResponse struct {
	Status  int         `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
	Result  string      `json:"result,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
