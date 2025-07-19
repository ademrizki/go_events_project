package models

type ErrorResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}

type EventsResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
	Event      any    `json:"events,omitempty"`
}

type UsersResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

type LoginResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
	Token      any    `json:"token,omitempty"`
}
