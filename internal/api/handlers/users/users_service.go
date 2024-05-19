package users

import (
	"github.com/kefniark/go-web-server/gen/api"
	"github.com/kefniark/go-web-server/gen/db"
	"github.com/rs/zerolog"
)

type UserService struct {
	db     *db.Queries
	logger zerolog.Logger
	api.UnimplementedUsersServer
}

func NewUserService(db *db.Queries, logger *zerolog.Logger) *UserService {
	return &UserService{
		db:     db,
		logger: logger.With().Str("service", "UserService").Logger(),
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
