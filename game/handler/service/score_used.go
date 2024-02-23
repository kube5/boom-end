package service

import (
	"context"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
)

func (a *Game) AddScoreUsedHistory(history *dto.UserScoreUsedHistory) error {
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
	return nil
}

func (a *Game) QueryTotalTypeScoreUsed(ctx context.Context, userId string) (map[uint64]uint64, error) {
	totalTypeScoreUsed, err := a.gameCache.GetTotalTypeScoreUsed(userId)
	if err != nil {
		return nil, err
	}
	return totalTypeScoreUsed, err
}

func (a *Game) QueryTotalScoreUsed(ctx context.Context, userId string) (uint64, error) {
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

func (a *Game) QueryTotalScoreUsedByType(ctx context.Context, userId string, scoreType uint64) (uint64, error) {
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

func (a *Game) QueryScoreUsedList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreUsedHistory, error) {
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

func (a *Game) QueryScoreUsedListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreUsedHistory, error) {
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

func (a *Game) QueryScoreUsedHistory(tx context.Context, id string) (*dto.UserScoreUsedHistory, error) {
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
