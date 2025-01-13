package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/romanmufid16/belajar-golang-restful-api/helper"
	"github.com/romanmufid16/belajar-golang-restful-api/model/web"
	"github.com/romanmufid16/belajar-golang-restful-api/service"
	"net/http"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
