package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/redis"
	redis_client "github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var ctx = context.TODO()

type GameCache interface {
	QueryCashList(userId string, pageNo, pageSize int64) ([]string, error)
	AddCashHistory(history *dto.UserCashHistory) error
	GetCashHistory(id string) (*dto.UserCashHistory, error)
	SetTotalCash(userId string, count uint64) error
	GetTotalCash(userId string) (uint64, error)

	// score
	QueryScoreList(userId string, pageNo, pageSize int64) ([]string, error)
	AddScoreHistory(history *dto.UserScoreHistory) error
	GetScoreHistory(id string) (*dto.UserScoreHistory, error)
	SetTotalScore(userId string, count uint64) error
	GetTotalScore(userId string) (uint64, error)

	QueryScoreListByType(userId string, pageNo, pageSize, scoreType int64) ([]string, error)
	GetTotalScoreByType(userId string, scoreType uint64) (uint64, error)
	GetTotalTypeScore(userId string) (map[uint64]uint64, error)
	SetTotalScoreByType(userId string, count uint64, scoreType uint64) error

	// scoreUsed
	QueryScoreUsedList(userId string, pageNo, pageSize int64) ([]string, error)
	AddScoreUsedHistory(history *dto.UserScoreUsedHistory) error
	GetScoreUsedHistory(id string) (*dto.UserScoreUsedHistory, error)
	SetTotalScoreUsed(userId string, count uint64) error
	GetTotalScoreUsed(userId string) (uint64, error)

	QueryScoreUsedListByType(userId string, pageNo, pageSize, scoreUsedType int64) ([]string, error)
	GetTotalScoreUsedByType(userId string, scoreUsedType uint64) (uint64, error)
	GetTotalTypeScoreUsed(userId string) (map[uint64]uint64, error)
	SetTotalScoreUsedByType(userId string, count uint64, scoreUsedType uint64) error

	TopIndexLB(pageNo, pageSize int64) (int64, []string, error)
	SetLeaderBoardUser(userId string, score uint64) error
	GetLeaderBoardUser(userId string) (uint64, error)
	GetLeaderBoardUserScore(userId string) (int64, error)

	SetUserStake(wallet string) error
	GetUserStake() ([]string, error)

	TGTopIndexLB(pageNo, pageSize int64) (int64, []string, error)
	TGSetLeaderBoardUser(userId string, score uint64) error
	TGGetLeaderBoardUser(userId string) (uint64, error)
	TGGetLeaderBoardUserScore(userId string) (int64, error)
	GetTGUser() ([]string, error)

	TGSetTotalScoreUsed24H(userId string, count int64) (uint64, error)
	TGGetTotalScoreUsed24H(userId string) (uint64, error)
	TGDelTotalScoreUsed24H() error
}

type Game struct {
	*redis.Cache
}

func (u *Game) TGGetTotalScoreUsed24H(userId string) (uint64, error) {
	return u.HGet(ctx, fmt.Sprintf("%s-TG-SCORE-USED-24H", ServicePrefix), userId).Uint64()
}

func (u *Game) TGSetTotalScoreUsed24H(userId string, count int64) (uint64, error) {
	return u.HIncrBy(ctx, fmt.Sprintf("%s-TG-SCORE-USED-24H", ServicePrefix), userId, count).Uint64()
}

func (u *Game) TGDelTotalScoreUsed24H() error {
	return u.Del(ctx, fmt.Sprintf("%s-TG-SCORE-USED-24H", ServicePrefix)).Err()
}

const ServicePrefix = "{SGAME}"

func (u *Game) Start(sd fx.Shutdowner) error {
	return nil
}

func (u *Game) Stop(sd fx.Shutdowner) error {
	return nil
}

func NewGameCache(cache *redis.Cache) (GameCache, error) {
	return &Game{cache}, nil
}

func cashTotalKey(userId string) string {
	return fmt.Sprintf("%s-CASH-TOTAL-%s", ServicePrefix, userId)
}

func cashKey(userId string) string {
	return fmt.Sprintf("%s-CASH-%s", ServicePrefix, userId)
}

func cashDetailKey(id string) string {
	return fmt.Sprintf("%s-CASH-DETAIL-%s", ServicePrefix, id)
}

func (u *Game) QueryCashList(userId string, pageNo, pageSize int64) ([]string, error) {
	nums, err := u.ZCard(ctx, cashKey(userId)).Result()
	if err != nil {
		return nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	return u.ZRevRange(ctx, cashKey(userId), start, end).Result()
}

func (u *Game) AddCashHistory(f *dto.UserCashHistory) error {
	history, err := u.GetCashHistory(f.UUID)
	if err != nil {
		return err
	}
	if history != nil {
		return nil
	}
	data, err := json.Marshal(f)
	if err != nil {
		return err
	}
	pipe := u.Pipeline()
	pipe.Set(ctx, cashDetailKey(f.UUID), data, 0)
	pipe.ZAdd(ctx, cashKey(f.UserID), redis_client.Z{Member: f.UUID, Score: float64(f.CreatedAt.UnixMilli())})
	_, err = pipe.Exec(ctx)
	return err
}

func (u *Game) GetCashHistory(id string) (*dto.UserCashHistory, error) {
	scoreHistory := &dto.UserCashHistory{}
	val, err := u.Get(ctx, cashDetailKey(id)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), scoreHistory)
	return scoreHistory, err
}

func (u *Game) SetTotalCash(userId string, count uint64) error {
	return u.Set(ctx, cashTotalKey(userId), count, 0).Err()
}

func (u *Game) GetTotalCash(userId string) (uint64, error) {
	result, err := u.Get(ctx, cashTotalKey(userId)).Result()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(result, 10, 64)
}

func topIndexLB() string {
	return fmt.Sprintf("%s-TOP-INDEX-LB", ServicePrefix)
}
func userStakes() string {
	return fmt.Sprintf("%s-USER-STAKE", ServicePrefix)
}
func leaderBoardUserKey() string {
	return fmt.Sprintf("%s-USER-LB", ServicePrefix)
}
func HIdKey(id string) string {
	return fmt.Sprintf("%s-%s", ServicePrefix, id)
}

func (u *Game) TopIndexLB(pageNo, pageSize int64) (int64, []string, error) {
	nums, err := u.ZCard(ctx, topIndexLB()).Result()
	if err != nil {
		return 0, nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	res, err := u.ZRevRange(ctx, topIndexLB(), start, end).Result()
	return nums, res, err
}

func (u *Game) SetLeaderBoardUser(userId string, score uint64) error {
	pipe := u.Pipeline()
	pipe.ZAdd(ctx, topIndexLB(), redis_client.Z{Member: userId, Score: float64(score)})
	pipe.HSet(ctx, leaderBoardUserKey(), HIdKey(userId), score)
	_, err := pipe.Exec(ctx)
	return err
}

func (u *Game) GetLeaderBoardUser(userId string) (uint64, error) {
	val, err := u.HGet(ctx, leaderBoardUserKey(), HIdKey(userId)).Uint64()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return val, err
}

func (u *Game) GetLeaderBoardUserScore(userId string) (int64, error) {
	val, err := u.ZRevRank(ctx, topIndexLB(), userId).Uint64()
	if err == redis_client.Nil {
		return -1, nil
	}
	if err != nil {
		return 0, err
	}
	return int64(val), err
}

func (u *Game) SetUserStake(wallet string) error {
	if err := u.ZAdd(ctx, userStakes(), redis_client.Z{Member: wallet, Score: float64(0)}).Err(); err != nil {
		return err
	}
	return nil
}

func (u *Game) GetUserStake() ([]string, error) {
	return u.ZRevRange(ctx, userStakes(), 0, -1).Result()
}

//  tg-----------------

func tgUsers() string {
	return fmt.Sprintf("{TGSUSER}-USER-TG")
}
func tgtopIndexLB() string {
	return fmt.Sprintf("%s-TOP-TG-INDEX-LB", ServicePrefix)
}
func tgleaderBoardUserKey() string {
	return fmt.Sprintf("%s-USER-TG-LB", ServicePrefix)
}
func (u *Game) GetTGUser() ([]string, error) {
	return u.ZRevRange(ctx, tgUsers(), 0, -1).Result()
}

func (u *Game) TGTopIndexLB(pageNo, pageSize int64) (int64, []string, error) {
	nums, err := u.ZCard(ctx, tgtopIndexLB()).Result()
	if err != nil {
		return 0, nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	res, err := u.ZRevRange(ctx, tgtopIndexLB(), start, end).Result()
	return nums, res, err
}

func (u *Game) TGSetLeaderBoardUser(userId string, score uint64) error {
	pipe := u.Pipeline()
	pipe.ZAdd(ctx, tgtopIndexLB(), redis_client.Z{Member: userId, Score: float64(score)})
	pipe.HSet(ctx, tgleaderBoardUserKey(), HIdKey(userId), score)
	_, err := pipe.Exec(ctx)
	return err
}

func (u *Game) TGGetLeaderBoardUser(userId string) (uint64, error) {
	val, err := u.HGet(ctx, tgleaderBoardUserKey(), HIdKey(userId)).Uint64()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return val, err
}

func (u *Game) TGGetLeaderBoardUserScore(userId string) (int64, error) {
	val, err := u.ZRevRank(ctx, tgtopIndexLB(), userId).Uint64()
	if err == redis_client.Nil {
		return -1, nil
	}
	if err != nil {
		return 0, err
	}
	return int64(val), err
}
