// Package types defines data structures for MPC API requests and responses.
package types

import (
	"chainup.com/go-sdk/utils"
	"github.com/shopspring/decimal"
)

// AppShowStatus represents wallet display status in the app.
type AppShowStatus int

// AppShowStatus constants.
const (
	AppShowStatusShow   AppShowStatus = 1
	AppShowStatusHidden AppShowStatus = 2
)

// Timestamp is an alias for utils.Timestamp for backward compatibility.
type Timestamp = utils.Timestamp

// FlexInt is an alias for utils.FlexInt for backward compatibility.
type FlexInt = utils.FlexInt

// -----------------------------------------------------------------------------
// Request Args Types (for functions with more than 2 parameters)
// -----------------------------------------------------------------------------

// QueryWalletAddressArgs represents arguments for querying wallet addresses.
type QueryWalletAddressArgs struct {
	WalletID int64  `json:"sub_wallet_id"` // Wallet ID (required)
	Symbol   string `json:"symbol"`        // Coin symbol (required)
	MaxID    int64  `json:"max_id"`        // Starting address ID, default 0
}

// SetAutoCollectSymbolArgs represents arguments for setting auto collect symbol.
type SetAutoCollectSymbolArgs struct {
	Symbol       string          `json:"symbol"`        // Coin symbol (required)
	CollectMin   decimal.Decimal `json:"collect_min"`   // Minimum amount for collection (required)
	FuelingLimit decimal.Decimal `json:"fueling_limit"` // Fueling limit (required)
}

// BuyTronResourceArgs represents arguments for buying Tron resource.
type BuyTronResourceArgs struct {
	WalletID     int64  `json:"sub_wallet_id"` // Wallet ID (required)
	Symbol       string `json:"symbol"`        // Resource symbol, e.g. TRX (required)
	ResourceType string `json:"resource_type"` // Resource type: ENERGY or BANDWIDTH (required)
	ResourceVal  int64  `json:"resource_val"`  // Resource amount (required)
	Duration     int    `json:"duration"`      // Duration in days (required)
}

// GetCoinDetailsArgs represents arguments for getting coin details.
type GetCoinDetailsArgs struct {
	Symbol          string `json:"symbol,omitempty"`           // Coin symbol (optional)
	ContractAddress string `json:"contract_address,omitempty"` // Contract address (optional)
	ShowBalance     bool   `json:"show_balance,omitempty"`     // Whether to show balance
	MaxID           int    `json:"max_id,omitempty"`           // Starting ID for pagination
	Limit           int    `json:"limit,omitempty"`            // Page size limit
}

// -----------------------------------------------------------------------------
// Withdraw Types
// -----------------------------------------------------------------------------

// WithdrawRequest represents a withdrawal request.
type WithdrawRequest struct {
	RequestID string          `json:"request_id"`
	WalletID  int64           `json:"sub_wallet_id"`
	Symbol    string          `json:"symbol"`
	Amount    decimal.Decimal `json:"amount"`
	AddressTo string          `json:"address_to"`
	From      string          `json:"from,omitempty"`
	Memo      string          `json:"memo,omitempty"`
	Remark    string          `json:"remark,omitempty"`
	Outputs   string          `json:"outputs,omitempty"`
	Sign      string          `json:"sign,omitempty"`
}

// WithdrawResponse represents a withdrawal response.
type WithdrawResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		WithdrawID int64 `json:"withdraw_id"`
	} `json:"data"`
}

// WithdrawRecord represents a withdrawal record.
type WithdrawRecord struct {
	ID              int64           `json:"id"`
	RequestID       string          `json:"request_id"`
	WalletID        int64           `json:"sub_wallet_id"`
	Symbol          string          `json:"symbol"`
	ContractAddress string          `json:"contract_address,omitempty"`
	BaseSymbol      string          `json:"base_symbol"`
	AddressFrom     string          `json:"address_from"`
	AddressTo       string          `json:"address_to"`
	Memo            string          `json:"memo,omitempty"`
	Amount          decimal.Decimal `json:"amount"`
	Txid            string          `json:"txid"`
	FeeSymbol       string          `json:"fee_symbol"`
	Fee             decimal.Decimal `json:"fee"`
	RealFee         decimal.Decimal `json:"real_fee"`
	Status          FlexInt         `json:"status"`
	Confirmations   FlexInt         `json:"confirmations"`
	TxHeight        FlexInt         `json:"tx_height"`
	WithdrawSource  int             `json:"withdraw_source"` // 1: app, 2: openapi, 3: web
	DelegateFee     string          `json:"delegate_fee,omitempty"`
	CreatedAt       Timestamp       `json:"created_at"`
	UpdatedAt       Timestamp       `json:"updated_at"`
	Outputs         []struct {
		Address    string `json:"address"`
		Amount     string `json:"amount"`
		Memo       string `json:"memo"`
		NeedActive bool   `json:"needActive"`
	} `json:"outputs"`
}

// WithdrawRecordResult represents withdraw records response.
type WithdrawRecordResult struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []*WithdrawRecord `json:"data"`
}

// -----------------------------------------------------------------------------
// Web3 Types
// -----------------------------------------------------------------------------

// Web3TransRequest represents a Web3 transaction request.
type Web3TransRequest struct {
	RequestID           string          `json:"request_id"`
	WalletID            int64           `json:"sub_wallet_id"`
	MainChainSymbol     string          `json:"main_chain_symbol"`
	InteractiveContract string          `json:"interactive_contract"`
	Amount              decimal.Decimal `json:"amount"`
	GasPrice            decimal.Decimal `json:"gas_price"`
	GasLimit            int64           `json:"gas_limit"`
	InputData           string          `json:"input_data"`
	TransType           string          `json:"trans_type"`
	From                string          `json:"from,omitempty"`
	DappName            string          `json:"dapp_name,omitempty"`
	DappURL             string          `json:"dapp_url,omitempty"`
	DappImg             string          `json:"dapp_img,omitempty"`
	Sign                string          `json:"sign,omitempty"`
}

// Web3TransResponse represents a Web3 transaction response.
type Web3TransResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TransID int64 `json:"trans_id"`
	} `json:"data"`
}

// Web3TransRecord represents a Web3 transaction record.
type Web3TransRecord struct {
	ID                  int64           `json:"id"`
	RequestID           string          `json:"request_id"`
	WalletID            int64           `json:"sub_wallet_id"`
	Txid                string          `json:"txid"`
	Symbol              string          `json:"symbol"`
	MainChainSymbol     string          `json:"main_chain_symbol"`
	Amount              decimal.Decimal `json:"amount"`
	FeeSymbol           string          `json:"fee_symbol"`
	Fee                 decimal.Decimal `json:"fee"`
	RealFee             decimal.Decimal `json:"real_fee"`
	From                string          `json:"from"`
	InteractiveContract string          `json:"interactive_contract"`
	GasPrice            decimal.Decimal `json:"gas_price"`
	GasLimit            FlexInt         `json:"gas_limit"`
	InputData           string          `json:"input_data"`
	TransType           FlexInt         `json:"trans_type"`
	DappName            string          `json:"dapp_name,omitempty"`
	DappURL             string          `json:"dapp_url,omitempty"`
	DappImg             string          `json:"dapp_img,omitempty"`
	Status              FlexInt         `json:"status"`
	Confirmations       FlexInt         `json:"confirmations"`
	TxHeight            FlexInt         `json:"tx_height"`
	CreatedAt           Timestamp       `json:"created_at"`
	UpdatedAt           Timestamp       `json:"updated_at"`
}

// Web3RecordResult represents Web3 transaction records response.
type Web3RecordResult struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []*Web3TransRecord `json:"data"`
}

// Web3AccelerationArgs represents Web3 transaction acceleration arguments.
// See: https://custodydocs-en.chainup.com/api-references/mpc-apis/apis/web3/web3-pending
type Web3AccelerationArgs struct {
	// TransID is the Web3 transaction ID (required)
	TransID int `json:"trans_id"`
	// GasPrice is the gas fee in Gwei (required)
	GasPrice string `json:"gas_price"`
	// GasLimit is the gas limit fee (required)
	GasLimit string `json:"gas_limit"`
}

// -----------------------------------------------------------------------------
// Wallet Types
// -----------------------------------------------------------------------------

// Wallet represents wallet information.
type Wallet struct {
	WalletID int64 `json:"sub_wallet_id"`
}

// WalletCreateResult represents wallet creation response.
type WalletCreateResult struct {
	Code string  `json:"code"`
	Msg  string  `json:"msg"`
	Data *Wallet `json:"data"`
}

// WalletAddress represents wallet address information.
type WalletAddress struct {
	ID       int64  `json:"id"`
	Address  string `json:"address"`
	AddrType int    `json:"addr_type"`
	Memo     string `json:"memo,omitempty"`
}

// WalletAddressResult represents wallet address response.
type WalletAddressResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data *WalletAddress `json:"data"`
}

// WalletAddressListResult represents wallet address list response.
type WalletAddressListResult struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []*WalletAddress `json:"data"`
}

// WalletAssets represents wallet assets.
type WalletAssets struct {
	NormalBalance     decimal.Decimal `json:"normal_balance"`
	LockBalance       decimal.Decimal `json:"lock_balance"`
	CollectingBalance decimal.Decimal `json:"collecting_balance"`
}

// WalletAssetsResult represents wallet assets response.
type WalletAssetsResult struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data *WalletAssets `json:"data"`
}

// WalletAddressInfo represents wallet address detailed information.
type WalletAddressInfo struct {
	WalletID           int64  `json:"sub_wallet_id"`
	AddrType           int    `json:"addr_type"`
	MergeAddressSymbol string `json:"merge_address_symbol"`
	Memo               string `json:"memo"`
}

// WalletAddressInfoResult represents wallet address info response.
type WalletAddressInfoResult struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data *WalletAddressInfo `json:"data"`
}

// -----------------------------------------------------------------------------
// Deposit Types
// -----------------------------------------------------------------------------

// DepositRecord represents a deposit record.
type DepositRecord struct {
	ID              int64           `json:"id"`
	WalletID        int64           `json:"sub_wallet_id"`
	Symbol          string          `json:"symbol"`
	BaseSymbol      string          `json:"base_symbol"`
	ContractAddress string          `json:"contract_address,omitempty"`
	Amount          decimal.Decimal `json:"amount"`
	AddressFrom     string          `json:"address_from"`
	AddressTo       string          `json:"address_to"`
	Memo            string          `json:"memo,omitempty"`
	Txid            string          `json:"txid"`
	Confirmations   FlexInt         `json:"confirmations"`
	TxHeight        FlexInt         `json:"tx_height"`
	Status          FlexInt         `json:"status"`
	DepositType     int             `json:"deposit_type"` // 1: Regular, 2: NFT
	TokenID         string          `json:"token_id,omitempty"`
	CreatedAt       Timestamp       `json:"created_at"`
	UpdatedAt       Timestamp       `json:"updated_at"`
	RefundAmount    string          `json:"refund_amount,omitempty"`
	KytStatus       string          `json:"kyt_status"`
}

// DepositRecordResult represents deposit records response.
type DepositRecordResult struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []*DepositRecord `json:"data"`
}

// -----------------------------------------------------------------------------
// Workspace Types
// -----------------------------------------------------------------------------

// SupportMainChain represents supported main chain information.
type SupportMainChain struct {
	CoinNet             string `json:"coin_net"`
	Symbol              string `json:"symbol"`
	IsSupportMemo       int    `json:"is_support_memo"`
	ChainID             string `json:"chain_id,omitempty"`
	EnableWithdraw      bool   `json:"enable_withdraw"`
	EnableDeposit       bool   `json:"enable_deposit"`
	SupportAcceleration bool   `json:"support_acceleration"`
	NeedPayment         bool   `json:"need_payment"`
	IfOpenChain         bool   `json:"if_open_chain,omitempty"`
	RealSymbol          string `json:"real_symbol"`
	SymbolAlias         string `json:"symbol_alias"`
	DisplayOrder        int    `json:"display_order,omitempty"`
}

// SupportMainChainData represents the data field of supported main chains response.
type SupportMainChainData struct {
	OpenMainChain    []*SupportMainChain `json:"open_main_chain"`
	SupportMainChain []*SupportMainChain `json:"support_main_chain"`
}

// SupportMainChainResult represents supported main chains response.
type SupportMainChainResult struct {
	Code string                `json:"code"`
	Msg  string                `json:"msg"`
	Data *SupportMainChainData `json:"data"`
}

// CoinDetails represents coin details information.
type CoinDetails struct {
	ID                   FlexInt `json:"id"`
	BaseSymbol           string  `json:"base_symbol"`
	CoinNet              string  `json:"coin_net"`
	Symbol               string  `json:"symbol"`
	SymbolAlias          string  `json:"symbol_alias"`
	AddressRegex         string  `json:"address_regex,omitempty"`
	AddressTagRegex      string  `json:"address_tag_regex,omitempty"`
	ContractAddress      string  `json:"contract_address,omitempty"`
	Decimals             FlexInt `json:"decimals"`
	DepositConfirmation  FlexInt `json:"deposit_confirmation"`
	AddressLink          string  `json:"address_link,omitempty"`
	TxidLink             string  `json:"txid_link,omitempty"`
	Icon                 string  `json:"icon,omitempty"`
	IfOpenChain          bool    `json:"if_open_chain"`
	RealSymbol           string  `json:"real_symbol"`
	SupportMemo          string  `json:"support_memo,omitempty"`
	SupportToken         string  `json:"support_token,omitempty"`
	SupportAcceleration  bool    `json:"support_acceleration"`
	SupportMultiAddr     bool    `json:"support_multi_addr"`
	MergeAddressSymbol   string  `json:"merge_address_symbol,omitempty"`
	CoinType             FlexInt `json:"coin_type"` // 0: account, 1: utxo, 2: memo
	MinWithdraw          string  `json:"min_withdraw,omitempty"`
	WithdrawConfirmation FlexInt `json:"withdraw_confirmation"`
}

// CoinDetailsResult represents coin details response.
type CoinDetailsResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []*CoinDetails `json:"data"`
}

// BlockHeight represents block height information.
type BlockHeight struct {
	BlockHeight int64 `json:"height"`
}

// BlockHeightResult represents block height response.
type BlockHeightResult struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data *BlockHeight `json:"data"`
}

// -----------------------------------------------------------------------------
// AutoSweep Types
// -----------------------------------------------------------------------------

// AutoCollectWallet represents auto collect wallet info.
type AutoCollectWallet struct {
	CollectWalletId int64  `json:"collect_sub_wallet_id"`
	FuelingWalletId int64  `json:"fueling_sub_wallet_id"`
	Symbol          string `json:"symbol"`
}

// AutoCollectResult represents auto collect response.
type AutoCollectResult struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data *AutoCollectWallet `json:"data"`
}

// AutoCollectRecord represents auto collect record.
// https://custodydocs-en.chainup.com/api-references/mpc-apis/apis/consolidation/consolidation-sync-list
type AutoCollectRecord struct {
	ID              int64           `json:"id"`
	WalletID        int64           `json:"sub_wallet_id"`
	Symbol          string          `json:"symbol"`
	Amount          decimal.Decimal `json:"amount"`
	FeeSymbol       string          `json:"fee_symbol"` // Fee currency, e.g.: ETH
	Fee             decimal.Decimal `json:"fee"`        // Fee amount
	RealFee         decimal.Decimal `json:"real_fee"`   // Actual consumed fee
	CreatedAt       Timestamp       `json:"created_at"`
	UpdatedAt       Timestamp       `json:"updated_at"`
	AddressFrom     string          `json:"address_from"`     // Sender's address
	AddressTo       string          `json:"address_to"`       // Consolidation address
	Txid            string          `json:"txid"`             // Transaction hash
	Confirmations   FlexInt         `json:"confirmations"`    // Number of block confirmations
	Status          FlexInt         `json:"status"`           // Consolidation status
	TransType       int             `json:"trans_type"`       // 10: Consolidation, 11: Consolidation Gas
	BaseSymbol      string          `json:"base_symbol"`      // Base currency on main chain
	ContractAddress string          `json:"contract_address"` // Contract address
	DelegateFee     decimal.Decimal `json:"delegate_fee"`     // TRON delegate fee
}

// AutoCollectRecordResult represents auto collect records response.
type AutoCollectRecordResult struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []*AutoCollectRecord `json:"data"`
}

// -----------------------------------------------------------------------------
// Notify Types
// -----------------------------------------------------------------------------

// NotifyData represents notification data from MPC async callback.
// https://custodydocs-zh.chainup.com/api-references/mpc-apis/notify
type NotifyData struct {
	// Common fields
	Side       string    `json:"side"`          // deposit or withdraw
	NotifyTime Timestamp `json:"notify_time"`   // Notification timestamp
	RequestID  string    `json:"request_id"`    // Request ID (withdraw/web3 only)
	ID         FlexInt   `json:"id"`            // Record ID
	WalletID   FlexInt   `json:"sub_wallet_id"` // Wallet ID

	// Token information
	Symbol          string          `json:"symbol"`                     // Coin symbol
	ContractAddress string          `json:"contract_address,omitempty"` // Contract address
	Amount          decimal.Decimal `json:"amount"`                     // Amount

	// Fee information
	FeeSymbol string          `json:"fee_symbol,omitempty"` // Fee symbol
	RealFee   decimal.Decimal `json:"real_fee,omitempty"`   // Actual fee

	// Address information
	AddressFrom string `json:"address_from,omitempty"` // From address
	AddressTo   string `json:"address_to,omitempty"`   // To address
	Memo        string `json:"memo,omitempty"`         // Memo/Tag

	// Transaction information
	Txid          string    `json:"txid,omitempty"`      // Transaction hash
	Confirmations FlexInt   `json:"confirmations"`       // Confirmation count
	Status        FlexInt   `json:"status"`              // Status
	TxHeight      FlexInt   `json:"tx_height,omitempty"` // Block height
	CreatedAt     Timestamp `json:"created_at"`          // Created time
	UpdatedAt     Timestamp `json:"updated_at"`          // Updated time

	// Chain information
	BaseSymbol string `json:"base_symbol,omitempty"` // Base chain symbol

	// Withdraw specific fields
	WithdrawSource int             `json:"withdraw_source,omitempty"` // Withdraw source: 1-app, 2-openapi, 3-web
	DelegateFee    decimal.Decimal `json:"delegate_fee,omitempty"`    // TRON delegate fee

	// Web3 specific fields
	MainChainSymbol     string `json:"main_chain_symbol,omitempty"`    // Web3 main chain symbol
	InteractiveContract string `json:"interactive_contract,omitempty"` // Web3 interactive contract
	InputData           string `json:"input_data,omitempty"`           // Web3 input data
	TransType           string `json:"trans_type,omitempty"`           // Web3 transaction type
	DappImg             string `json:"dapp_img,omitempty"`             // DApp icon URL
	DappName            string `json:"dapp_name,omitempty"`            // DApp name
	DappURL             string `json:"dapp_url,omitempty"`             // DApp URL
}

// -----------------------------------------------------------------------------
// Tron Resource Types
// -----------------------------------------------------------------------------

// TronBuyResourceArgs represents Tron resource purchase arguments.
type TronBuyResourceArgs struct {
	RequestID         string `json:"request_id"`
	BuyType           int    `json:"buy_type,omitempty"`         // Purchase type
	ResourceType      int    `json:"resource_type,omitempty"`    // Resource type: 0-energy, 1-bandwidth
	EnergyNum         int    `json:"energy_num,omitempty"`       // Energy amount
	NetNum            int    `json:"net_num,omitempty"`          // Bandwidth amount
	ServiceChargeType string `json:"service_charge_type"`        // Service charge type
	AddressFrom       string `json:"address_from"`               // From address
	AddressTo         string `json:"address_to,omitempty"`       // To address
	ContractAddress   string `json:"contract_address,omitempty"` // Contract address
}

// TronBuyResource represents Tron resource purchase result.
type TronBuyResource struct {
	TransID int64 `json:"trans_id"`
}

// TronBuyResourceResult represents Tron resource purchase response.
type TronBuyResourceResult struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data *TronBuyResource `json:"data"`
}

// TronBuyResourceRecord represents Tron resource purchase record.
type TronBuyResourceRecord struct {
	ID                int             `json:"id"`
	RequestID         string          `json:"request_id"`
	AddressFrom       string          `json:"address_from"`
	ServiceChargeRate decimal.Decimal `json:"service_charge_rate"`
	ServiceCharge     string          `json:"service_charge"`
	ContractAddress   string          `json:"contract_address"`
	AddressTo         string          `json:"address_to"`
	ResourceType      FlexInt         `json:"resource_type"`
	BuyType           FlexInt         `json:"buy_type"`
	NetNum            FlexInt         `json:"net_num"`
	EnergyNum         FlexInt         `json:"energy_num"`
	NetTxid           string          `json:"net_txid"`
	EnergyTxid        string          `json:"energy_txid"`
	ReclaimNetTxid    string          `json:"reclaim_net_txid"`
	ReclaimEnergyTxid string          `json:"reclaim_energy_txid"`
	NetTime           Timestamp       `json:"net_time"`
	EnergyTime        Timestamp       `json:"energy_time"`
	ReclaimNetTime    Timestamp       `json:"reclaim_net_time"`
	ReclaimEnergyTime Timestamp       `json:"reclaim_energy_time"`
	NetPrice          decimal.Decimal `json:"net_price"`
	EnergyPrice       decimal.Decimal `json:"energy_price"`
	Status            FlexInt         `json:"status"`
}

// TronBuyResourceRecordResult represents Tron resource purchase records response.
type TronBuyResourceRecordResult struct {
	Code string                   `json:"code"`
	Msg  string                   `json:"msg"`
	Data []*TronBuyResourceRecord `json:"data"`
}

// TronResourceOrder represents Tron resource order (legacy).
type TronResourceOrder struct {
	OrderID     string          `json:"order_id"`
	PayAmount   decimal.Decimal `json:"pay_amount"`
	PaySymbol   string          `json:"pay_symbol"`
	Duration    int             `json:"duration"`
	ResourceVal int64           `json:"resource_val"`
}

// TronResourceOrderResult represents Tron resource order response (legacy).
type TronResourceOrderResult struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data *TronResourceOrder `json:"data"`
}
