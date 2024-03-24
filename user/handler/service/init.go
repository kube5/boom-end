package service

import (
	"github.com/Mu-Exchange/Mu-End/common"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/user/pkg/base"
)

func init() {
	base.RegisterComponents(
		NewBaseService,
		NewAuthService,
		NewTgUserService,
		NewGameService,
	)
}

func NewGameService(cfg repo.CommonComponents) pb.GameService {
	return pb.NewGameService(common.GameService, cfg.Client)
}
