package v1

import (
	"context"
	"encoding/json"
	"gohomdb/pkg/pb"
	"log/slog"
	"net/http"
)

func newLinksHandler(linksClient linksClient) *linksHandler {
	return &linksHandler{client: linksClient}
}

type linksHandler struct {
	client linksClient
}

func (h *linksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) PostLinks(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	req := pb.CreateLinkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("json.NewDecoder Decode", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if _, err := h.client.CreateLink(context.TODO(), &req); err != nil{
		slog.Error("gRPS create link", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *linksHandler) DeleteLinksId(w http.ResponseWriter, r *http.Request, id string) {
	req := pb.DeleteLinkRequest{Id: id}
	if _, err := h.client.DeleteLink(context.TODO(), &req); err != nil{
		slog.Error("gRPS delete link", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *linksHandler) GetLinksId(w http.ResponseWriter, r *http.Request, id string) {
	req := pb.GetLinkRequest{Id: id}
	link, err := h.client.GetLink(context.TODO(), &req)
	if err != nil{
		slog.Error("gRPS GetLinksId link", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var resStr string
	if link != nil{
		timeCreate := "Дата создания: " + link.CreatedAt
		timeUpdate := "Дата создания: " + link.UpdatedAt
		resStr = resStr + "ID: " + link.Id + "\n" + "Title: " + link.Title + "\nUrl: " + link.Url + 
				"\nUserId: " + link.UserId + "\n" + timeCreate + "\n" + timeUpdate + "\n"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resStr))
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) PutLinksId(w http.ResponseWriter, r *http.Request, id string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *linksHandler) GetLinksUserUserID(w http.ResponseWriter, r *http.Request, userID string) {
	// TODO implement me
	w.WriteHeader(http.StatusNotImplemented)
}

