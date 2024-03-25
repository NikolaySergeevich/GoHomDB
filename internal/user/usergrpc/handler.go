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

	// TODO implement me
	panic("implement me")
}

func (h Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
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
	// TODO implement me
	return &pb.Empty{}, nil
}

func (h Handler) ListUsers(
	ctx context.Context,
	in *pb.Empty,
) (*pb.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	// TODO implement me
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}