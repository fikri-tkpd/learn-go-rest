package service

import (
	"context"
	"todo/model/web"
)

type ToDoService interface {
	Save(c context.Context, req web.ToDoCreateRequest) web.ToDoSingleResponse
	Read(c context.Context, req web.ToDoCreateRequest) web.ToDoMultiResponse
	Update(c context.Context, req web.ToDoPutRequest) web.ToDoSingleResponse
	Delete(c context.Context, req web.ToDoDeleteRequest) web.ToDoNoResponse
}
