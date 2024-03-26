package handler

import (
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/utils/context"
	"github.com/Mu-Exchange/Mu-End/common/validator"
	"github.com/gin-gonic/gin"
)

func (s *Server) TGGameRandom(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	req := &model.GameRandomReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.gameService.TGGameRandom(ctx.Ctx(), &pb.GameRandomReq{
		Level:  uint64(req.Level),
		UserId: token.UserId})
}

func (s *Server) TGLeaderBoard(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token, err := s.TGExtractTokenMetadata(c)
	userId := ""
	if err == nil {
		userId = token.UserId
	}
	return s.gameService.TGLeaderBoard(ctx.Ctx(), &pb.LeaderBoardReq{UserId: userId, Limit: 20})
}

func (s *Server) TGMissionProfile(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.gameService.MissionProfile(ctx.Ctx(), &pb.UserIdReq{Id: token.UserId})
}

func (s *Server) BuyDice(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.BuyVipReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.userService.SetUserDiceMint(ctx.Ctx(), &pb.SetUserDiceMintReq{Id: token.UserId, Hash: req.Hash})
}

func (s *Server) BuyVip(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.BuyVipReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.userService.SetUserVIP(ctx.Ctx(), &pb.SetUserVIPReq{Id: token.UserId, Hash: req.Hash})
}

func (s *Server) Game24H(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.gameService.Game24H(ctx.Ctx(), &pb.Game24HReq{UserId: token.UserId})
}
