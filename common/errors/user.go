package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrAuthToken                   = errors.New("error auth token")               // 10112
	ErrRefreshToken                = errors.New("error refresh token")            // 10109
	ErrAuthValidationExpired       = errors.New("error access expiration token")  // 10117
	ErrRefreshTokenExpired         = errors.New("error refresh expiration token") // 10114
	ErrUserExists                  = errors.New("user already exists")
	ErrUserNameExists              = errors.New("user name already exists")
	ErrUserNotExist                = errors.New("user not exist")
	ErrUserProfileNotExist         = errors.New("user profile not exist")
	ErrNoPermission                = errors.New("no permission")
	ErrFollow                      = errors.New("error follow")
	ErrUnFollow                    = errors.New("error unfollow")
	ErrCreateMch                   = errors.New("error create mch")
	ErrCreateMeta                  = errors.New("error query mch")
	ErrSocialName                  = errors.New("error social name")
	ErrMirrorUri                   = errors.New("can not found mirror uri")
	ErrTelegramVerify              = errors.New("error verify telelegram failed")
	ErrMetaMaskVerifySign          = errors.New("error metamask verify sign")
	ErrHeaderNonceVerify           = errors.New("error header nonce verify")
	ErrGetVerifyCode               = errors.New("error send verify code")
	ErrWaitCulMpt                  = errors.New("calculating mpt,please wait a monent...")
	ErrVerify3Elements             = errors.New("error verify 3 elements")
	ErrVerifyNotExpired            = errors.New("error verify's time not expired")
	ErrVerifyNotExist              = errors.New("error verify not exist")
	ErrVerifyAlready               = errors.New("error user already verify")
	ErrNotExpirationCode           = errors.New("error not expiration code")
	ErrUserAlreadyGuided           = errors.New("error use has already finished guide")
	ErrUserNotVerify               = errors.New("error user not pass 3 elements")
	ErrGuideAnswerIncomplete       = errors.New("error guide answers are not completed")
	ErrProfileImageExpected        = errors.New("error profile image should be provided")
	ErrEmailChangeNotAllowed       = errors.New("error change user's email is not allowed")
	ErrWalletChangeNotAllowed      = errors.New("error change user's wallet is not allowed")
	ErrUnknownSortType             = errors.New("error unknown sort type")
	ErrUnknownDupChkType           = errors.New("error unknown duplicate check type")
	ErrEmailHasBeenBind            = errors.New("error email has been bind")
	ErrSocialHasBeenBind           = errors.New("error social account has been bind")
	ErrWalletHasBeenBind           = errors.New("error wallet has been bind")
	ErrEmailBindCodeNotFound       = errors.New("error email bind code not found")
	ErrEmailCodeTimeOut            = errors.New("error email code has been timeout")
	ErrEnsNotFound                 = errors.New("ENS Not Found") // frontend need to show a different error message
	ErrWalletNotBind               = errors.New("error wallet not bind")
	ErrTwitterNotBind              = errors.New("error twitter not bind")
	ErrTwitterFansNotEnough        = errors.New("error twitter fans not enough")
	ErrKOLApplyExist               = errors.New("error kol apply exist")
	ErrInstitutionApplyExist       = errors.New("error institution apply exist")
	ErrInstitutionApplyNotExist    = errors.New("error institution apply not exist")
	ErrInvalidVerificationLink     = errors.New("error invalid verification link")
	ErrEmailHasBeenBindWhenLogin   = errors.New("error email has been bind, please login by correct wallet")
	ErrHasEmailWhenLoginWhenLogin  = errors.New("error email has been bind, please login by other wallet to bind this email")
	ErrCommingSoon                 = errors.New("Comming soon!")
	ErrTwitterHomeNotMatch         = errors.New("error twitter home not match")
	ErrInstitutionAuthTypeMismatch = errors.New("error institution auth type mismatch")
	ErrInstitutionApplySubmiited   = errors.New("error institution apply submiited")
	ErrWrongVerificationCode       = errors.New("error wrong verification code")
	ErrRequestCodeTooFrequent      = errors.New("error request code too frequent")
	ErrInstitutionEmailNotVerified = errors.New("error institution email not verified")
	ErrInstitutionEmailNotMatch    = errors.New("error institution email not match")

	ErrInviteCodeNotFound    = errors.New("error invite code not found")
	ErrInviteCodeHasUsed     = errors.New("error invite code has been used")
	ErrSocialNotBind         = errors.New("error social not bind")
	ErrTxErr                 = errors.New("error tx")
	ErrNotOneDayApart        = errors.New("error not one day apart")
	ErrTweetLimit            = errors.New("error tweet limit, please check tomorrow")
	ErrTweetNotFoundKeyword  = errors.New("error tweet not found keyword")
	ErrTweetNotBelongToday   = errors.New("error tweet not belong today")
	ErrTweetNotFound         = errors.New("error tweet not found")
	ErrTweetNotBelongUser    = errors.New("error tweet not belong you")
	ErrWrongScoreHistoryId   = errors.New("error wrong score history id")
	ErrUserPostTooFrequently = errors.New("post comment or reply too frequently")
	ErrNotEnoughDice         = errors.New("not enough dice")
	ErrUserHasBeenRegister   = errors.New("err user has been register")
	ErrInvalidParams         = errors.New("error invalid params")
)
