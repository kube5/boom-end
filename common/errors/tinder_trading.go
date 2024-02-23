package errors

import "github.com/pkg/errors"

var (
	ErrNoTraderConfigure      = errors.New("error no trader configure")
	ErrMarginNotEnough        = errors.New("error margin not enough")
	ErrOrderNotFound          = errors.New("error order not found")
	ErrOwnerNotMatch          = errors.New("error owner not match")
	ErrOrderNotOpened         = errors.New("error order not opened")
	ErrOrderIsClosing         = errors.New("error order is closing")
	ErrOrderAlreadyClosed     = errors.New("error order already closed")
	ErrOpenOrderCountReachMax = errors.New("error open order count reach max")
	ErrOpenOrderTooFast       = errors.New("error open order too fast")
	ErrCloseOrderTooFast      = errors.New("error close order too fast")
	ErrOpenOrder              = errors.New("error open order")
	ErrCloseOrder             = errors.New("error close order")

	ErrGetPriceData         = errors.New("error get price data")
	ErrOutOfPriceLimit      = errors.New("error out of price limit")
	ErrOutOfCollateralLimit = errors.New("error out of collateral limit")
	ErrOutOfInterestLimit   = errors.New("error out of interest limit")
	ErrOutOfSpread          = errors.New("error out of spread")
	ErrTradeAlreadyExisted  = errors.New("error trade already existed")
	ErrIllegal              = errors.New("error illegal")
	ErrOutSize              = errors.New("error out size")
	ErrNotExist             = errors.New("error not exist")
	ErrContract             = errors.New("error contract")
)
