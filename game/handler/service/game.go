package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Mu-Exchange/Mu-End/common"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/redis"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/basic"
	"github.com/Mu-Exchange/Mu-End/common/utils/uuid"
	"github.com/Mu-Exchange/Mu-End/game/handler/cache"
	"github.com/Mu-Exchange/Mu-End/game/handler/dao"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type GameService interface {
	DoMission(ctx context.Context, userId, address string, score int64) error
	Cron(ctx context.Context) error
	GameRandom(ctx context.Context, userId string, level int) (int, int, int, error)
	LeaderBoard(ctx context.Context, limit uint64) ([]*proto.LeaderBoardResp_LeaderBoardItem, error)
	CheckDailyMission(ctx context.Context, userId string, missionType uint64) (bool, error)

	AddScoreHistory(history *dto.UserScoreHistory) error
	QueryTotalTypeScore(ctx context.Context, userId string) (map[uint64]uint64, error)
	QueryTotalScore(ctx context.Context, userId string) (uint64, error)
	QueryTotalScoreByType(ctx context.Context, userId string, scoreType uint64) (uint64, error)
	QueryScoreList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreHistory, error)
	QueryScoreListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreHistory, error)
	QueryScoreHistory(tx context.Context, id string) (*dto.UserScoreHistory, error)

	AddScoreUsedHistory(history *dto.UserScoreUsedHistory) error
	QueryTotalTypeScoreUsed(ctx context.Context, userId string) (map[uint64]uint64, error)
	QueryTotalScoreUsed(ctx context.Context, userId string) (uint64, error)
	QueryTotalScoreUsedByType(ctx context.Context, userId string, scoreUsedType uint64) (uint64, error)
	QueryScoreUsedList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreUsedHistory, error)
	QueryScoreUsedListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreUsedHistory, error)
	QueryScoreUsedHistory(tx context.Context, id string) (*dto.UserScoreUsedHistory, error)

	QueryTotalCash(userId string) (uint64, error)
}

type Game struct {
	BaseService
	repo.CommonComponents
	gameCache    cache.GameCache
	userService  proto.UserService
	scoreUsedDao dao.ScoreUsedDao
	scoreDao     dao.ScoreDao
	cashDao      dao.CashDao
	logger       *logrus.Logger
	lock         *redis.DistributedLock
}

func (ref *Game) Start(sd fx.Shutdowner) error {
	go func() {
		var wg = 0
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				now := time.Now()
				ref.Logger.Infof("check self: %s", now.Format("2006-01-02 15:04:05"))
				if wg == 1 {
					continue
				}
				nextHour := now.Truncate(time.Minute).Add(time.Minute)
				time.AfterFunc(nextHour.Sub(now), func() {
					ref.Logger.Infof("do self:%s", time.Now().Format("2006-01-02 15:04:05"))
					if err := ref.Cron(context.Background()); err != nil {
						ref.Logger.Errorf("do self:%s, err:%s", time.Now().Format("2006-01-02 15:04:05"), err.Error())
					}
					wg = 0
				})
				ref.Logger.Infof("add mission: %s", nextHour.Format("2006-01-02 15:04:05"))
				wg = 1
			}
		}
	}()
	return nil
}

func (ref *Game) Stop(sd fx.Shutdowner) error {
	return nil
}

func (a *Game) Cron(ctx context.Context) error {
	userStakes, err := a.gameCache.GetUserStake()
	if err != nil {
		return err
	}
	a.Logger.Infof("start to send dice for users length:[%d]", len(userStakes))
	for _, userId := range userStakes {
		user, err := a.userService.Profile(ctx, &proto.UserMutualReq{
			TargetUserId: userId,
		})
		if err != nil {
			return err
		}
		if err := a.DoMission(ctx, userId, "", int64(user.UserInfo.DiceSpeed)); err != nil {
			a.Logger.Errorf("do mission for user [%s] err [%s]", userId, err.Error())
		}
	}
	return nil
}

func (a *Game) DoMission(ctx context.Context, userId, address string, score int64) error {
	if len(userId) == 0 {
		if len(address) == 0 {
			a.Logger.Warnf("cannot found address [%s] user", address)
			return nil
		}
		user, err := a.userService.Profile(ctx, &proto.UserMutualReq{
			Wallet: address,
		})
		if err != nil {
			return err
		}
		if nil == user {
			a.Logger.Warnf("cannot found address [%s] user ", address)
			return nil
		}
		userId = user.Id
	}

	if err := a.AddScoreHistory(&dto.UserScoreHistory{
		UUID:   uuid.GenUUID().Encode(),
		UserID: userId,
		Score:  score,
		Type:   0,
		Desc:   "",
		DescId: "",
	}); err != nil {
		return err
	}
	return nil
}

func (ref *Game) CheckDailyMission(ctx context.Context, userId string, missionType uint64) (bool, error) {
	scoreList, err := ref.QueryScoreListByType(ctx, userId, 1, 1, missionType)
	if err != nil {
		return false, err
	}
	if len(scoreList) > 0 {
		date1 := scoreList[0].CreatedAt.Truncate(24 * time.Hour)
		date2 := time.Now().Truncate(24 * time.Hour)
		if date1.Equal(date2) {
			return false, nil
		}
	}
	return true, nil
}

const DistributionLockExpire = 5 * time.Second

func gameLockScoreKey(userId string) string {
	return fmt.Sprintf("{LOCK}-GAME-LOCK-KEY-%s", userId)
}

func (a *Game) AddCashHistory(history *dto.UserCashHistory) error {
	if err := a.cashDao.Create(history); err != nil {
		return err
	}
	cashTotal, err := a.cashDao.QueryTotalCash(history.UserID)
	if err != nil {
		return err
	}
	if err := a.gameCache.AddCashHistory(history); err != nil {
		return err
	}
	if err := a.gameCache.SetTotalCash(history.UserID, cashTotal); err != nil {
		return err
	}
	if err := a.gameCache.SetLeaderBoardUser(history.UserID, cashTotal); err != nil {
		return err
	}
	return nil
}

func (a *Game) QueryTotalCash(userId string) (uint64, error) {
	cash, err := a.gameCache.GetTotalCash(userId)
	if err != nil {
		return 0, err
	}
	if cash != 0 {
		return cash, nil
	}
	cash, err = a.cashDao.QueryTotalCash(userId)
	if err != nil {
		return 0, err
	}
	return cash, a.gameCache.SetTotalScoreUsed(userId, cash)
}

func random(level int, dual bool) (int, int, int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	dice1 := rand.Intn(6) + 1
	if !dual {
		return dice1, 0, dice1 * level
	}
	dice2 := rand.Intn(6) + 1
	base := (dice1 + dice2) * level
	score := 0
	if dice1 != dice2 {
		score = base
	} else {
		switch dice1 {
		case 1:
			score = base * 2
		case 2:
			score = base * 4
		case 3:
			score = base * 6
		case 4:
			score = base * 8
		case 5:
			score = base * 10
		case 6:
			score = base * 12
		default:
			break
		}
	}
	return dice1, dice2, score
}

func (ref *Game) GameRandom(ctx context.Context, userId string, level int) (int, int, int, error) {
	if !ref.lock.AcquireLock(gameLockScoreKey(userId), DistributionLockExpire) {
		return 0, 0, 0, derrors.ErrUserPostTooFrequently
	}
	defer ref.lock.ReleaseLock(gameLockScoreKey(userId))
	amount := 1 * common.ScoreAccuracy * level
	// 1. getTotalScore
	totalScore, err := ref.QueryTotalScore(ctx, userId)
	if err != nil {
		return 0, 0, 0, err
	}
	totalScoreUsed, err := ref.QueryTotalScoreUsed(ctx, userId)
	if err != nil {
		return 0, 0, 0, err
	}
	if totalScore < totalScoreUsed+uint64(amount) {
		return 0, 0, 0, fmt.Errorf("not enough score")
	}
	profile, err := ref.userService.Profile(ctx, &proto.UserMutualReq{TargetUserId: userId})
	if err != nil {
		return 0, 0, 0, err
	}
	dice1, dice2, cash := random(level, profile.UserInfo.MintDice)
	if err := ref.TransactionExecute(func(db *gorm.DB) error {
		scoreUsedHis := &dto.UserScoreUsedHistory{
			UUID:   uuid.GenUUID().Encode(),
			UserID: userId,
			Score:  int64(amount),
			Type:   common.ConsumeRandom.ScoreType,
			Desc:   common.ConsumeRandom.Desc,
		}
		if err := ref.AddScoreUsedHistory(scoreUsedHis); err != nil {
			return err
		}
		cashHistory := &dto.UserCashHistory{
			UUID:                   uuid.GenUUID().Encode(),
			UserScoreUsedHistoryID: scoreUsedHis.UUID,
			UserID:                 userId,
			Cash:                   int64(cash),
		}
		if err := ref.AddCashHistory(cashHistory); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, 0, 0, err
	}
	return dice1, dice2, cash, nil
}

func (ref *Game) LeaderBoard(ctx context.Context, limit uint64) ([]*proto.LeaderBoardResp_LeaderBoardItem, error) {
	if limit == 0 {
		limit = 10
	}
	_, userIds, err := ref.gameCache.TopIndexLB(1, int64(limit))
	if err != nil {
		return nil, err
	}
	var items []*proto.LeaderBoardResp_LeaderBoardItem
	for _, userId := range userIds {
		user, err := ref.userService.Profile(ctx, &proto.UserMutualReq{
			TargetUserId: userId,
		})
		if err != nil {
			return nil, err
		}
		cash, err := ref.gameCache.GetLeaderBoardUser(user.Id)
		if err != nil {
			return nil, err
		}
		items = append(items, &proto.LeaderBoardResp_LeaderBoardItem{
			UserId:     user.Id,
			Cash:       cash,
			ProfileUrl: user.UserInfo.ProfileUrl,
		})
	}
	return items, nil
}

func NewGameService(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents,
	service BaseService,
	userService proto.UserService,
	scoreUsedDao dao.ScoreUsedDao,
	scoreDao dao.ScoreDao,
	cashDao dao.CashDao,
	gameCache cache.GameCache) (GameService, error) {
	lock, err := redis.NewDistributedLock(cfg.Cfg.Redis)
	if err != nil {
		return nil, err
	}
	erc := &Game{
		BaseService:      service,
		CommonComponents: cfg,
		userService:      userService,
		scoreUsedDao:     scoreUsedDao,
		scoreDao:         scoreDao,
		gameCache:        gameCache,
		cashDao:          cashDao,
		logger:           cfg.Logger,
		lock:             lock,
	}
	basic.RegisterHook(lc, sd, erc)
	return erc, nil
}
