package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/go_rest/internal/handlers"
	"github.com/morheus9/go_rest/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	UsersURL = "/users"
	UserURL  = "/users/:id"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, UsersURL, h.GetList)
	router.HandlerFunc(http.MethodPost, UsersURL, h.CreateUser)
	router.HandlerFunc(http.MethodPost, UserURL, h.GetUserByUUID)
	router.HandlerFunc(http.MethodPut, UserURL, h.UpdateUser)
	router.HandlerFunc(http.MethodPatch, UserURL, h.PartiallyUpdateUser)
	router.HandlerFunc(http.MethodDelete, UserURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("this is a list of users"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("GetUserByUUID"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("CreateUser"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("UpdateUser"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("PartiallyUpdateUser"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("DeleteUser"))
}
