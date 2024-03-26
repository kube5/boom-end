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

type TGGameService interface {
	DoMission(ctx context.Context, userId string, mission common.Mission) error
	Cron(ctx context.Context) error
	TGGameRandom(ctx context.Context, userId string, level int) (int, int, int, error)
	LeaderBoard(ctx context.Context, userId string, limit uint64) ([]*proto.LeaderBoardResp_LeaderBoardItem, *proto.LeaderBoardResp_LeaderBoardItem, error)
	Game24H(ctx context.Context, userId string) (uint64, error)

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

type TGGame struct {
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

func (ref *TGGame) Game24H(ctx context.Context, userId string) (uint64, error) {
	return ref.gameCache.TGGetTotalScoreUsed24H(userId)
}

func (ref *TGGame) Start(sd fx.Shutdowner) error {
	go func() {
		var wg = 0
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				now := time.Now().UTC()
				ref.Logger.Infof("check self 1h: %s", now.Format("2006-01-02 15:04:05"))
				if wg == 1 {
					continue
				}
				nextHour := now.Truncate(time.Hour).Add(time.Hour)
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

	go func() {
		var wg = 0
		ticker := time.NewTicker(time.Hour * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				now := time.Now().UTC()
				ref.Logger.Infof("check self 24h: %s", now.Format("2006-01-02 15:04:05"))
				if wg == 1 {
					continue
				}
				nextDayUp := now.AddDate(0, 0, 1)
				nextDay := time.Date(nextDayUp.Year(), nextDayUp.Month(), nextDayUp.Day(), 0, 0, 0, 0, time.UTC)
				time.AfterFunc(nextDay.Sub(now), func() {
					ref.Logger.Infof("do self:%s", time.Now().Format("2006-01-02 15:04:05"))
					if err := ref.gameCache.TGDelTotalScoreUsed24H(); err != nil {
						ref.Logger.Errorf("TGDelTotalScoreUsed24H:%s, err:%s", time.Now().Format("2006-01-02 15:04:05"), err.Error())
					}
					wg = 0
				})
				ref.Logger.Infof("TGDelTotalScoreUsed24H: %s", nextDay.Format("2006-01-02 15:04:05"))
				wg = 1
			}
		}
	}()
	return nil
}

func (ref *TGGame) Stop(sd fx.Shutdowner) error {
	return nil
}

func (a *TGGame) Cron(ctx context.Context) error {
	userStakes, err := a.gameCache.GetTGUser()
	if err != nil {
		return err
	}
	for _, wallet := range userStakes {
		user, err := a.userService.ProfileTG(ctx, &proto.UserMutualReq{
			Wallet: wallet,
		})
		if err != nil {
			return err
		}
		mission := common.MissionCron
		mission.ScoreAmount = user.UserInfo.DiceSpeed
		if err := a.DoMission(ctx, user.Id, mission); err != nil {
			a.Logger.Errorf("do mission for user [%s] err [%s]", wallet, err.Error())
		}
	}
	return nil
}

func (a *TGGame) DoMission(ctx context.Context, userId string, mission common.Mission) error {
	if len(userId) == 0 {
		a.Logger.Warnf("cannot found userId [%s] user", userId)
		return nil
	}
	if err := a.AddScoreHistory(&dto.UserScoreHistory{
		UUID:   uuid.GenUUID().Encode(),
		UserID: userId,
		Score:  int64(mission.ScoreAmount),
		Type:   mission.ScoreType,
		Desc:   mission.Desc,
		DescId: mission.DescId,
	}); err != nil {
		return err
	}
	return nil
}

const tgDistributionLockExpire = 5 * time.Second

func tggameLockScoreKey(userId string) string {
	return fmt.Sprintf("{LOCK}-TG-GAME-LOCK-KEY-%s", userId)
}

func (a *TGGame) AddCashHistory(history *dto.UserCashHistory) error {
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
	if err := a.gameCache.TGSetLeaderBoardUser(history.UserID, cashTotal); err != nil {
		return err
	}
	return nil
}

func (a *TGGame) QueryTotalCash(userId string) (uint64, error) {
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

func tgrandom(level int, dual bool) (int, int, int) {
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

func (ref *TGGame) TGGameRandom(ctx context.Context, userId string, level int) (int, int, int, error) {
	if !ref.lock.AcquireLock(tggameLockScoreKey(userId), tgDistributionLockExpire) {
		return 0, 0, 0, derrors.ErrUserPostTooFrequently
	}
	defer ref.lock.ReleaseLock(tggameLockScoreKey(userId))
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
		return 0, 0, 0, derrors.ErrNotEnoughDice
	}
	profile, err := ref.userService.ProfileTG(ctx, &proto.UserMutualReq{TargetUserId: userId})
	if err != nil {
		return 0, 0, 0, err
	}
	dice1, dice2, cash := tgrandom(level, profile.UserInfo.MintDice)
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
			Dice1:                  int64(dice1),
			Dice2:                  int64(dice2),
			Level:                  int64(level),
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

func (ref *TGGame) LeaderBoard(ctx context.Context, userId string, limit uint64) ([]*proto.LeaderBoardResp_LeaderBoardItem, *proto.LeaderBoardResp_LeaderBoardItem, error) {
	if limit == 0 {
		limit = 10
	}
	_, userIds, err := ref.gameCache.TGTopIndexLB(1, int64(limit))
	if err != nil {
		return nil, nil, err
	}
	var items []*proto.LeaderBoardResp_LeaderBoardItem
	for _, userId := range userIds {
		user, err := ref.userService.ProfileTG(ctx, &proto.UserMutualReq{
			TargetUserId: userId,
		})
		if err != nil {
			return nil, nil, err
		}
		cash, err := ref.gameCache.TGGetLeaderBoardUser(user.Id)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, &proto.LeaderBoardResp_LeaderBoardItem{
			Wallet: user.UserInfo.Wallet,
			Cash:   cash,
		})
	}
	if len(userId) > 0 {
		user, err := ref.userService.ProfileTG(ctx, &proto.UserMutualReq{
			TargetUserId: userId,
		})
		if err != nil {
			return nil, nil, err
		}
		rank, err := ref.gameCache.TGGetLeaderBoardUserScore(user.Id)
		if err != nil {
			return nil, nil, err
		}
		if -1 == rank {
			return items, nil, nil
		}
		rank++
		cash, err := ref.gameCache.TGGetLeaderBoardUser(user.Id)
		if err != nil {
			return nil, nil, err
		}
		return items, &proto.LeaderBoardResp_LeaderBoardItem{
			Wallet: user.UserInfo.Wallet,
			Cash:   cash,
			Rank:   uint64(rank),
		}, nil
	} else {
		return items, nil, nil
	}
}

func NewTGGameService(lc fx.Lifecycle, sd fx.Shutdowner, cfg repo.CommonComponents,
	service BaseService,
	userService proto.UserService,
	scoreUsedDao dao.ScoreUsedDao,
	scoreDao dao.ScoreDao,
	cashDao dao.CashDao,
	gameCache cache.GameCache) (TGGameService, error) {
	lock, err := redis.NewDistributedLock(cfg.Cfg.Redis)
	if err != nil {
		return nil, err
	}
	erc := &TGGame{
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

func (a *TGGame) AddScoreHistory(history *dto.UserScoreHistory) error {
	if err := a.scoreDao.Create(history); err != nil {
		return err
	}
	scoreTotal, err := a.scoreDao.QueryTotalScore(history.UserID)
	if err != nil {
		return err
	}
	totalScoreByType, err := a.scoreDao.QueryTotalScoreByType(history.UserID, history.Type)
	if err != nil {
		return err
	}

	if err := a.gameCache.AddScoreHistory(history); err != nil {
		return err
	}
	if err := a.gameCache.SetTotalScore(history.UserID, scoreTotal); err != nil {
		return err
	}
	if err := a.gameCache.SetTotalScoreByType(history.UserID, totalScoreByType, history.Type); err != nil {
		return err
	}
	return nil
}

func (a *TGGame) QueryTotalTypeScore(ctx context.Context, userId string) (map[uint64]uint64, error) {
	totalTypeScore, err := a.gameCache.GetTotalTypeScore(userId)
	if err != nil {
		return nil, err
	}
	return totalTypeScore, err
}

func (a *TGGame) QueryTotalScore(ctx context.Context, userId string) (uint64, error) {
	totalScore, err := a.gameCache.GetTotalScore(userId)
	if err != nil {
		return 0, err
	}
	if totalScore != 0 {
		return totalScore, nil
	}
	score, err := a.scoreDao.QueryTotalScore(userId)
	if err != nil {
		return 0, err
	}
	return score, a.gameCache.SetTotalScore(userId, score)
}

func (a *TGGame) QueryTotalScoreByType(ctx context.Context, userId string, scoreType uint64) (uint64, error) {
	totalScore, err := a.gameCache.GetTotalScoreByType(userId, scoreType)
	if err != nil {
		return 0, err
	}
	if totalScore != 0 {
		return totalScore, nil
	}

	score, err := a.scoreDao.QueryTotalScoreByType(userId, scoreType)
	if err != nil {
		return 0, err
	}
	return score, a.gameCache.SetTotalScoreByType(userId, score, scoreType)
}

func (a *TGGame) QueryScoreList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreHistory, error) {
	score, err := a.gameCache.QueryScoreList(userId, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	if len(score) == 0 {
		scoreList, err := a.scoreDao.QueryScoreHistoryByUserId(userId)
		if err != nil {
			return nil, err
		}
		for _, sc := range scoreList {
			if err := a.gameCache.AddScoreHistory(sc); err != nil {
				return nil, err
			}
		}
		if len(scoreList) != 0 {
			score, err = a.gameCache.QueryScoreList(userId, pageNo, pageSize)
			if err != nil {
				return nil, err
			}
		}
	}
	shs := make([]*dto.UserScoreHistory, 0, len(score))
	for _, id := range score {
		sh, err := a.QueryScoreHistory(ctx, id)
		if err != nil {
			return nil, err
		}
		shs = append(shs, sh)
	}
	return shs, err
}

func (a *TGGame) QueryScoreListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreHistory, error) {
	score, err := a.gameCache.QueryScoreListByType(userId, pageNo, pageSize, int64(scoreType))
	if err != nil {
		return nil, err
	}
	if len(score) == 0 {
		scoreList, err := a.scoreDao.QueryScoreHistoryByUserIdAndType(userId, int64(scoreType))
		if err != nil {
			return nil, err
		}
		for _, sc := range scoreList {
			if err := a.gameCache.AddScoreHistory(sc); err != nil {
				return nil, err
			}
		}
		if len(scoreList) != 0 {
			score, err = a.gameCache.QueryScoreListByType(userId, pageNo, pageSize, int64(scoreType))
			if err != nil {
				return nil, err
			}
		}
	}
	shs := make([]*dto.UserScoreHistory, 0, len(score))
	for _, id := range score {
		sh, err := a.QueryScoreHistory(ctx, id)
		if err != nil {
			return nil, err
		}
		shs = append(shs, sh)
	}
	return shs, err
}

func (a *TGGame) QueryScoreHistory(tx context.Context, id string) (*dto.UserScoreHistory, error) {
	scoreDetail, err := a.gameCache.GetScoreHistory(id)
	if err != nil {
		return nil, err
	}
	if scoreDetail != nil {
		return scoreDetail, nil
	}
	scoreDetail, err = a.scoreDao.QueryScoreHistory(id)
	if err != nil {
		return nil, err
	}
	if scoreDetail == nil {
		return nil, derrors.ErrWrongScoreHistoryId
	}
	return scoreDetail, a.gameCache.AddScoreHistory(scoreDetail)
}

func (a *TGGame) AddScoreUsedHistory(history *dto.UserScoreUsedHistory) error {
	if err := a.scoreUsedDao.Create(history); err != nil {
		return err
	}
	scoreUsedTotal, err := a.scoreUsedDao.QueryTotalScoreUsed(history.UserID)
	if err != nil {
		return err
	}
	totalScoreUsedByType, err := a.scoreUsedDao.QueryTotalScoreUsedByType(history.UserID, history.Type)
	if err != nil {
		return err
	}
	if err := a.gameCache.AddScoreUsedHistory(history); err != nil {
		return err
	}
	if err := a.gameCache.SetTotalScoreUsed(history.UserID, scoreUsedTotal); err != nil {
		return err
	}
	if err := a.gameCache.SetTotalScoreUsedByType(history.UserID, totalScoreUsedByType, history.Type); err != nil {
		return err
	}
	score, err := a.gameCache.TGSetTotalScoreUsed24H(history.UserID, history.Score)
	if err != nil {
		return err
	}
	if score >= 2400 {
		_, err := a.userService.AddDiceSpeed(context.Background(), &proto.AddDiceSpeedReq{
			Id:   history.UserID,
			Dice: 10,
			Desc: "24H",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *TGGame) QueryTotalTypeScoreUsed(ctx context.Context, userId string) (map[uint64]uint64, error) {
	totalTypeScoreUsed, err := a.gameCache.GetTotalTypeScoreUsed(userId)
	if err != nil {
		return nil, err
	}
	return totalTypeScoreUsed, err
}

func (a *TGGame) QueryTotalScoreUsed(ctx context.Context, userId string) (uint64, error) {
	totalScoreUsed, err := a.gameCache.GetTotalScoreUsed(userId)
	if err != nil {
		return 0, err
	}
	if totalScoreUsed != 0 {
		return totalScoreUsed, nil
	}
	score, err := a.scoreUsedDao.QueryTotalScoreUsed(userId)
	if err != nil {
		return 0, err
	}
	return score, a.gameCache.SetTotalScoreUsed(userId, score)
}

func (a *TGGame) QueryTotalScoreUsedByType(ctx context.Context, userId string, scoreType uint64) (uint64, error) {
	totalScoreUsed, err := a.gameCache.GetTotalScoreUsedByType(userId, scoreType)
	if err != nil {
		return 0, err
	}
	if totalScoreUsed != 0 {
		return totalScoreUsed, nil
	}

	score, err := a.scoreUsedDao.QueryTotalScoreUsedByType(userId, scoreType)
	if err != nil {
		return 0, err
	}
	return score, a.gameCache.SetTotalScoreUsedByType(userId, score, scoreType)
}

func (a *TGGame) QueryScoreUsedList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreUsedHistory, error) {
	score, err := a.gameCache.QueryScoreUsedList(userId, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	if len(score) == 0 {
		scoreList, err := a.scoreUsedDao.QueryScoreUsedHistoryByUserId(userId)
		if err != nil {
			return nil, err
		}
		for _, sc := range scoreList {
			if err := a.gameCache.AddScoreUsedHistory(sc); err != nil {
				return nil, err
			}
		}
		if len(scoreList) != 0 {
			score, err = a.gameCache.QueryScoreUsedList(userId, pageNo, pageSize)
			if err != nil {
				return nil, err
			}
		}
	}
	shs := make([]*dto.UserScoreUsedHistory, 0, len(score))
	for _, id := range score {
		sh, err := a.QueryScoreUsedHistory(ctx, id)
		if err != nil {
			return nil, err
		}
		shs = append(shs, sh)
	}
	return shs, err
}

func (a *TGGame) QueryScoreUsedListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreUsedHistory, error) {
	score, err := a.gameCache.QueryScoreUsedListByType(userId, pageNo, pageSize, int64(scoreType))
	if err != nil {
		return nil, err
	}
	if len(score) == 0 {
		scoreList, err := a.scoreUsedDao.QueryScoreUsedHistoryByUserIdAndType(userId, int64(scoreType))
		if err != nil {
			return nil, err
		}
		for _, sc := range scoreList {
			if err := a.gameCache.AddScoreUsedHistory(sc); err != nil {
				return nil, err
			}
		}
		if len(scoreList) != 0 {
			score, err = a.gameCache.QueryScoreUsedListByType(userId, pageNo, pageSize, int64(scoreType))
			if err != nil {
				return nil, err
			}
		}
	}
	shs := make([]*dto.UserScoreUsedHistory, 0, len(score))
	for _, id := range score {
		sh, err := a.QueryScoreUsedHistory(ctx, id)
		if err != nil {
			return nil, err
		}
		shs = append(shs, sh)
	}
	return shs, err
}

func (a *TGGame) QueryScoreUsedHistory(tx context.Context, id string) (*dto.UserScoreUsedHistory, error) {
	scoreDetail, err := a.gameCache.GetScoreUsedHistory(id)
	if err != nil {
		return nil, err
	}
	if scoreDetail != nil {
		return scoreDetail, nil
	}
	scoreDetail, err = a.scoreUsedDao.QueryScoreUsedHistory(id)
	if err != nil {
		return nil, err
	}
	if scoreDetail == nil {
		return nil, derrors.ErrWrongScoreHistoryId
	}
	return scoreDetail, a.gameCache.AddScoreUsedHistory(scoreDetail)
}
