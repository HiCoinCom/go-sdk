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
	DepositAddress string          `json:"deposit_address"`
	LockBalance    decimal.Decimal `json:"lock_balance"`
	NormalBalance  decimal.Decimal `json:"normal_balance"`
}

type CompanyAccount struct {
	Balance           decimal.Decimal `json:"balance"`
	FeeAccountBalance decimal.Decimal `json:"fee_account_balance"`
	Symbol            string          `json:"symbol"`
	TotalBalance      decimal.Decimal `json:"total_balance"`
}

// AccountResult represents account response.
type AccountResult struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data *Account `json:"data"`
}

type CompanyAccountResult struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data *CompanyAccount `json:"data"`
}

// UserAddress represents user deposit address.
type UserAddress struct {
	Id      FlexInt `json:"id"`
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
	AddressLink          string          `json:"address_link"`
	AddressRegex         string          `json:"address_regex"`
	AddressTagRegex      string          `json:"address_tag_regex"`
	BaseSymbol           string          `json:"base_symbol"`
	CoinNet              string          `json:"coin_net"`
	ContractAddress      string          `json:"contract_address"`
	Decimals             FlexInt         `json:"decimals"`
	DepositConfirmation  FlexInt         `json:"deposit_confirmation"`
	Explorer             string          `json:"explorer"`
	Icon                 string          `json:"icon"`
	MarginSymbol         bool            `json:"margin_symbol"`
	MergeAddressSymbol   string          `json:"merge_address_symbol"`
	MinDeposit           decimal.Decimal `json:"min_deposit"`
	RealSymbol           string          `json:"real_symbol"`
	SupportMemo          string          `json:"support_memo"`
	SupportToken         string          `json:"support_token"`
	Symbol               string          `json:"symbol"`
	SymbolAlias          string          `json:"symbol_alias"`
	TxidLink             string          `json:"txid_link"`
	WithdrawConfirmation FlexInt         `json:"withdraw_confirmation"`
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
	RequestID         string          `json:"request_id"`
	AddressFrom       string          `json:"address_from"`
	AddressTo         string          `json:"address_to"`
	Amount            decimal.Decimal `json:"amount"`
	BaseSymbol        string          `json:"base_symbol"`
	CompanyStatus     FlexInt         `json:"company_status"`
	Confirmations     FlexInt         `json:"confirmations"`
	ContractAddress   string          `json:"contract_address"`
	CreatedAt         Timestamp       `json:"created_at"`
	Email             string          `json:"email"`
	Fee               decimal.Decimal `json:"fee"`
	FeeSymbol         string          `json:"fee_symbol"`
	Id                FlexInt         `json:"id"`
	RealFee           decimal.Decimal `json:"real_fee"`
	RequestId         string          `json:"request_id"`
	SaasStatus        FlexInt         `json:"saas_status"`
	Status            int64           `json:"status"`
	Symbol            string          `json:"symbol"`
	Txid              string          `json:"txid"`
	TxidType          string          `json:"txid_type"`
	Uid               int64           `json:"uid"`
	UpdatedAt         Timestamp       `json:"updated_at"`
	WithdrawFee       decimal.Decimal `json:"withdraw_fee"`
	WithdrawFeeSymbol string          `json:"withdraw_fee_symbol"`
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
	AddressTo       string          `json:"address_to"`
	Amount          decimal.Decimal `json:"amount"`
	BaseSymbol      string          `json:"base_symbol"`
	Confirmations   FlexInt         `json:"confirmations"`
	ContractAddress string          `json:"contract_address"`
	CreatedAt       Timestamp       `json:"created_at"`
	Email           string          `json:"email"`
	ID              int             `json:"id"`
	IsMining        int             `json:"is_mining"`
	Status          FlexInt         `json:"status"`
	Symbol          string          `json:"symbol"`
	Txid            string          `json:"txid"`
	TxidType        FlexInt         `json:"txid_type"`
	Uid             int             `json:"uid"`
	UpdatedAt       Timestamp       `json:"updated_at"`
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
	AddressFrom     string          `json:"address_from"`
	AddressTo       string          `json:"address_to"`
	Amount          decimal.Decimal `json:"amount"`
	BaseSymbol      string          `json:"base_symbol"`
	Confirmations   FlexInt         `json:"confirmations"`
	ContractAddress string          `json:"contract_address"`
	CreatedAt       Timestamp       `json:"created_at"`
	Email           string          `json:"email"`
	Fee             decimal.Decimal `json:"fee"`
	ID              int             `json:"id"`
	Status          FlexInt         `json:"status"`
	Symbol          string          `json:"symbol"`
	Txid            string          `json:"txid"`
	TxidType        FlexInt         `json:"txid_type"`
	UpdatedAt       decimal.Decimal `json:"updated_at"`
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
