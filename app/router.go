package app

import (
	"todo/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoController controller.ToDoController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", todoController.Read)
	router.POST("/", todoController.Create)
	router.PUT("/", todoController.Update)
	router.DELETE("/", todoController.Delete)

	return router
}
