package cache

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	redis_client "github.com/redis/go-redis/v9"
)

func scoreTotalKey(userId string) string {
	return fmt.Sprintf("%s-SCORE-TOTAL-%s", ServicePrefix, userId)
}
func scoreTypeTotalKey(userId string) string {
	return fmt.Sprintf("%s-SCORE-TYPE-TOTAL-%s", ServicePrefix, userId)
}
func scoreKey(userId string) string {
	return fmt.Sprintf("%s-SCORE-%s", ServicePrefix, userId)
}
func scoreTypeKey(userId string, scoreType int64) string {
	return fmt.Sprintf("%s-SCORE-%s-%d", ServicePrefix, userId, scoreType)
}
func scoreDetailKey(id string) string {
	return fmt.Sprintf("%s-SCORE-DETAIL-%s", ServicePrefix, id)
}

func (u *Game) QueryScoreList(userId string, pageNo, pageSize int64) ([]string, error) {
	nums, err := u.ZCard(ctx, scoreKey(userId)).Result()
	if err != nil {
		return nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	return u.ZRevRange(ctx, scoreKey(userId), start, end).Result()
}

func (u *Game) QueryScoreListByType(userId string, pageNo, pageSize, scoreType int64) ([]string, error) {
	nums, err := u.ZCard(ctx, scoreTypeKey(userId, scoreType)).Result()
	if err != nil {
		return nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	return u.ZRevRange(ctx, scoreTypeKey(userId, scoreType), start, end).Result()
}

func (u *Game) AddScoreHistory(f *dto.UserScoreHistory) error {
	history, err := u.GetScoreHistory(f.UUID)
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
	pipe.Set(ctx, scoreDetailKey(f.UUID), data, 0)
	pipe.ZAdd(ctx, scoreTypeKey(f.UserID, int64(f.Type)), redis_client.Z{Member: f.UUID, Score: float64(f.CreatedAt.UnixMilli())})
	pipe.ZAdd(ctx, scoreKey(f.UserID), redis_client.Z{Member: f.UUID, Score: float64(f.CreatedAt.UnixMilli())})
	_, err = pipe.Exec(ctx)
	return err
}

func (u *Game) GetScoreHistory(id string) (*dto.UserScoreHistory, error) {
	scoreHistory := &dto.UserScoreHistory{}
	val, err := u.Get(ctx, scoreDetailKey(id)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), scoreHistory)
	return scoreHistory, err
}

func (u *Game) SetTotalScore(userId string, count uint64) error {
	return u.Set(ctx, scoreTotalKey(userId), count, 0).Err()
}

func (u *Game) GetTotalScore(userId string) (uint64, error) {
	result, err := u.Get(ctx, scoreTotalKey(userId)).Result()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(result, 10, 64)
}

func (u *Game) GetTotalTypeScore(userId string) (map[uint64]uint64, error) {
	result, err := u.HKeys(ctx, scoreTypeTotalKey(userId)).Result()
	value := make(map[uint64]uint64)
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil || len(result) == 0 {
		return value, err
	}
	for _, v := range result {
		parseUint, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		score, err := u.GetTotalScoreByType(userId, parseUint)
		if err != nil {
			return nil, err
		}
		value[parseUint] = score
	}
	return value, err
}

func (u *Game) GetTotalScoreByType(userId string, scoreType uint64) (uint64, error) {
	val, err := u.HGet(ctx, scoreTypeTotalKey(userId), strconv.FormatUint(scoreType, 10)).Result()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(val, 10, 64)
}

func (u *Game) SetTotalScoreByType(userId string, count uint64, scoreType uint64) error {
	return u.HSet(ctx, scoreTypeTotalKey(userId), strconv.FormatUint(scoreType, 10), count).Err()
}
