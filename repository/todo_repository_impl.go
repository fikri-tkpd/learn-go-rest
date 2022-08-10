package repository

import (
	"context"
	"todo/model/domain"
)

type ToDoRepositoryImpl struct {
}

func NewToDoRepository() ToDoRepository {
	return &ToDoRepositoryImpl{}
}

func (repo *ToDoRepositoryImpl) Save(c context.Context, data *[]domain.Todo, todo domain.Todo) domain.Todo {
	id := len(*data) + 1
	todo.Id = id
	*data = append(*data, todo)

	return todo
}

func (repo *ToDoRepositoryImpl) Read(c context.Context, data *[]domain.Todo) []domain.Todo {

	return *data
}

func (repo *ToDoRepositoryImpl) Update(c context.Context, data *[]domain.Todo, todo domain.Todo) domain.Todo {
	id := todo.Id
	foundData := (*data)[id-1]
	foundData.Id = id
	foundData.Description = todo.Description
	foundData.Title = todo.Title
	foundData.Status = todo.Status

	(*data)[id-1] = foundData

	return foundData
}
func (repo *ToDoRepositoryImpl) Delete(c context.Context, data *[]domain.Todo, id int) {
	index := id - 1
	if len(*data) > 1 {
		(*data) = append((*data)[:index], (*data)[index+1:]...)
	} else {
		var kosong []domain.Todo
		(*data) = kosong
	}
}
