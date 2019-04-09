package main

import "golang.org/x/net/context"
import pb "github.com/binhgo/GoMicro-User/proto/user"

type Handler struct {
	repo Repository
	tokenService Authable
}


func (han *Handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := han.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (han *Handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := han.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}


func (han *Handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := han.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}

func (han *Handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := han.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (han *Handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
