package service

import (
	"github.com/Mu-Exchange/Mu-End/common"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/game/pkg/base"
)

func init() {
	base.RegisterComponents(NewBaseService, NewGameService, NewTGGameService, NewUserService, NewStakingService)
}

func NewUserService(cfg repo.CommonComponents) pb.UserService {
	return pb.NewUserService(common.UserService, cfg.Client)
}
