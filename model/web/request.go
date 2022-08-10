package web

type ToDoCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ToDoPutRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Id          int    `json:"id"`
}

type ToDoDeleteRequest struct {
	Id int `json:"id"`
}
