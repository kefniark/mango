package users

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
)

const defaultOffset = int64(0)
const defaultLimit = int64(50)

// List multiple Users (filter, paginate).
func (service *UserService) Search(
	ctx context.Context,
	req *connect.Request[api.UserSearchRequest],
) (*connect.Response[api.UserSearchResponse], error) {
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
