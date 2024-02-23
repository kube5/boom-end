package cache

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	redis_client "github.com/redis/go-redis/v9"
)

func scoreUsedTotalKey(userId string) string {
	return fmt.Sprintf("%s-SCOREUSED-TOTAL-%s", ServicePrefix, userId)
}
func scoreUsedTypeTotalKey(userId string) string {
	return fmt.Sprintf("%s-SCOREUSED-TYPE-TOTAL-%s", ServicePrefix, userId)
}
func scoreUsedKey(userId string) string {
	return fmt.Sprintf("%s-SCOREUSED-%s", ServicePrefix, userId)
}
func scoreUsedTypeKey(userId string, scoreUsedType int64) string {
	return fmt.Sprintf("%s-SCOREUSED-%s-%d", ServicePrefix, userId, scoreUsedType)
}
func scoreUsedDetailKey(id string) string {
	return fmt.Sprintf("%s-SCOREUSED-DETAIL-%s", ServicePrefix, id)
}

func (u *Game) QueryScoreUsedList(userId string, pageNo, pageSize int64) ([]string, error) {
	nums, err := u.ZCard(ctx, scoreUsedKey(userId)).Result()
	if err != nil {
		return nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	return u.ZRevRange(ctx, scoreUsedKey(userId), start, end).Result()
}

func (u *Game) QueryScoreUsedListByType(userId string, pageNo, pageSize, scoreUsedType int64) ([]string, error) {
	nums, err := u.ZCard(ctx, scoreUsedTypeKey(userId, scoreUsedType)).Result()
	if err != nil {
		return nil, err
	}
	start, end := (pageNo-1)*pageSize, pageSize*pageNo-1
	if pageNo*pageSize > nums {
		end = nums
	}
	return u.ZRevRange(ctx, scoreUsedTypeKey(userId, scoreUsedType), start, end).Result()
}

func (u *Game) AddScoreUsedHistory(f *dto.UserScoreUsedHistory) error {
	history, err := u.GetScoreUsedHistory(f.UUID)
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
	pipe.Set(ctx, scoreUsedDetailKey(f.UUID), data, 0)
	pipe.ZAdd(ctx, scoreUsedTypeKey(f.UserID, int64(f.Type)), redis_client.Z{Member: f.UUID, Score: float64(f.CreatedAt.UnixMilli())})
	pipe.ZAdd(ctx, scoreUsedKey(f.UserID), redis_client.Z{Member: f.UUID, Score: float64(f.CreatedAt.UnixMilli())})
	_, err = pipe.Exec(ctx)
	return err
}

func (u *Game) GetScoreUsedHistory(id string) (*dto.UserScoreUsedHistory, error) {
	scoreUsedHistory := &dto.UserScoreUsedHistory{}
	val, err := u.Get(ctx, scoreUsedDetailKey(id)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), scoreUsedHistory)
	return scoreUsedHistory, err
}

func (u *Game) SetTotalScoreUsed(userId string, count uint64) error {
	return u.Set(ctx, scoreUsedTotalKey(userId), count, 0).Err()
}

func (u *Game) GetTotalScoreUsed(userId string) (uint64, error) {
	result, err := u.Get(ctx, scoreUsedTotalKey(userId)).Result()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(result, 10, 64)
}

func (u *Game) GetTotalTypeScoreUsed(userId string) (map[uint64]uint64, error) {
	result, err := u.HKeys(ctx, scoreUsedTypeTotalKey(userId)).Result()
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
		scoreUsed, err := u.GetTotalScoreUsedByType(userId, parseUint)
		if err != nil {
			return nil, err
		}
		value[parseUint] = scoreUsed
	}
	return value, err
}

func (u *Game) GetTotalScoreUsedByType(userId string, scoreUsedType uint64) (uint64, error) {
	val, err := u.HGet(ctx, scoreUsedTypeTotalKey(userId), strconv.FormatUint(scoreUsedType, 10)).Result()
	if err == redis_client.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(val, 10, 64)
}

func (u *Game) SetTotalScoreUsedByType(userId string, count uint64, scoreUsedType uint64) error {
	return u.HSet(ctx, scoreUsedTypeTotalKey(userId), strconv.FormatUint(scoreUsedType, 10), count).Err()
}
