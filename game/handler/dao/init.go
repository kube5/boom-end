package dao

import "github.com/Mu-Exchange/Mu-End/game/pkg/base"

func init() {
	base.RegisterComponents(
		NewBaseDao,
		NewScoreUsedDao,
		NewScoreDao,
		NewCashDao,
		NewGameContractDao,
		NewGameMonitorRecordDao,
	)
}
