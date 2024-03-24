package dao

import "github.com/Mu-Exchange/Mu-End/user/pkg/base"

func init() {
	base.RegisterComponents(
		NewBaseDao,
		NewUserDao,
		NewTgUserDao,
	)
}
