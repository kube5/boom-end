package handler

import (
	"github.com/Mu-Exchange/Mu-End/api/pkg/base"
	"github.com/Mu-Exchange/Mu-End/common"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
)

func init() {
	base.RegisterComponents(NewRedis,
		NewTokenCache,
		NewUserService,
		NewGameService,
		New)
}

func NewUserService(cfg repo.CommonComponents) pb.UserService {
	return pb.NewUserService(common.UserService, cfg.Client)
}

func NewGameService(cfg repo.CommonComponents) pb.GameService {
	return pb.NewGameService(common.GameService, cfg.Client)
}
