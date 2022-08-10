package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ToDoController interface {
	Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Read(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}
