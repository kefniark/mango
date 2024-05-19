package users

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
)

// Get a User by Id.
func (service *UserService) Get(ctx context.Context, req *connect.Request[api.UserGetRequest]) (*connect.Response[api.UserData], error) {
	user, err := service.db.GetUser(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(mapUserSQLToGrpc(user)), nil
}
