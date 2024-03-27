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
	usList, err := h.client.ListUsers(context.TODO(), &pb.Empty{})
	if err != nil {
		slog.Error("gRPC GetUsers client", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var resStr string
	for _, v := range usList.Users {
		timeCreate := "Дата создания: " + v.CreatedAt
		timeUpdate := "Дата создания: " + v.UpdatedAt
		resStr = resStr + "ID: " + v.Id + "\n" + "Пользователь: " + v.Username + "\nПароль: " + v.Password + "\n" + timeCreate + "\n" + timeUpdate + "\n"
	
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resStr))
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
	req := pb.GetUserRequest{Id: id}
	us, err := h.client.GetUser(context.TODO(), &req)
	if err != nil{
		slog.Error("gRPC getuser client", slog.String("err", err.Error()))
		w. WriteHeader(http.StatusInternalServerError)
	}
	var resStr string
	if us != nil{
		timeCreate := "Дата создания: " + us.CreatedAt
		timeUpdate := "Дата создания: " + us.UpdatedAt
		resStr = resStr + "ID: " + us.Id + "\n" + "Пользователь: " + us.Username + "\nПароль: " + us.Password + "\n" + timeCreate + "\n" + timeUpdate + "\n"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resStr))
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *usersHandler) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	req := pb.UpdateUserRequest{Id: id}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("json.NewDecoder Decode", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := h.client.UpdateUser(context.TODO(), &req); err != nil{
		slog.Error("gRPC Updateuser client", slog.String("err", err.Error()))
		w. WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
