package services

import (
	"context"

	app "github.com/storyofhis/auth-service/repositories/proto"
	"github.com/storyofhis/auth-service/services"
)

type Server struct {
	app.UnimplementedAuthServer
}

func (s *Server) CreateUser(ctx context.Context, in *app.User) (*app.Empty, error) {
	err := services.CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Server) Login(ctx context.Context, in *app.Credentials) (res *app.User, err error) {
	res, err = services.Login(ctx, in)
	if err != nil {
		return res, err
	}
	return res, nil
}
