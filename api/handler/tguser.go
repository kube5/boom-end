package handler

import (
	"fmt"
	"os"
	"strings"

	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	pb "github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/utils/context"
	"github.com/Mu-Exchange/Mu-End/common/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) TGMetaMaskLoginPre(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.MetaMaskLoginPreReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.userService.LoginPreByTG(ctx.Ctx(), &pb.LoginPreByTGReq{
		PublicAddress: req.PublicAddress})
}

func (s *Server) TGMetaMaskLogin(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.MetaMaskLoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.userService.LoginByTg(ctx.Ctx(), &pb.LoginTgReq{
		PublicAddress: req.PublicAddress,
		Signature:     req.Signature})
}

func (s *Server) TGLoginInternal(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.LoginInternalReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	return s.userService.LoginInternal(ctx.Ctx(), &pb.LoginInternalReq{
		Address: req.Address})
}

func (s *Server) TGRefreshToken(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		return nil, derrors.ErrRefreshToken
	}
	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_TOKEN")), nil
	})
	if err != nil && strings.Contains(err.Error(), "expired") {
		return nil, derrors.ErrRefreshTokenExpired
	}
	if err != nil {
		return nil, derrors.ErrRefreshToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, derrors.ErrAuthToken
	}
	refreshUuid, ok := claims["refresh_uuid"].(string)
	if !ok || refreshUuid == "" {
		return nil, derrors.ErrRefreshToken
	}
	accessUuid, ok := claims["access_uuid"].(string)
	if !ok || accessUuid == "" {
		return nil, derrors.ErrRefreshToken
	}
	userId := fmt.Sprintf("%s", claims["user_id"])
	s.Logger.Infof("userid:%s,refreshuuid:%s,accessuuid:%s", userId, refreshUuid, accessUuid)
	return s.userService.TgRefreshToken(ctx.Ctx(), &pb.RefreshTokenReq{UserId: userId, RefreshUuid: refreshUuid, AccessUuid: accessUuid})
}

func (s *Server) TGLogout(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	token := c.MustGet(AccessDetails).(*model.AccessDetails)
	return s.userService.TGLogout(ctx.Ctx(), &pb.LogoutReq{UserId: token.UserId, AccessUuid: token.AccessUuid})
}

func (s *Server) TGProfile(ctx *context.Context, c *gin.Context) (res interface{}, err error) {
	req := &model.UserIdReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, derrors.Wrap(derrors.ErrRequestParameter, validator.GetErrorMsg(req, err))
	}
	pbReq := &pb.UserMutualReq{}
	token, err := s.TGExtractTokenMetadata(c)
	if err == nil {
		pbReq.Id = token.UserId
	}
	if len(req.TargetUserId) != 0 {
		pbReq.TargetUserId = req.TargetUserId
	} else {
		if err != nil {
			return nil, err
		}
	}
	return s.userService.ProfileTG(ctx.Ctx(), pbReq)
}
