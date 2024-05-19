package users

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
)

// Delete a User.
func (service *UserService) Delete(
	ctx context.Context,
	req *connect.Request[api.UserGetRequest],
) (*connect.Response[api.UserEmptyResponse], error) {
	err := service.db.DeleteUser(ctx, req.Msg.GetId())
	return nil, err
}
