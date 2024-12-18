package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/go_rest/internal/handlers"
)

var _ handlers.Handler = &handler{}

const (
	UsersURL = "/users"
	UserURL  = "/users/:id"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(UsersURL, h.GetList)
	router.POST(UsersURL, h.CreateUser)
	router.POST(UserURL, h.GetUserByUUID)
	router.PUT(UserURL, h.UpdateUser)
	router.PATCH(UserURL, h.PartiallyUpdateUser)
	router.DELETE(UserURL, h.DeleteUser)

}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is a list of users"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("GetUserByUUID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("CreateUser"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("UpdateUser"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("PartiallyUpdateUser"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("DeleteUser"))
}
