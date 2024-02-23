package model

type GameRandomReq struct {
	Level int `json:"level" form:"level" binding:"required"`
}

func (req GameRandomReq) ErrMessages() map[string]string {
	return map[string]string{
		"Level.required": "Level is empty",
	}
}

type MetaMaskLoginPreReq struct {
	PublicAddress string `json:"public_address" form:"public_address" binding:"required"`
}

func (req MetaMaskLoginPreReq) ErrMessages() map[string]string {
	return map[string]string{
		"PublicAddress.required": "PublicAddress is empty",
	}
}

type RegisterByInviteCodeReq struct {
	Code string `json:"code" form:"code" binding:"required"`
}

func (req RegisterByInviteCodeReq) ErrMessages() map[string]string {
	return map[string]string{
		"Code.required": "code is empty",
	}
}

type MissionTweetCheckInReq struct {
	TweetId string `json:"tweet_id" form:"tweet_id" binding:"required"`
}

func (req MissionTweetCheckInReq) ErrMessages() map[string]string {
	return map[string]string{
		"TweetId.required": "tweet_id is empty",
	}
}

type LoginInternalReq struct {
	Address string `json:"address" form:"address" binding:"required"`
}

func (req LoginInternalReq) ErrMessages() map[string]string {
	return map[string]string{
		"Address.required": "Address is empty",
	}
}

type MetaMaskLoginReq struct {
	PublicAddress string `json:"public_address" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
	EmailCode     string `json:"email_code"`
}

func (req MetaMaskLoginReq) ErrMessages() map[string]string {
	return map[string]string{
		"PublicAddress.required": "PublicAddress is empty",
		"Signature.required":     "Signature is empty",
	}
}

type SocialBindReq struct {
	SocialName string `json:"social_name" binding:"required"`
	Code       string `json:"code" binding:"required"`
}

func (req SocialBindReq) ErrMessages() map[string]string {
	return map[string]string{
		"SocialName.required": "SocialName is empty",
		"Code.required":       "Code is empty",
	}
}

type SocialBindUriReq struct {
	SocialName string `form:"social_name" json:"social_name" binding:"required"`
}

func (req SocialBindUriReq) ErrMessages() map[string]string {
	return map[string]string{
		"SocialName.required": "SocialName is empty",
	}
}

type UserReq struct {
	PhoneNumber string `json:"phone_number" binding:"required,phone"`
	Code        string `json:"code" binding:"required,max=6"`
	InvitedCode string `json:"invited_code" binding:"max=8"`
}

func (req UserReq) ErrMessages() map[string]string {
	return map[string]string{
		"PhoneNumber.required": "phone_number is empty",
		"Code.required":        "code is empty",
		"PhoneNumber.phone":    "phone_number format incorrect",
		"Code.max":             "maximum code length is 6",
		"InvitedCode.max":      "maximum invited_code length is 8",
	}
}

type TargetUserReq struct {
	TargetUserId string `json:"target_user_id" form:"target_user_id" binding:"required"`
}

func (req TargetUserReq) ErrMessages() map[string]string {
	return map[string]string{
		"TargetUserId.required": "target_user_id is empty",
	}
}

type UserIdReq struct {
	TargetUserId string `json:"target_user_id" form:"target_user_id"`
}

func (req UserIdReq) ErrMessages() map[string]string {
	return map[string]string{}
}

type UserIdPageReq struct {
	TargetUserId string `form:"target_user_id" json:"target_user_id"`
	PageNo       int64  `form:"page_no" json:"page_no" binding:"required,gt=0"`
	PageSize     int64  `form:"page_size" json:"page_size" binding:"required,gte=5,max=20"`
}

func (req UserIdPageReq) ErrMessages() map[string]string {
	return map[string]string{
		"PageNo.required":   "page_no is empty",
		"PageNo.gt":         "required page_no > 0",
		"PageSize.gte":      "required page_size >= 5",
		"PageSize.required": "page_size is empty",
		"PageSize.max":      "maximum page_size length is 20",
	}
}

type UrlReq struct {
	Url string `json:"url" form:"url" binding:"required,url"`
}

func (req UrlReq) ErrMessages() map[string]string {
	return map[string]string{
		"Url.required": "url is empty",
		"Url.url":      "url is not URL String",
	}
}
