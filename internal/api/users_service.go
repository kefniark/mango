package api

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/kefniark/go-web-server/internal/core"
	"github.com/moroz/uuidv7-go"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	db     *db.Queries
	logger zerolog.Logger
	api.UnimplementedUsersServer
}

func NewUserService(options *core.ServerOptions) *UserService {
	return &UserService{
		db:     options.DB,
		logger: options.Logger.With().Str("service", "UserService").Logger(),
	}
}

func mapUserSQLToGrpc(user db.User) *api.UserData {
	return &api.UserData{
		Id:   user.ID.(string),
		Name: user.Name,
		Bio:  user.Bio,
	}
}

func mapUsersSQLToGrpc(users []db.User) []*api.UserData {
	res := []*api.UserData{}
	for _, user := range users {
		res = append(res, mapUserSQLToGrpc(user))
	}
	return res
}

// Get a User by Id.
func (service *UserService) Get(ctx context.Context, req *connect.Request[api.UserGetRequest]) (*connect.Response[api.UserData], error) {
	user, err := service.db.GetUser(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(mapUserSQLToGrpc(user)), nil
}

// Create or Update User.
func (service *UserService) Set(ctx context.Context, req *connect.Request[api.UserSetRequest]) (*connect.Response[api.UserData], error) {
	var id string
	if req.Msg.Id == nil {
		id = uuidv7.Generate().String()
	} else {
		id = req.Msg.GetId()
	}

	// get and use User from Auth Middleware
	// userInfo, _ := authn.GetInfo(ctx).(core.AuthInfo)
	// fmt.Println("Set User", userInfo)

	user, err := service.db.SetUser(ctx, db.SetUserParams{
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
	err := service.db.DeleteUser(ctx, req.Msg.GetId())
	return nil, err
}

const defaultOffset = int64(0)
const defaultLimit = int64(50)

// List multiple Users (filter, paginate).
func (service *UserService) Search(ctx context.Context, req *connect.Request[api.UserSearchRequest]) (*connect.Response[api.UserSearchResponse], error) {
	count, err := service.db.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	res, err := service.db.SearchUsers(ctx, db.SearchUsersParams{
		Offset: defaultOffset,
		Limit:  defaultLimit,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api.UserSearchResponse{
		Users: mapUsersSQLToGrpc(res),
		Total: count,
	}), nil
}
