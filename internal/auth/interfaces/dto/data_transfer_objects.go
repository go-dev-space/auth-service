package dto

type Response struct {
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	Data       []map[string]any `json:"data"`
	Error      bool             `json:"error"`
}
