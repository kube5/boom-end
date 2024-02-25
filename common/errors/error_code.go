package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// UnknownErrorCode means unknown error
	UnknownErrorCode = 10001
	UnknownErrorMsg  = "something wrong!"
)

var errorCodeMap = map[string]uint32{
	// common
	ErrRequestParameter.Error():         10003,
	ErrRequestParameterAllEmpty.Error(): 10004,
	ErrDuplicateSubscription.Error():    10005,
	ErrInternal.Error():                 10006,

	// user
	ErrUserExists.Error(): 10101,

	ErrNoPermission.Error():                10105,
	ErrFollow.Error():                      10106,
	ErrUnFollow.Error():                    10107,
	ErrCreateMch.Error():                   10108,
	ErrRefreshToken.Error():                10109,
	ErrGetVerifyCode.Error():               10110,
	ErrVerify3Elements.Error():             10111,
	ErrAuthToken.Error():                   10112,
	ErrRefreshTokenExpired.Error():         10114,
	ErrNotExpirationCode.Error():           10116,
	ErrAuthValidationExpired.Error():       10117,
	ErrUserNotVerify.Error():               10118,
	ErrCreateMeta.Error():                  10119,
	ErrVerifyNotExist.Error():              10121,
	ErrVerifyNotExpired.Error():            10122,
	ErrVerifyAlready.Error():               10123,
	ErrUserNameExists.Error():              10124,
	ErrUserProfileNotExist.Error():         10125,
	ErrMetaMaskVerifySign.Error():          10126,
	ErrHeaderNonceVerify.Error():           10127,
	ErrUserAlreadyGuided.Error():           10128,
	ErrGuideAnswerIncomplete.Error():       10129,
	ErrProfileImageExpected.Error():        10130,
	ErrEmailChangeNotAllowed.Error():       10131,
	ErrWalletChangeNotAllowed.Error():      10132,
	ErrUnknownSortType.Error():             10133,
	ErrUnknownDupChkType.Error():           10134,
	ErrEmailHasBeenBind.Error():            10135,
	ErrEmailCodeTimeOut.Error():            10136,
	ErrEmailBindCodeNotFound.Error():       10137,
	ErrMirrorUri.Error():                   10138,
	ErrSocialName.Error():                  10139,
	ErrWalletHasBeenBind.Error():           10140,
	ErrEnsNotFound.Error():                 10141,
	ErrTelegramVerify.Error():              10142,
	ErrSocialHasBeenBind.Error():           10143,
	ErrWaitCulMpt.Error():                  10144,
	ErrWalletNotBind.Error():               10145,
	ErrTwitterNotBind.Error():              10146,
	ErrTwitterFansNotEnough.Error():        10147,
	ErrKOLApplyExist.Error():               10148,
	ErrInstitutionApplyExist.Error():       10149,
	ErrInstitutionApplyNotExist.Error():    10150,
	ErrInvalidVerificationLink.Error():     10151,
	ErrEmailHasBeenBindWhenLogin.Error():   10152,
	ErrHasEmailWhenLoginWhenLogin.Error():  10153,
	ErrCommingSoon.Error():                 10154,
	ErrTwitterHomeNotMatch.Error():         10155,
	ErrInstitutionAuthTypeMismatch.Error(): 10156,
	ErrInstitutionApplySubmiited.Error():   10157,
	ErrWrongVerificationCode.Error():       10158,
	ErrRequestCodeTooFrequent.Error():      10159,
	ErrInstitutionEmailNotVerified.Error(): 10160,
	ErrInstitutionEmailNotMatch.Error():    10161,

	ErrInviteCodeNotFound.Error():   10162,
	ErrInviteCodeHasUsed.Error():    10163,
	ErrSocialNotBind.Error():        10164,
	ErrTxErr.Error():                10165,
	ErrNotOneDayApart.Error():       10166,
	ErrTweetLimit.Error():           10167,
	ErrTweetNotFoundKeyword.Error(): 10168,
	ErrTweetNotBelongToday.Error():  10169,
	ErrWrongScoreHistoryId.Error():  10170,
	ErrTweetNotFound.Error():        10171,
	ErrTweetNotBelongUser.Error():   10172,
	ErrUserNotExist.Error():         10102,
	ErrUserHasBeenRegister.Error():  10173,

	// utility
	ErrExceedMaxFilesSize.Error():      12001,
	ErrGenPresignUrlNotAllowed.Error(): 12002,
	ErrUnknownPresignUrlType.Error():   12003,
	ErrUnsupportedFileFormat.Error():   12004,
	ErrInvalidProjectId.Error():        12005,

	ErrUserPostTooFrequently.Error(): 13009,
	ErrNotEnoughDice.Error():         13010,

	// monitor
	ErrBlockchainNotFound.Error():      16001,
	ErrRpcUrlNotSet.Error():            16002,
	ErrDialBlockchainRpc.Error():       16003,
	ErrInvalidContractType.Error():     16004,
	ErrAllRPCFailed.Error():            16005,
	ErrInvalidChainId.Error():          16006,
	ErrInvalidStartBlock.Error():       16007,
	ErrDuplicateChannel.Error():        16008,
	ErrMonitorHasPublish.Error():       16009,
	ErrCanNotGetChildContract.Error():  16010,
	ErrUnknownContractType.Error():     16011,
	ErrStopMonitorNotAllowed.Error():   16012,
	ErrMonitorStartBlockTooLow.Error(): 16013,

	// tinder trading
	ErrNoTraderConfigure.Error():      17001,
	ErrMarginNotEnough.Error():        17002,
	ErrOrderNotFound.Error():          17003,
	ErrOwnerNotMatch.Error():          17004,
	ErrOrderNotOpened.Error():         17005,
	ErrOrderIsClosing.Error():         17006,
	ErrOrderAlreadyClosed.Error():     17007,
	ErrOpenOrderCountReachMax.Error(): 17008,
	ErrOpenOrder.Error():              17009,
	ErrCloseOrder.Error():             17010,
	ErrOpenOrderTooFast.Error():       17011,
	ErrCloseOrderTooFast.Error():      17012,
}

type CustomMsgError struct {
	error
	errMsg string
}

func (e CustomMsgError) Error() string {
	return fmt.Sprintf("%v.Error(): %s", e.error, e.errMsg)
}

func Wrap(err error, errMsg string) error {
	return CustomMsgError{
		error:  err,
		errMsg: errMsg,
	}
}

func DecodeError(customErr error) uint32 {
	customMsgError, ok := customErr.(CustomMsgError)
	if ok {
		customErr = customMsgError.error
	}

	code, ok := errorCodeMap[errors.Cause(customErr).Error()]
	if !ok {
		return UnknownErrorCode
	}

	return code
}
