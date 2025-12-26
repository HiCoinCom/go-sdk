// Package types defines data structures for Custody/WaaS API requests and responses.
package types

import (
	"chainup.com/go-sdk/utils"
	"github.com/shopspring/decimal"
)

// -----------------------------------------------------------------------------
// Common Types (aliases for utils types for backward compatibility)
// -----------------------------------------------------------------------------

// FlexInt is an alias for utils.FlexInt for backward compatibility.
type FlexInt = utils.FlexInt

// Timestamp is an alias for utils.Timestamp for backward compatibility.
type Timestamp = utils.Timestamp

// -----------------------------------------------------------------------------
// User Types
// -----------------------------------------------------------------------------

// UserInfo represents user information.
type UserInfo struct {
	UID      FlexInt `json:"uid"`
	Nickname string  `json:"nickname"`
}

// UserInfoResult represents user info response.
type UserInfoResult struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data *UserInfo `json:"data"`
}

// UserListResult represents user list response.
type UserListResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*UserInfo `json:"data"`
}

// -----------------------------------------------------------------------------
// Account Types
// -----------------------------------------------------------------------------

// Account represents account information.
type Account struct {
	Balance decimal.Decimal `json:"balance"`
	Symbol  string          `json:"symbol"`
}

// AccountResult represents account response.
type AccountResult struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data *Account `json:"data"`
}

// UserAddress represents user deposit address.
type UserAddress struct {
	UID     FlexInt `json:"uid"`
	Address string  `json:"address"`
	Symbol  string  `json:"symbol"`
	Memo    string  `json:"memo,omitempty"`
}

// UserAddressResult represents user address response.
type UserAddressResult struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data *UserAddress `json:"data"`
}

// UserAddressListResult represents user address list response.
type UserAddressListResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []*UserAddress `json:"data"`
}

// -----------------------------------------------------------------------------
// Coin Types
// -----------------------------------------------------------------------------

// CoinInfo represents coin information.
type CoinInfo struct {
	Symbol      string  `json:"symbol"`
	Name        string  `json:"name"`
	Decimals    FlexInt `json:"decimals"`
	TokenStatus FlexInt     `json:"token_status"`
}

// CoinInfoListResult represents coin list response.
type CoinInfoListResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*CoinInfo `json:"data"`
}

// -----------------------------------------------------------------------------
// Withdraw Types
// -----------------------------------------------------------------------------

// Withdraw represents withdrawal information.
type Withdraw struct {
	ID                int64           `json:"id"`
	RequestID         string          `json:"request_id"`
	UID               FlexInt         `json:"uid"`
	Symbol            string          `json:"symbol"`
	Amount            decimal.Decimal `json:"amount"`
	WithdrawFeeSymbol string          `json:"withdraw_fee_symbol"`
	WithdrawFee       decimal.Decimal `json:"withdraw_fee"`
	FeeSymbol         string          `json:"fee_symbol"`
	RealFee           decimal.Decimal `json:"real_fee"`
	AddressTo         string          `json:"address_to"`
	CreatedAt         Timestamp       `json:"created_at"`
	UpdatedAt         Timestamp       `json:"updated_at"`
	TxID              string          `json:"txid"`
	Confirmations     FlexInt         `json:"confirmations"`
	Status            FlexInt         `json:"status"`
}

// WithdrawResult represents withdrawal response.
type WithdrawResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID int64 `json:"id"`
	} `json:"data"`
}

// WithdrawListResult represents withdrawal list response.
type WithdrawListResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*Withdraw `json:"data"`
}

// -----------------------------------------------------------------------------
// Deposit Types
// -----------------------------------------------------------------------------

// Deposit represents deposit information.
type Deposit struct {
	ID            int64           `json:"id"`
	UID           FlexInt         `json:"uid"`
	Symbol        string          `json:"symbol"`
	Amount        decimal.Decimal `json:"amount"`
	AddressFrom   string          `json:"address_from"`
	AddressTo     string          `json:"address_to"`
	TxID          string          `json:"txid"`
	Confirmations FlexInt         `json:"confirmations"`
	Status        FlexInt         `json:"status"`
	CreatedAt     Timestamp       `json:"created_at"`
}

// DepositListResult represents deposit list response.
type DepositListResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data []*Deposit `json:"data"`
}

// -----------------------------------------------------------------------------
// MinerFee Types
// -----------------------------------------------------------------------------

// MinerFee represents miner fee information.
type MinerFee struct {
	ID        int64           `json:"id"`
	Symbol    string          `json:"symbol"`
	Fee       decimal.Decimal `json:"fee"`
	FeeSymbol string          `json:"fee_symbol"`
	TxID      string          `json:"txid"`
	Status    FlexInt             `json:"status"`
}

// MinerFeeListResult represents miner fee list response.
type MinerFeeListResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*MinerFee `json:"data"`
}

// -----------------------------------------------------------------------------
// Transfer Types
// -----------------------------------------------------------------------------

// Transfer represents transfer information.
type Transfer struct {
	ID        int64           `json:"id"`
	RequestID string          `json:"request_id"`
	Symbol    string          `json:"symbol"`
	Amount    decimal.Decimal `json:"amount"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	CreatedAt Timestamp       `json:"created_at"`
	Receipt   string          `json:"receipt"`
	Remark    string          `json:"remark"`
}

// TransferResult represents transfer response.
type TransferResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID int64 `json:"id"`
	} `json:"data"`
}

// TransferListResult represents transfer list response.
type TransferListResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data []*Transfer `json:"data"`
}

// -----------------------------------------------------------------------------
// Async Notify Types
// -----------------------------------------------------------------------------

// AsyncNotifyArgs represents async notification data from webhook callbacks.
type AsyncNotifyArgs struct {
	Side              string          `json:"side"`                // Notification side (deposit/withdraw)
	NotifyTime        Timestamp       `json:"notify_time"`         // Notification time
	RequestID         string          `json:"request_id"`          // Request ID
	ID                FlexInt         `json:"id"`                  // Record ID
	UID               FlexInt         `json:"uid"`                 // User ID
	Symbol            string          `json:"symbol"`              // Cryptocurrency symbol
	Amount            decimal.Decimal `json:"amount"`              // Amount
	WithdrawFeeSymbol string          `json:"withdraw_fee_symbol"` // Withdraw fee symbol
	WithdrawFee       decimal.Decimal `json:"withdraw_fee"`        // Withdraw fee amount
	FeeSymbol         string          `json:"fee_symbol"`          // Fee symbol
	RealFee           decimal.Decimal `json:"real_fee"`            // Actual fee
	AddressTo         string          `json:"address_to"`          // Destination address
	CreatedAt         Timestamp       `json:"created_at"`          // Creation time
	UpdatedAt         Timestamp       `json:"updated_at"`          // Update time
	TxID              string          `json:"txid"`                // Transaction ID
	Confirmations     FlexInt         `json:"confirmations"`       // Number of confirmations
	Status            FlexInt         `json:"status"`              // Status code
}
