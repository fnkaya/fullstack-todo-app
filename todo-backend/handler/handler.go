package handler

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"io"
	"net/http"
)

type IHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
	insert(w http.ResponseWriter, r *http.Request)
	getAll(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IService
}

func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	Cors(&w, r)
	switch r.Method {
	case http.MethodOptions:
		return
	case http.MethodPost:
		h.insert(w, r)
	case http.MethodGet:
		h.getAll(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (h *Handler) insert(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		serverError(w, err)
		return
	}
	defer r.Body.Close()

	todo := &model.Todo{}
	if err := json.Unmarshal(body, todo); err != nil {
		serverError(w, err)
		return
	}

	if err := h.service.Insert(*todo); err != nil {
		serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll()
	if err != nil {
		serverError(w, err)
		return
	}

	response, err := json.Marshal(todos)
	if err != nil {
		serverError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func NewHandler(service service.IService) IHandler {
	return &Handler{service: service}
}

func serverError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func Cors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}
