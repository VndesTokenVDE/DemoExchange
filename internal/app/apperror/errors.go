package apperror

//lint:file-ignore ST1005 strings capitalized
var (
	ErrAccountNotFound             = New("Account not found")
	ErrPositionNotFound            = New("Position not found")
	ErrTokenNotFound               = New("Token not found")
	ErrTokenLimitExceeded          = New("Token limit exceeded")
	ErrRequestError                = New("Request error")
	ErrBalanceLimitExceeded        = New("Balance limit exceeded")
	ErrAppendBalanceCoinNotAllowed = New("Append balance coin not allowed")
	ErrBalanceNotFound             = New("Balance not found")
	ErrInvalidPositionSide         = New("Invalid position side")
	ErrInsufficientFunds           = New("Insufficient funds")
	ErrSetPositionMode             = New("Error set position mode")
	ErrSetLeverage                 = New("Error set leverage")
	ErrSetMarginType               = New("Error set margin type")
	ErrOpenOrdersExists            = New("Open orders exists")
	ErrPositionExists              = New("Position exists")
	ErrOrderAlreadyCancelled       = New("Order already cancelled")
	ErrExchangeIsNotValid          = New("Exchange is not valid")
	ErrSymbolIsNotValid            = New("Symbol is not valid")
	ErrAmountIsNotValid            = New("Amount is not valid")
	ErrAmountIsOutOfRange          = New("Amount is out of range")
	ErrPriceIsNotValid             = New("Price is not valid")
	ErrMarketPriceIsWrong          = New("Market price is wrong")
	ErrOrderTypeNotValid           = New("Order Type not valid")
	ErrOrderSideIsNotValid         = New("Order Side is not valid")
	ErrOrderPositionModeIsNotValid = New("Order Position mode is not valid")
	ErrOrderNotFound               = New("Order not found")
)
