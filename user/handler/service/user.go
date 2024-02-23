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

type UserService interface {
	QueryUser(ctx context.Context, userId string) (*dto.User, error)
	QueryUserByWallet(ctx context.Context, wallet string) (*dto.User, error)
	LoginPreByMetaMask(ctx context.Context, publicAddress string) (string, error)
	LoginByMetaMask(ctx context.Context, publicAddress, signature string) (*model.TokenDetails, error)
	LoginInternal(ctx context.Context, address string) (*model.TokenDetails, error)
	UpdateUserByWallet(ctx context.Context, user *dto.User) error
	RefreshToken(ctx context.Context, refreshUuid, accessUuid string, userId string) (*model.TokenDetails, error)
	Logout(ctx context.Context, givenUuid string, userId string) error

	FindUserIdByWallet(ctx context.Context, wallet string) (string, error)
}

var _ UserService = (*Account)(nil)

const DistributionLockExpire = 5 * time.Second
const InviteCodeLockKey = "{LOCK}-InviteCodeLock"

type Account struct {
	BaseService
	repo.CommonComponents
	userDao     dao.UserDao
	userCache   cache.UserCache
	ethClient   *ethclient.Client
	gameService proto.GameService
	lock        *redis.DistributedLock
}

func (a *Account) QueryUser(ctx context.Context, userId string) (*dto.User, error) {
	user, err := a.userCache.GetUser(userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	user, err = a.userDao.FindByUserID(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, derrors.ErrUserNotExist
	}
	return user, a.userCache.SetUser(user)
}

func (a *Account) QueryUserByWallet(ctx context.Context, wallet string) (*dto.User, error) {
	if len(wallet) == 0 {
		return nil, derrors.ErrUserNotExist
	}
	user, err := a.userDao.FindByAddress(wallet)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthService(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents,
	service BaseService,
	gameService proto.GameService,
	userDao dao.UserDao,
	userCache cache.UserCache,
) (UserService, error) {
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
	ser := &Account{
		BaseService:      service,
		CommonComponents: cfg,
		userCache:        userCache,
		userDao:          userDao,
		ethClient:        ethClient,
		gameService:      gameService,
		lock:             lock,
	}

	basic.RegisterHook(lc, sd, ser)
	return ser, nil
}

func (a *Account) Start(sd fx.Shutdowner) error {
	return nil
}

func (a *Account) Stop(sd fx.Shutdowner) error {
	return nil
}

func (a *Account) LoginPreByMetaMask(ctx context.Context, publicAddress string) (string, error) {
	nonce := time.Now().GoString()
	signText := "Welcome to rollie!\n\n" +
		"Click to sign in and accept the rollie Terms of Service: https://rollie.finance/\n\n" +
		"This request will not trigger a blockchain transaction or cost any gas fees.\n\n" +
		"Your authentication status will reset after 24 hours.\n\n" +
		"Wallet address:\n" +
		publicAddress +
		"\n\nNonce:\n" +
		nonce
	if err := a.userCache.SetMetamaskLoginNonce(publicAddress, signText); err != nil {
		return "", err
	}
	return signText, nil
}

func (a *Account) LoginByMetaMask(ctx context.Context, publicAddress, signature string) (*model.TokenDetails, error) {
	if len(publicAddress) == 0 {
		return nil, derrors.ErrUserNotExist
	}
	nonce, err := a.userCache.GetMetamaskLoginNonce(publicAddress)
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
	user, err := a.userDao.FindByAddress(publicAddress)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user = &dto.User{
			UUID:   uuid.GenUUID().Encode(),
			Wallet: publicAddress,
		}
		if err := a.userDao.Create(user); err != nil {
			return nil, err
		}
		a.Logger.Infof("New user[%s] with address[%s] is registering", user.UUID, user.Wallet)
	}
	td, err := utils.CreateToken(user.UUID)
	if err != nil {
		return nil, err
	}
	if err := a.userCache.DeleteUserAuth(user.UUID); err != nil {
		return nil, err
	}
	if err := a.userCache.CreateAuth(user.UUID, td); err != nil {
		return nil, err
	}
	a.Logger.Infof("New user[%s] with address[%s] login in", user.UUID, user.Wallet)
	return td, a.userCache.SetUser(user)
}

func (a *Account) LoginInternal(ctx context.Context, publicAddress string) (*model.TokenDetails, error) {
	user, err := a.userDao.FindByAddress(publicAddress)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, derrors.ErrUserNotExist
	}
	td, err := utils.CreateToken(user.UUID)
	if err != nil {
		return nil, err
	}
	if err := a.userCache.DeleteUserAuth(user.UUID); err != nil {
		return nil, err
	}
	if err := a.userCache.CreateAuth(user.UUID, td); err != nil {
		return nil, err
	}
	a.Logger.Infof("!!internal login, user[%s] with address[%s] login in", user.UUID, user.Wallet)
	return td, a.userCache.SetUser(user)
}

func (a *Account) Logout(ctx context.Context, givenUuid, userId string) error {
	deleted, err := a.userCache.DeleteAuth(userId, givenUuid)
	if deleted == 0 || err != nil {
		return derrors.ErrAuthToken
	}
	return err
}

func (a *Account) RefreshToken(ctx context.Context, refreshUuid, accessUuid, userId string) (*model.TokenDetails, error) {
	deleted, err := a.userCache.DeleteAuth(userId, refreshUuid)
	if deleted == 0 || err != nil {
		return nil, derrors.ErrRefreshTokenExpired
	}

	deleted, err = a.userCache.DeleteAuth(userId, accessUuid)
	if deleted == 0 || err != nil {
		return nil, derrors.ErrAuthValidationExpired
	}
	td, err := utils.CreateToken(userId)
	if err != nil {
		return nil, err
	}

	if err := a.userCache.CreateAuth(userId, td); err != nil {
		return nil, err
	}

	return td, nil
}

func (a *Account) UpdateUserByWallet(ctx context.Context, user *dto.User) error {
	if len(user.Wallet) == 0 {
		return nil
	}

	oldUser, err := a.QueryUserByWallet(ctx, user.Wallet)
	if err != nil {
		return err
	}
	if nil == oldUser {
		newUser := &dto.User{
			UUID:              uuid.GenUUID().Encode(),
			Wallet:            user.Wallet,
			StakingAmountETH:  user.StakingAmountETH,
			StakingAmountUSDB: user.StakingAmountUSDB,
			MintDice:          user.MintDice,
			DiceSpeed:         user.DiceSpeed,
		}
		if err := a.userDao.Create(newUser); err != nil {
			return err
		}
		a.Logger.Infof("New user[%s] with address[%s] is registering", user.UUID, user.Wallet)
		return a.userCache.SetUser(newUser)
	} else {
		user.UUID = oldUser.UUID
		if oldUser.MintDice {
			user.MintDice = oldUser.MintDice
		}
		_, err := a.userDao.UpdateUser(user)
		if err != nil {
			return err
		}
		return a.userCache.SetUser(user)
	}
}

func (a *Account) FindUserIdByWallet(ctx context.Context, wallet string) (string, error) {
	user, err := a.userDao.FindByAddress(wallet)
	if err != nil {
		return "", err
	}
	if nil == user {
		return "", derrors.ErrUserNotExist
	}
	return user.UUID, nil
}
