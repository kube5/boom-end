package service

import (
	"context"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
)

func (a *Game) AddScoreHistory(history *dto.UserScoreHistory) error {
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

func (a *Game) QueryTotalTypeScore(ctx context.Context, userId string) (map[uint64]uint64, error) {
	totalTypeScore, err := a.gameCache.GetTotalTypeScore(userId)
	if err != nil {
		return nil, err
	}
	return totalTypeScore, err
}

func (a *Game) QueryTotalScore(ctx context.Context, userId string) (uint64, error) {
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

func (a *Game) QueryTotalScoreByType(ctx context.Context, userId string, scoreType uint64) (uint64, error) {
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

func (a *Game) QueryScoreList(ctx context.Context, userId string, pageNo, pageSize int64) ([]*dto.UserScoreHistory, error) {
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

func (a *Game) QueryScoreListByType(ctx context.Context, userId string, pageNo, pageSize int64, scoreType uint64) ([]*dto.UserScoreHistory, error) {
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

func (a *Game) QueryScoreHistory(tx context.Context, id string) (*dto.UserScoreHistory, error) {
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
