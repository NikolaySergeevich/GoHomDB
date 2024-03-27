package linkgrpc

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gohomdb/internal/database"
	"gohomdb/pkg/pb"
)

var _ pb.LinkServiceServer = (*Handler)(nil)

func New(linksRepository linksRepository, timeout time.Duration) *Handler {
	return &Handler{linksRepository: linksRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedLinkServiceServer
	linksRepository linksRepository
	timeout         time.Duration
}

func (h Handler) GetLinkByUserID(ctx context.Context, id *pb.GetLinksByUserId) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	list, err := h.linksRepository.FindByUserID(ctx, id.UserId)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("empty response from DB")
	}
	listLinks := make([]*pb.Link, 0)
	for _, v := range list {
		listLinks = append(listLinks, &pb.Link{
			Id:        v.ID.String(),
			Title:     v.Title,
			Url:       v.URL,
			Images:    v.Images,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
			UserId:    v.UserID,
		})
	}
	return &pb.ListLinkResponse{Links: listLinks}, nil
}

func (h Handler) mustEmbedUnimplementedLinkServiceServer() {
	// TODO implement me
	panic("implement me")
}

func (h Handler) CreateLink(ctx context.Context, request *pb.CreateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "link is empty")
	}

	reqToDB := database.CreateLinkReq{
		ID:     primitive.NewObjectID(),
		URL:    request.Url,
		Title:  request.Title,
		Tags:   request.Tags,
		Images: request.Images,
		UserID: request.UserId,
	}
	if _, err := h.linksRepository.Create(ctx, reqToDB); err != nil {
		return nil, err
	}

	// TODO implement me
	return &pb.Empty{}, nil
}

func (h Handler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	idPrimit, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "primitive Parse cannot")
	}
	linkDB, errr := h.linksRepository.FindByID(ctx, idPrimit)
	if errr != nil {
		return nil, errr
	}
	res := pb.Link{
		Id:        linkDB.ID.String(),
		Title:     linkDB.Title,
		Url:       linkDB.URL,
		Images:    linkDB.Images,
		Tags:      linkDB.Tags,
		UserId:    linkDB.UserID,
		CreatedAt: linkDB.CreatedAt.String(),
		UpdatedAt: linkDB.UpdatedAt.String(),
	}

	return &res, nil
}

func (h Handler) UpdateLink(ctx context.Context, request *pb.UpdateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	idPrimit, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "primitive Parse cannot")
	}

	req := database.UpdateLinkReq{
		ID: idPrimit,
		URL: request.Url,
		Title: request.Title,
		Images: request.Images,
		Tags: request.Tags,
		UserID: request.UserId,
	}
	if _, err := h.linksRepository.Update(ctx, req); err != nil{
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h Handler) DeleteLink(ctx context.Context, request *pb.DeleteLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	idPrimit, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "primitive Parse cannot")
	}
	if err := h.linksRepository.Delete(ctx, idPrimit); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h Handler) ListLinks(ctx context.Context, request *pb.Empty) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request is empty")
	}
	listDB, err := h.linksRepository.FindAll(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "dont try find all from links")
	}
	if len(listDB) == 0 {
		return nil, status.Error(codes.InvalidArgument, "links in db is empty")
	}
	res := make([]*pb.Link, 0)
	for _, linkDB := range listDB {
		res = append(res, &pb.Link{
			Id:        linkDB.ID.String(),
			Title:     linkDB.Title,
			Url:       linkDB.URL,
			Images:    linkDB.Images,
			Tags:      linkDB.Tags,
			UserId:    linkDB.UserID,
			CreatedAt: linkDB.CreatedAt.String(),
			UpdatedAt: linkDB.UpdatedAt.String(),
		})
	}

	return &pb.ListLinkResponse{Links: res}, nil
}
