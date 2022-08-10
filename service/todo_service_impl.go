package service

import (
	"context"
	"encoding/json"
	"todo/model/domain"
	"todo/model/web"
	"todo/repository"
)

type ToDoServiceImpl struct {
	repository repository.ToDoRepository
	data       *[]domain.Todo
}

func NewToDoService(repo repository.ToDoRepository, data *[]domain.Todo) ToDoService {
	return &ToDoServiceImpl{
		repository: repo,
		data:       data,
	}
}

func (service *ToDoServiceImpl) Save(c context.Context, req web.ToDoCreateRequest) web.ToDoSingleResponse {
	todo := domain.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	todo = service.repository.Save(c, service.data, todo)
	return web.ToDoSingleResponse{
		StatusCode: 200,
		Message:    "OK",
		Data:       todo,
	}
}

func (service *ToDoServiceImpl) Read(c context.Context, req web.ToDoCreateRequest) web.ToDoMultiResponse {
	todos := service.repository.Read(c, service.data)
	return web.ToDoMultiResponse{
		StatusCode: 200,
		Message:    "OK",
		Data:       todos,
	}
}

func (service *ToDoServiceImpl) Update(c context.Context, req web.ToDoPutRequest) web.ToDoSingleResponse {
	todo := domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Id:          req.Id,
	}
	service.repository.Update(c, service.data, todo)
	json.Marshal(&todo)
	return web.ToDoSingleResponse{
		StatusCode: 200,
		Message:    "OK",
		Data:       todo,
	}
}

func (service *ToDoServiceImpl) Delete(c context.Context, req web.ToDoDeleteRequest) web.ToDoNoResponse {
	id := req.Id
	service.repository.Delete(c, service.data, id)
	return web.ToDoNoResponse{
		StatusCode: 204,
		Message:    "No Content",
	}
}
