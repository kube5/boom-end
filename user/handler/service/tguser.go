package service

import (
	"context"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/redis"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/Mu-Exchange/Mu-End/common/utils/metamask"
	"github.com/Mu-Exchange/Mu-End/common/utils/uuid"
	"github.com/Mu-Exchange/Mu-End/user/handler/cache"
	"github.com/Mu-Exchange/Mu-End/user/handler/dao"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/fx"
)

type TgUserService interface {
	QueryTgUser(ctx context.Context, TgUserId string) (*dto.TgUser, error)
	QueryTgUserByWallet(ctx context.Context, wallet string) (*dto.TgUser, error)
	LoginPreByMetaMask(ctx context.Context, publicAddress string) (string, error)
	LoginByMetaMask(ctx context.Context, publicAddress, signature string) (*model.TokenDetails, error)
	LoginInternal(ctx context.Context, address string) (*model.TokenDetails, error)
	UpdateTgUserByWallet(ctx context.Context, TgUser *dto.TgUser) error
	RefreshToken(ctx context.Context, refreshUuid, accessUuid string, TgUserId string) (*model.TokenDetails, error)
	Logout(ctx context.Context, givenUuid string, TgUserId string) error

	FindTgUserIdByWallet(ctx context.Context, wallet string) (string, error)
}

var _ TgUserService = (*TgAccount)(nil)

const DistributionLockExpire = 5 * time.Second
const InviteCodeLockKey = "{LOCK}-InviteCodeLock"

type TgAccount struct {
	BaseService
	repo.CommonComponents
	TgUserDao   dao.TgUserDao
	TgUserCache cache.TgUserCache
	ethClient   *ethclient.Client
	gameService proto.GameService
	lock        *redis.DistributedLock
}

func (a *TgAccount) QueryTgUser(ctx context.Context, TgUserId string) (*dto.TgUser, error) {
	TgUser, err := a.TgUserCache.GetTgUser(TgUserId)
	if err != nil {
		return nil, err
	}
	if TgUser != nil {
		return TgUser, nil
	}
	TgUser, err = a.TgUserDao.FindByTgUserID(TgUserId)
	if err != nil {
		return nil, err
	}
	if TgUser == nil {
		return nil, derrors.ErrUserNotExist
	}
	return TgUser, a.TgUserCache.SetTgUser(TgUser)
}

func (a *TgAccount) QueryTgUserByWallet(ctx context.Context, wallet string) (*dto.TgUser, error) {
	if len(wallet) == 0 {
		return nil, derrors.ErrUserNotExist
	}
	TgUser, err := a.TgUserDao.FindByAddress(wallet)
	if err != nil {
		return nil, err
	}
	return TgUser, nil
}

func NewTgUserService(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents,
	service BaseService,
	gameService proto.GameService,
	TgUserDao dao.TgUserDao,
	TgUserCache cache.TgUserCache,
) (TgUserService, error) {
	blockchain := cfg.Cfg.Blockchains[0]
	rpcClient, err := rpc.DialContext(context.Background(), blockchain.RpcUrls[0])
	if err != nil {
		return nil, err
	}
	lock, err := redis.NewDistributedLock(cfg.Cfg.Redis)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)
	ser := &TgAccount{
		BaseService:      service,
		CommonComponents: cfg,
		TgUserCache:      TgUserCache,
		TgUserDao:        TgUserDao,
		ethClient:        ethClient,
		gameService:      gameService,
		lock:             lock,
	}

	basic.RegisterHook(lc, sd, ser)
	return ser, nil
}

func (a *TgAccount) Start(sd fx.Shutdowner) error {
	return nil
}

func (a *TgAccount) Stop(sd fx.Shutdowner) error {
	return nil
}

func (a *TgAccount) LoginPreByMetaMask(ctx context.Context, publicAddress string) (string, error) {
	nonce := time.Now().GoString()
	signText := "Welcome to boomup.fun!\n\n" +
		"Click to sign in and accept the boomup.fun Terms of Service: https://boomup.fun/\n\n" +
		"This request will not trigger a blockchain transaction or cost any gas fees.\n\n" +
		"Your authentication status will reset after 24 hours.\n\n" +
		"Wallet address:\n" +
		publicAddress +
		"\n\nNonce:\n" +
		nonce
	if err := a.TgUserCache.SetMetamaskLoginNonce(publicAddress, signText); err != nil {
		return "", err
	}
	return signText, nil
}

func (a *TgAccount) LoginByMetaMask(ctx context.Context, publicAddress, signature string) (*model.TokenDetails, error) {
	if len(publicAddress) == 0 {
		return nil, derrors.ErrUserNotExist
	}
	nonce, err := a.TgUserCache.GetMetamaskLoginNonce(publicAddress)
	if err != nil {
		return nil, err
	}
	verify, err := metamask.SignVerify(publicAddress, []byte(nonce), signature)
	if err != nil {
		return nil, err
	}
	if !verify {
		return nil, derrors.ErrMetaMaskVerifySign
	}
	TgUser, err := a.TgUserDao.FindByAddress(publicAddress)
	if err != nil {
		return nil, err
	}
	if TgUser == nil {
		TgUser = &dto.TgUser{
			UUID:   uuid.GenUUID().Encode(),
			Wallet: publicAddress,
		}
		if err := a.CreateTgUser(ctx, TgUser); err != nil {
			return nil, err
		}
		a.Logger.Infof("New TgUser[%s] with address[%s] is registering", TgUser.UUID, TgUser.Wallet)
	}
	td, err := utils.CreateToken(TgUser.UUID)
	if err != nil {
		return nil, err
	}
	if err := a.TgUserCache.DeleteTgUserAuth(TgUser.UUID); err != nil {
		return nil, err
	}
	if err := a.TgUserCache.CreateAuth(TgUser.UUID, td); err != nil {
		return nil, err
	}
	a.Logger.Infof("New TgUser[%s] with address[%s] login in", TgUser.UUID, TgUser.Wallet)
	return td, a.TgUserCache.SetTgUser(TgUser)
}

func (a *TgAccount) LoginInternal(ctx context.Context, publicAddress string) (*model.TokenDetails, error) {
	TgUser, err := a.TgUserDao.FindByAddress(publicAddress)
	if err != nil {
		return nil, err
	}
	//if TgUser == nil {
	//	return nil, derrors.ErrTgUserNotExist
	//}
	if TgUser == nil {
		TgUser = &dto.TgUser{
			UUID:   uuid.GenUUID().Encode(),
			Wallet: publicAddress,
		}
		if err := a.CreateTgUser(ctx, TgUser); err != nil {
			return nil, err
		}
		a.Logger.Infof("New TgUser[%s] with address[%s] is registering", TgUser.UUID, TgUser.Wallet)
	}
	td, err := utils.CreateToken(TgUser.UUID)
	if err != nil {
		return nil, err
	}
	if err := a.TgUserCache.DeleteTgUserAuth(TgUser.UUID); err != nil {
		return nil, err
	}
	if err := a.TgUserCache.CreateAuth(TgUser.UUID, td); err != nil {
		return nil, err
	}
	a.Logger.Infof("!!internal login, TgUser[%s] with address[%s] login in", TgUser.UUID, TgUser.Wallet)
	return td, a.TgUserCache.SetTgUser(TgUser)
}

func (a *TgAccount) Logout(ctx context.Context, givenUuid, TgUserId string) error {
	deleted, err := a.TgUserCache.DeleteAuth(TgUserId, givenUuid)
	if deleted == 0 || err != nil {
		return derrors.ErrAuthToken
	}
	return err
}

func (a *TgAccount) RefreshToken(ctx context.Context, refreshUuid, accessUuid, TgUserId string) (*model.TokenDetails, error) {
	deleted, err := a.TgUserCache.DeleteAuth(TgUserId, refreshUuid)
	if deleted == 0 || err != nil {
		return nil, derrors.ErrRefreshTokenExpired
	}

	deleted, err = a.TgUserCache.DeleteAuth(TgUserId, accessUuid)
	if err != nil {
		return nil, derrors.ErrAuthValidationExpired
	}
	td, err := utils.CreateToken(TgUserId)
	if err != nil {
		return nil, err
	}

	if err := a.TgUserCache.CreateAuth(TgUserId, td); err != nil {
		return nil, err
	}

	return td, nil
}

func (a *TgAccount) UpdateTgUserByWallet(ctx context.Context, TgUser *dto.TgUser) error {
	if len(TgUser.Wallet) == 0 {
		return nil
	}

	oldTgUser, err := a.QueryTgUserByWallet(ctx, TgUser.Wallet)
	if err != nil {
		return err
	}
	if nil == oldTgUser {
		newTgUser := &dto.TgUser{
			UUID:      uuid.GenUUID().Encode(),
			Wallet:    TgUser.Wallet,
			MintDice:  TgUser.MintDice,
			DiceSpeed: TgUser.DiceSpeed,
		}
		if err := a.CreateTgUser(ctx, newTgUser); err != nil {
			return err
		}
		a.Logger.Infof("New TgUser[%s] with address[%s] is registering", TgUser.UUID, TgUser.Wallet)
		return a.TgUserCache.SetTgUser(newTgUser)
	} else {
		TgUser.UUID = oldTgUser.UUID
		if oldTgUser.MintDice {
			TgUser.MintDice = oldTgUser.MintDice
		}
		_, err := a.TgUserDao.UpdateTgUser(TgUser)
		if err != nil {
			return err
		}
		return a.TgUserCache.SetTgUser(TgUser)
	}
}

func (a *TgAccount) CreateTgUser(ctx context.Context, TgUser *dto.TgUser) error {
	if err := a.TgUserDao.Create(TgUser); err != nil {
		return err
	}
	_, err := a.gameService.MissionCheckIn(ctx, &proto.UserIdReq{Id: TgUser.UUID})
	if err != nil {
		a.Logger.Errorf("MissionCheckIn err:%s", err.Error())
	}
	return nil
}

func (a *TgAccount) FindTgUserIdByWallet(ctx context.Context, wallet string) (string, error) {
	TgUser, err := a.TgUserDao.FindByAddress(wallet)
	if err != nil {
		return "", err
	}
	if nil == TgUser {
		return "", derrors.ErrUserNotExist
	}
	return TgUser.UUID, nil
}
