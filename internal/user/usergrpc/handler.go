package usergrpc

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gohomdb/internal/database"
	"gohomdb/pkg/pb"
)

var _ pb.UserServiceServer = (*Handler)(nil)

func New(usersRepository usersRepository, timeout time.Duration) *Handler {
	return &Handler{usersRepository: usersRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedUserServiceServer
	usersRepository usersRepository
	timeout         time.Duration
}

func (h Handler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	var us database.CreateUserReq
	us.ID = uuid.New()
	us.Username = in.Username
	us.Password = in.Password
	if _, err := h.usersRepository.Create(ctx, us); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h Handler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	idUU, err := uuid.Parse(in.Id)
	if err != nil{
		return nil, fmt.Errorf("uuid Pars: %w", err)
	}
	us, err := h.usersRepository.FindByID(ctx, idUU)// !
	if err != nil{
		return nil, err
	}
	user := pb.User{
		Id: us.ID.String(), 
		Username: us.Username, 
		Password: us.Password, 
		CreatedAt: us.CreatedAt.String(), 
		UpdatedAt: us.UpdatedAt.String()}
	return &user, nil
}

func (h Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	idUUID, err := uuid.Parse(in.Id)
	if err != nil{
		return nil, fmt.Errorf("not correctly string UUID: %w", err)
	}

	req := database.CreateUserReq{
		ID: idUUID,
		Username: in.Username,
		Password: in.Password,
	}
	_, errr :=  h.usersRepository.Create(ctx, req)
	if errr != nil{
		return nil, err
	}
	// TODO implement me
	return &pb.Empty{}, nil
}

func (h Handler) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest,) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	idUU, err := uuid.Parse(in.Id)
	if err != nil{
		return nil, fmt.Errorf("uuid Pars: %w", err)
	}
	if err := h.usersRepository.DeleteByUserID(ctx, idUU); err != nil{
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h Handler) ListUsers(ctx context.Context, in *pb.Empty,) (*pb.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	listUsers, err := h.usersRepository.FindAll(ctx)
	if err != nil{
		return nil, err
	}
	UsList := make([]*pb.User, 0)
	for _, v := range listUsers {
		UsList = append(UsList, &pb.User{
			Id: v.ID.String(), 
			Username: v.Username, 
			Password: v.Password,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		})
	}
	return &pb.ListUsersResponse{Users: UsList}, nil
}
