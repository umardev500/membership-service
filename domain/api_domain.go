package domain

type APIResponse struct {
	StatusCode int64       `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}
