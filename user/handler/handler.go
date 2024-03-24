package handler

import (
	"context"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/user/handler/service"
	"github.com/Mu-Exchange/Mu-End/user/pkg/base"
	aws2 "github.com/aws/aws-sdk-go-v2/aws"
	credentials2 "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var _ proto.UserHandler = (*userHandler)(nil)

func init() {
	base.RegisterComponents(NewUserHandler)
}

type userHandler struct {
	repo.CommonComponents
	us            service.UserService
	tus           service.TgUserService
	s3Uploader    *s3manager.Uploader
	presignClient *s3.PresignClient
}

func (i *userHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserReq, empty *proto.Empty) error {
	return i.us.UpdateUserByWallet(ctx, &dto.User{
		Wallet:            req.Wallet,
		StakingAmountETH:  req.EthStakeAmount,
		StakingAmountUSDB: req.UsdbStakeAmount,
		MintDice:          req.MintDice,
		DiceSpeed:         uint64(req.Dice),
	})
}

func (i *userHandler) FindUserIdByWallet(ctx context.Context, wallet *proto.Wallet, resp *proto.UserIdResp) error {
	userId, err := i.us.FindUserIdByWallet(ctx, wallet.Address)
	if err != nil {
		return err
	}
	resp.UserId = userId
	return nil
}

func (i *userHandler) LoginInternal(ctx context.Context, req *proto.LoginInternalReq, resp *proto.LoginByMetaMaskResp) error {
	token, err := i.us.LoginInternal(ctx, req.Address)
	if err != nil {
		return err
	}
	resp.AccessToken = token.AccessToken
	resp.AtExpires = int64(token.AtExpires)
	resp.RefreshToken = token.RefreshToken
	resp.RtExpires = int64(token.RtExpires)
	return nil
}

func (i *userHandler) Profile(ctx context.Context, req *proto.UserMutualReq, resp *proto.ProfileResp) error {
	id := req.Id
	if len(req.TargetUserId) != 0 {
		id = req.TargetUserId
	}
	var user *dto.User
	var err error
	if len(id) == 0 {
		user, err = i.us.QueryUserByWallet(ctx, req.Wallet)
		if err != nil {
			return err
		}
	} else {
		user, err = i.us.QueryUser(ctx, id)
		if err != nil {
			return err
		}
	}

	resp.Id = user.UUID
	resp.UserInfo = &proto.BasicUserInfo{
		Username:          user.Username,
		Wallet:            user.Wallet,
		UserId:            user.UUID,
		ProfileUrl:        user.ProfileUrl,
		StakingAmountEth:  user.StakingAmountETH,
		StakingAmountUsdb: user.StakingAmountUSDB,
		MintDice:          user.MintDice,
		DiceSpeed:         user.DiceSpeed,
	}

	return nil
}

func NewUserHandler(commonComponents repo.CommonComponents,
	us service.UserService, tus service.TgUserService) (proto.UserHandler, error) {
	s3Config := commonComponents.Cfg.S3
	creds := credentials.NewStaticCredentials(s3Config.AccessKeyId, s3Config.AccessKeySecret, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.Region),
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(s3.Options{
		Credentials: credentials2.StaticCredentialsProvider{Value: aws2.Credentials{
			AccessKeyID:     s3Config.AccessKeyId,
			SecretAccessKey: s3Config.AccessKeySecret,
		}},
		Region: s3Config.Region,
	})
	presignClient := s3.NewPresignClient(s3Client)

	return &userHandler{
		CommonComponents: commonComponents,
		us:               us,
		tus:              tus,
		s3Uploader:       s3manager.NewUploader(sess),
		presignClient:    presignClient,
	}, nil
}

func (i *userHandler) LoginPreByTG(ctx context.Context, req *proto.LoginPreByTGReq, resp *proto.LoginPreByTGResp) error {
	nonce, err := i.tus.LoginPreByMetaMask(ctx, req.PublicAddress)
	if err != nil {
		return err
	}
	resp.Nonce = nonce
	return nil
}

func (i *userHandler) ProfileTG(ctx context.Context, req *proto.UserMutualReq, resp *proto.ProfileResp) error {
	id := req.Id
	if len(req.TargetUserId) != 0 {
		id = req.TargetUserId
	}
	var user *dto.TgUser
	var err error
	if len(id) == 0 {
		user, err = i.tus.QueryTgUserByWallet(ctx, req.Wallet)
		if err != nil {
			return err
		}
	} else {
		user, err = i.tus.QueryTgUser(ctx, id)
		if err != nil {
			return err
		}
	}

	resp.Id = user.UUID
	resp.UserInfo = &proto.BasicUserInfo{
		Username:   user.Tgname,
		Wallet:     user.Wallet,
		UserId:     user.UUID,
		ProfileUrl: user.ProfileUrl,
		MintDice:   user.MintDice,
		DiceSpeed:  user.DiceSpeed,
	}
	return nil
}

func (i *userHandler) LoginByTg(ctx context.Context, req *proto.LoginTgReq, resp *proto.LoginTgResp) error {
	token, err := i.tus.LoginInternal(ctx, req.PublicAddress)
	if err != nil {
		return err
	}
	resp.AccessToken = token.AccessToken
	resp.AtExpires = int64(token.AtExpires)
	resp.RefreshToken = token.RefreshToken
	resp.RtExpires = int64(token.RtExpires)
	return nil
}

func (i *userHandler) LoginPreByMetaMask(ctx context.Context, req *proto.LoginPreByMetaMaskReq, resp *proto.LoginPreByMetaMaskResp) error {
	nonce, err := i.us.LoginPreByMetaMask(ctx, req.PublicAddress)
	if err != nil {
		return err
	}
	resp.Nonce = nonce
	return nil
}

func (i *userHandler) LoginByMetaMask(ctx context.Context, req *proto.LoginByMetaMaskReq, resp *proto.LoginByMetaMaskResp) error {
	token, err := i.us.LoginByMetaMask(ctx, req.PublicAddress, req.Signature)
	if err != nil {
		return err
	}
	resp.AccessToken = token.AccessToken
	resp.AtExpires = int64(token.AtExpires)
	resp.RefreshToken = token.RefreshToken
	resp.RtExpires = int64(token.RtExpires)
	return nil
}

func (i *userHandler) RefreshToken(ctx context.Context, req *proto.RefreshTokenReq, resp *proto.RefreshTokenResp) error {
	token, err := i.us.RefreshToken(ctx, req.RefreshUuid, req.AccessUuid, req.UserId)
	if err != nil {
		return err
	}
	resp.AccessToken = token.AccessToken
	resp.AtExpires = int64(token.AtExpires)
	resp.RefreshToken = token.RefreshToken
	resp.RtExpires = int64(token.RtExpires)
	return nil
}

func (i *userHandler) Logout(ctx context.Context, req *proto.LogoutReq, empty *proto.Empty) error {
	return i.us.Logout(ctx, req.AccessUuid, req.UserId)
}

func (i *userHandler) TgRefreshToken(ctx context.Context, req *proto.RefreshTokenReq, resp *proto.RefreshTokenResp) error {
	token, err := i.tus.RefreshToken(ctx, req.RefreshUuid, req.AccessUuid, req.UserId)
	if err != nil {
		return err
	}
	resp.AccessToken = token.AccessToken
	resp.AtExpires = int64(token.AtExpires)
	resp.RefreshToken = token.RefreshToken
	resp.RtExpires = int64(token.RtExpires)
	return nil
}

func (i *userHandler) TGLogout(ctx context.Context, req *proto.LogoutReq, empty *proto.Empty) error {
	return i.tus.Logout(ctx, req.AccessUuid, req.UserId)
}

func (i *userHandler) Call(ctx context.Context, empty *proto.Empty, empty2 *proto.Empty) error {
	//TODO implement me
	panic("implement me")
}
