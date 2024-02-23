package common

var ConsumeTypeMap = map[uint64]Consume{
	1: ConsumeRandom,
	2: ConsumeGetInvitedCode,
}

type Consume struct {
	Name          string
	Desc          string
	DescId        string
	ScoreAccuracy uint64
	ScoreType     uint64
	ScoreAmount   uint64
}

var ConsumeRandom = Consume{
	ScoreType:     1,
	Name:          "Random",
	Desc:          "Play Random",
	DescId:        "",
	ScoreAccuracy: ScoreAccuracy,
	ScoreAmount:   10,
}

var ConsumeGetInvitedCode = Consume{
	ScoreType:     2,
	Name:          "InvitedCode",
	Desc:          "Get InvitedCode",
	DescId:        "",
	ScoreAccuracy: ScoreAccuracy,
	ScoreAmount:   10,
}
