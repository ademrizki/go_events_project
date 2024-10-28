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
	User       any    `json:"users,omitempty"`
}
