package dto

type JSONResponse struct {
	Meta  any `json:"meta,omitempty"`
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}
