package v1

import (
	"context"
	"encoding/json"
	"gohomdb/pkg/pb"
	"log/slog"
	"net/http"
)

func newUsersHandler(usersClient usersClient) *usersHandler {
	return &usersHandler{client: usersClient}
}

type usersHandler struct {
	client usersClient
}

func (h *usersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}
/*
Создать нового пользователя
*/
func (h *usersHandler) PostUsers(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("json.NewDecoder Decode", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := h.client.CreateUser(context.TODO(), &req); err != nil{
		slog.Error("gRPS create client", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *usersHandler) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	req := pb.DeleteUserRequest{Id: id}
	if _, err := h.client.DeleteUser(context.TODO(), &req); err != nil{
		slog.Error("gRPS delete client", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *usersHandler) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *usersHandler) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}
