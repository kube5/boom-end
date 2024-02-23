package handler

import (
	"context"

	common2 "github.com/Mu-Exchange/Mu-End/common"
	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/social"
	"github.com/Mu-Exchange/Mu-End/game/handler/service"
	"github.com/Mu-Exchange/Mu-End/game/pkg/base"

	"github.com/Mu-Exchange/Mu-End/common/proto"
)

type Game struct {
	repo.CommonComponents
	gs           service.GameService
	tvs          service.StakingService
	userService  proto.UserService
	tweetLimiter *social.TweetRateLimiter
}

func (w *Game) MissionInvited(ctx context.Context, req *proto.MissionInvitedReq, empty *proto.Empty) error {
	return nil
}

func (w *Game) MissionProfile(ctx context.Context, req *proto.UserIdReq, resp *proto.MissionProfileResp) error {
	score, err := w.gs.QueryTotalScore(ctx, req.Id)
	if err != nil {
		return err
	}
	scoreUsed, err := w.gs.QueryTotalScoreUsed(ctx, req.Id)
	if err != nil {
		return err
	}

	cash, err := w.gs.QueryTotalCash(req.Id)
	if err != nil {
		return err
	}

	typeScore, err := w.gs.QueryTotalTypeScore(ctx, req.Id)
	if err != nil {
		return err
	}

	for scoreType, score := range typeScore {
		resp.ScoreAccuracy = common2.MissionTypeMap[scoreType].ScoreAccuracy
		scoreDetail := proto.Mission{
			Name:          common2.MissionTypeMap[scoreType].Name,
			Desc:          common2.MissionTypeMap[scoreType].Desc,
			ScoreAccuracy: common2.MissionTypeMap[scoreType].ScoreAccuracy,
			ScoreType:     scoreType,
			ScoreAmount:   score,
			Daily:         common2.MissionTypeMap[scoreType].Daily,
			MissionScore:  common2.MissionTypeMap[scoreType].ScoreAmount,
		}
		scoreDetail.Undo = true
		if scoreDetail.Daily {
			checkDailyMission, err := w.gs.CheckDailyMission(ctx, req.Id, scoreType)
			if err != nil {
				return err
			}
			scoreDetail.Undo = checkDailyMission
		}
		resp.Mission = append(resp.Mission, &scoreDetail)
	}

	resp.Id = req.Id
	resp.Score = score
	resp.ScoreUsed = scoreUsed
	resp.Cash = cash
	return nil
}

func (w *Game) ScoreList(ctx context.Context, req *proto.ScoreListReq, resp *proto.ScoreListResp) error {
	var scoreHistory []*dto.UserScoreHistory
	var score uint64
	var err error
	if req.ScoreType == 0 {
		scoreHistory, err = w.gs.QueryScoreList(ctx, req.UserId, req.PageNo, req.PageSize)
		if err != nil {
			return err
		}
		score, err = w.gs.QueryTotalScore(ctx, req.UserId)
		if err != nil {
			return err
		}
	} else {
		scoreHistory, err = w.gs.QueryScoreListByType(ctx, req.UserId, req.PageNo, req.PageSize, req.ScoreType)
		if err != nil {
			return err
		}
		score, err = w.gs.QueryTotalScoreByType(ctx, req.UserId, req.ScoreType)
		if err != nil {
			return err
		}
	}
	resp.CurrentScore = score
	resp.ScoreAccuracy = common2.ScoreAccuracy
	for _, item := range scoreHistory {
		scoreDetail := proto.ScoreDetail{
			Id:         item.UUID,
			Score:      item.Score,
			Desc:       item.Desc,
			CreateTime: item.CreatedAt.Unix(),
		}
		resp.Score = append(resp.Score, &scoreDetail)
	}
	return nil
}

func (w *Game) ConsumeList(ctx context.Context, req *proto.ConsumeListReq, resp *proto.ConsumeListResp) error {
	var scoreUsedHistory []*dto.UserScoreUsedHistory
	var score uint64
	var err error
	if req.ConsumeType == 0 {
		scoreUsedHistory, err = w.gs.QueryScoreUsedList(ctx, req.UserId, req.PageNo, req.PageSize)
		if err != nil {
			return err
		}
		score, err = w.gs.QueryTotalScoreUsed(ctx, req.UserId)
		if err != nil {
			return err
		}
	} else {
		scoreUsedHistory, err = w.gs.QueryScoreUsedListByType(ctx, req.UserId, req.PageNo, req.PageSize, req.ConsumeType)
		if err != nil {
			return err
		}
		score, err = w.gs.QueryTotalScoreByType(ctx, req.UserId, req.ConsumeType)
		if err != nil {
			return err
		}
	}
	resp.CurrentConsume = score
	resp.ScoreAccuracy = common2.ScoreAccuracy
	for _, item := range scoreUsedHistory {
		scoreDetail := proto.ScoreDetail{
			Id:         item.UUID,
			Score:      item.Score,
			Desc:       item.Desc,
			CreateTime: item.CreatedAt.Unix(),
		}
		resp.Score = append(resp.Score, &scoreDetail)
	}
	return nil
}

func (w *Game) MissionCheckIn(ctx context.Context, req *proto.UserIdReq, empty *proto.Empty) error {
	return nil
}

func (w *Game) MissionTweet(ctx context.Context, req *proto.MissionTweetReq, empty *proto.Empty) error {
	return nil
}

func init() {
	base.RegisterComponents(NewGame)
}

func NewGame(c repo.CommonComponents,
	gs service.GameService,
	userService proto.UserService,
	tvs service.StakingService) (proto.GameHandler, error) {
	return &Game{
		CommonComponents: c,
		gs:               gs,
		tweetLimiter:     social.NewRateLimiter(),
		tvs:              tvs,
		userService:      userService,
	}, nil
}

func (w *Game) LeaderBoard(ctx context.Context, req *proto.LeaderBoardReq, resp *proto.LeaderBoardResp) error {
	board, err := w.gs.LeaderBoard(ctx, req.Limit)
	if err != nil {
		return err
	}
	resp.Items = board
	return nil
}

func (w *Game) GameGetInviteCode(ctx context.Context, req *proto.UserIdReq, resp *proto.GameGetInviteCodeResp) error {
	return nil
}

func (w *Game) GameRandom(ctx context.Context, req *proto.GameRandomReq, resp *proto.GameRandomResp) error {
	dice1, dice2, score, err := w.gs.GameRandom(ctx, req.UserId, int(req.Level))
	if err != nil {
		return err
	}
	resp.Cash = uint64(score)
	resp.Dice1 = uint64(dice1)
	resp.Dice2 = uint64(dice2)
	return nil
}
