package handlers

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/kefniark/mango/example/codegen/api"
	"github.com/kefniark/mango/example/codegen/database"
	"github.com/kefniark/mango/example/config"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	api.UnimplementedUsersServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func mapUserSQLToGrpc(user database.User) *api.UserData {
	return &api.UserData{
		Id:   user.ID.String(),
		Name: user.Name,
		Bio:  user.Bio,
	}
}

func mapUsersSQLToGrpc(users []database.User) []*api.UserData {
	res := []*api.UserData{}
	for _, user := range users {
		res = append(res, mapUserSQLToGrpc(user))
	}
	return res
}

// Get a User by Id.
func (service *UserService) Get(ctx context.Context, req *connect.Request[api.UserGetRequest]) (*connect.Response[api.UserData], error) {
	db := config.GetDB(ctx)

	id, err := uuid.FromBytes([]byte(req.Msg.GetId()))
	if err != nil {
		return nil, err
	}

	user, err := db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(mapUserSQLToGrpc(user)), nil
}

func getUUID(val string) uuid.UUID {
	if id, err := uuid.FromBytes([]byte(val)); err != nil {
		return uuid.New()
	} else {
		return id
	}
}

// Create or Update User.
func (service *UserService) Set(ctx context.Context, req *connect.Request[api.UserSetRequest]) (*connect.Response[api.UserData], error) {
	db := config.GetDB(ctx)
	id := getUUID(req.Msg.GetId())

	user, err := db.SetUser(ctx, database.SetUserParams{
		ID:   id,
		Name: req.Msg.GetName(),
		Bio:  req.Msg.GetBio(),
	})

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(mapUserSQLToGrpc(user)), nil
}

// Delete a User.
func (service *UserService) Delete(ctx context.Context, req *connect.Request[api.UserGetRequest]) (*connect.Response[emptypb.Empty], error) {
	db := config.GetDB(ctx)
	id := getUUID(req.Msg.GetId())

	err := db.DeleteUser(ctx, id)
	return nil, err
}

// const defaultOffset = int64(0)
// const defaultLimit = int64(50)

// List multiple Users (filter, paginate).
func (service *UserService) Search(ctx context.Context, req *connect.Request[api.UserSearchRequest]) (*connect.Response[api.UserSearchResponse], error) {
	db := config.GetDB(ctx)
	count, err := db.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	res, err := db.SearchUsers(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.UserSearchResponse{
		Users: mapUsersSQLToGrpc(res),
		Total: count,
	}), nil
}
