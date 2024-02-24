package handler

import (
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/utils/context"
	"github.com/Mu-Exchange/Mu-End/common/validator"
	"github.com/gin-gonic/gin"
)

func (s *Server) GameRandom(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	req := &model.GameRandomReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.gameService.GameRandom(ctx.Ctx(), &pb.GameRandomReq{
		Level:  uint64(req.Level),
		UserId: token.UserId})
}
func (s *Server) GameGetCode(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.gameService.GameGetInviteCode(ctx.Ctx(), &pb.UserIdReq{
		Id: token.UserId})
}

func (s *Server) LeaderBoard(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token, err := s.ExtractTokenMetadata(c)
	userId := ""
	if err == nil {
		userId = token.UserId
	}
	return s.gameService.LeaderBoard(ctx.Ctx(), &pb.LeaderBoardReq{UserId: userId, Limit: 20})
}

func (s *Server) ScoreList(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)

	req := &model.KindPage{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.gameService.ScoreList(ctx.Ctx(), &pb.ScoreListReq{
		UserId: token.UserId, PageNo: req.PageNo, PageSize: req.PageSize, ScoreType: uint64(req.Kind)})
}

func (s *Server) ConsumeList(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)

	req := &model.KindPage{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.gameService.ConsumeList(ctx.Ctx(), &pb.ConsumeListReq{
		UserId: token.UserId, PageNo: req.PageNo, PageSize: req.PageSize, ConsumeType: uint64(req.Kind)})
}

func (s *Server) MissionProfile(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.gameService.MissionProfile(ctx.Ctx(), &pb.UserIdReq{Id: token.UserId})
}
