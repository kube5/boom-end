package common

import (
	"math"
	"math/big"

	"github.com/pkg/errors"
)

const (
	MissionReferralFatherType = 7
	MissionReferralGrandType  = 8
	ScoreAccuracy             = 1
)

func init() {
	MissionTypeMap = make(map[uint64]Mission, len(Missions))
	for _, v := range Missions {
		MissionTypeMap[v.ScoreType] = v
	}
}

var MissionTypeMap = map[uint64]Mission{}

type Mission struct {
	Name             string
	Desc             string
	DescId           string
	ScoreAccuracy    uint64
	ScoreType        uint64
	ScoreAmount      uint64
	Daily            bool
	Referral         bool
	ReferralMissions []Mission
	ReferralRule     func(uint64) uint64

	ScoreLimit    uint64
	ScoreStart    uint64
	ScoreRule     func([]interface{}) (uint64, error)
	ScoreRuleArgs []interface{}
}

var Missions = []Mission{
	MissionCheckIn, MissionTweet, MissionInvite, MissionTrading, MissionBridging, MissionDeposit, MissionReferralFather, MissionReferralGrand,
}

var MissionCheckIn = Mission{
	ScoreType:     1,
	Name:          "CheckIn",
	Desc:          "Check in daily",
	ScoreAccuracy: ScoreAccuracy,
	ScoreAmount:   5,
	Daily:         true,
	Referral:      false,
}

var MissionTweet = Mission{
	ScoreType:     2,
	Name:          "Tweet",
	Desc:          "Post a tweet daily",
	ScoreAccuracy: ScoreAccuracy,
	ScoreAmount:   10,
	Daily:         true,
	Referral:      false,
}

var MissionInvite = Mission{
	ScoreType:     3,
	Name:          "Invite",
	Desc:          "Invite a user",
	ScoreAccuracy: ScoreAccuracy,
	ScoreAmount:   10,
	Daily:         false,
	Referral:      false,
}

var MissionTrading = Mission{
	ScoreType:     4,
	Name:          "Trading",
	Desc:          "Exec one trade",
	ScoreAccuracy: ScoreAccuracy,

	Daily:            false,
	Referral:         true,
	ReferralMissions: []Mission{MissionReferralFather, MissionReferralGrand},
	ScoreLimit:       10000,
	ScoreStart:       10,
	ScoreRule:        tradingScoreRule,
}

var MissionBridging = Mission{
	ScoreType:     5,
	Name:          "Bridging",
	Desc:          "Exec one Bridging",
	ScoreAccuracy: ScoreAccuracy,

	Daily:            false,
	Referral:         true,
	ReferralMissions: []Mission{MissionReferralFather, MissionReferralGrand},
	ScoreRule:        bridgingScoreRule,
}

var MissionDeposit = Mission{
	ScoreType:     6,
	Name:          "Deposit",
	Desc:          "Exec one Deposit",
	ScoreAccuracy: ScoreAccuracy,

	Daily:            false,
	Referral:         true,
	ReferralMissions: []Mission{MissionReferralFather, MissionReferralGrand},
	ScoreRule:        depositScoreRule,
}

var MissionReferralFather = Mission{
	ScoreType:     MissionReferralFatherType,
	Name:          "Referral Father",
	Desc:          "Exec one to father",
	ScoreAccuracy: ScoreAccuracy,

	Daily:        false,
	ReferralRule: referralFatherRule,
}

var MissionReferralGrand = Mission{
	ScoreType:     MissionReferralGrandType,
	Name:          "Referral Grand",
	Desc:          "Exec one to grand",
	ScoreAccuracy: ScoreAccuracy,

	Daily:        false,
	ReferralRule: referralGrandRule,
}

func referralGrandRule(amount uint64) uint64 {
	if amount == 0 {
		return 0
	}
	result := 8 * amount / 100
	return result

}
func referralFatherRule(amount uint64) uint64 {
	if amount == 0 {
		return 0
	}
	result := 16 * amount / 100
	return result
}

func tradingScoreRule(args []interface{}) (uint64, error) {
	//convert an expression of the type 'interface{}' to the type 'float64'
	if len(args) == 0 {
		return 0, errors.New("invalid tradingScoreRule arguments arg len == 0")
	}
	switch v := args[0].(type) {
	case *big.Int:
		b := v
		result := uint64(0.5 * (math.Sqrt(float64(b.Int64())) * ScoreAccuracy))
		return result, nil
	default:
		return 0, errors.New("invalid tradingScoreRule arguments not big.Int")
	}

}

func bridgingScoreRule(args []interface{}) (uint64, error) {
	//convert an expression of the type 'interface{}' to the type 'float64'
	if len(args) == 0 {
		return 0, errors.New("invalid bridgingScoreRule arguments arg len == 0")
	}
	switch v := args[0].(type) {
	case *big.Int:
		b := v
		result := uint64(5 * (math.Sqrt(float64(b.Int64())) * ScoreAccuracy))
		return result, nil
	default:
		return 0, errors.New("invalid bridgingScoreRule arguments not big.Int")
	}

}

func depositScoreRule(args []interface{}) (uint64, error) {
	//convert an expression of the type 'interface{}' to the type 'float64'
	if len(args) == 0 {
		return 0, errors.New("invalid depositScoreRule arguments arg len == 0")
	}
	switch v := args[0].(type) {
	case *big.Int:
		b := v
		result := uint64(5 * (math.Sqrt(float64(b.Int64())) * ScoreAccuracy))
		return result, nil
	default:
		return 0, errors.New("invalid depositScoreRule arguments not big.Int")
	}

}
