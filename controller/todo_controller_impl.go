package controller

import (
	"net/http"
	"todo/helper"
	"todo/model/web"
	"todo/service"

	"github.com/julienschmidt/httprouter"
)

type ToDoControllerImpl struct {
	todoService service.ToDoService
}

func NewToDoController(todoService service.ToDoService) ToDoController {
	return &ToDoControllerImpl{
		todoService: todoService,
	}
}

func (todoImpl *ToDoControllerImpl) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	todoRequest := web.ToDoCreateRequest{}
	helper.ReadFromBody(req, &todoRequest)

	response := todoImpl.todoService.Save(req.Context(), todoRequest)
	helper.WriteToResponseBody(writer, response)
}

func (todoImpl *ToDoControllerImpl) Read(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	response := todoImpl.todoService.Read(req.Context(), web.ToDoCreateRequest{})
	helper.WriteToResponseBody(writer, response)
}

func (todoImpl *ToDoControllerImpl) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	todoRequest := web.ToDoPutRequest{}
	helper.ReadFromBody(req, &todoRequest)
	response := todoImpl.todoService.Update(req.Context(), todoRequest)
	helper.WriteToResponseBody(writer, response)
}

func (todoImpl *ToDoControllerImpl) Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	todoRequest := web.ToDoDeleteRequest{}
	helper.ReadFromBody(req, &todoRequest)

	response := todoImpl.todoService.Delete(req.Context(), todoRequest)
	helper.WriteToResponseBody(writer, response)
}
