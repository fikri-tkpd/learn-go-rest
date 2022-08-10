package web

import "todo/model/domain"

type ToDoSingleResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       domain.Todo `json:"data"`
}

type ToDoMultiResponse struct {
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	Data       []domain.Todo `json:"data"`
}

type ToDoNoResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
