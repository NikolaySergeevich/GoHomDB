package linkgrpc

import (
	"context"
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
	// TODO implement me
	panic("implement me")
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
		ID: primitive.NewObjectID(),
		URL: request.Url,
		Title: request.Title,
		Tags: request.Tags,
		Images: request.Images,
		UserID: request.UserId,
	}
	if _, err := h.linksRepository.Create(ctx,reqToDB); err != nil{
		return nil, err
	}

	// TODO implement me
	return &pb.Empty{}, nil
}

func (h Handler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	idPrimit, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, "primitive Parse cannot") 
	}
	linkDB, errr := h.linksRepository.FindByID(ctx, idPrimit)
	if err != nil{
		return nil, errr
	}
	res := pb.Link{
		Id: linkDB.ID.String(),
		Title: linkDB.Title,
		Url: linkDB.URL,
		Images: linkDB.Images,
		Tags: linkDB.Tags,
		UserId: linkDB.UserID,
		CreatedAt: linkDB.CreatedAt.String(),
		UpdatedAt: linkDB.UpdatedAt.String(),
	}
	
	return &res, nil
}

func (h Handler) UpdateLink(ctx context.Context, request *pb.UpdateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (h Handler) DeleteLink(ctx context.Context, request *pb.DeleteLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	idPrimit, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil{
		return nil, status.Error(codes.InvalidArgument, "primitive Parse cannot") 
	}
	if err := h.linksRepository.Delete(ctx, idPrimit); err != nil{
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h Handler) ListLinks(ctx context.Context, request *pb.Empty) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
