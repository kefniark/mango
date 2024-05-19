package users

import (
	"context"

	"connectrpc.com/connect"
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/moroz/uuidv7-go"
)

// Create or Update User.
func (service *UserService) Set(ctx context.Context, req *connect.Request[api.UserSetRequest]) (*connect.Response[api.UserData], error) {
	var id string
	if req.Msg.Id == nil {
		id = uuidv7.Generate().String()
	} else {
		id = req.Msg.GetId()
	}

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
