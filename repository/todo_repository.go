package repository

import (
	"context"
	"todo/model/domain"
)

type ToDoRepository interface {
	Save(c context.Context, data *[]domain.Todo, todo domain.Todo) domain.Todo
	Read(c context.Context, data *[]domain.Todo) []domain.Todo
	Update(c context.Context, data *[]domain.Todo, todo domain.Todo) domain.Todo
	Delete(c context.Context, data *[]domain.Todo, id int)
}
