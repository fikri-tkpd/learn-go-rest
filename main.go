package main

import (
	"net/http"
	"todo/app"
	"todo/controller"
	"todo/model/domain"
	"todo/repository"
	"todo/service"
)

func main() {
	var data []domain.Todo

	todoRepo := repository.NewToDoRepository()
	todoService := service.NewToDoService(todoRepo, &data)
	todoController := controller.NewToDoController(todoService)
	router := app.NewRouter(todoController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
