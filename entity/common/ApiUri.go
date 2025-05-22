package common

type ApiUri string

const (
	MpcWalletCoin          ApiUri = "/mpc/wallet/open_coin"
	MpcCoinList            ApiUri = "/mpc/coin_list"
	ChainHeight            ApiUri = "/mpc/chain_height"
	CreateSubWallet        ApiUri = "/mpc/sub_wallet/create"
	CreateSubWalletAddress ApiUri = "/mpc/sub_wallet/create/address"
	GetSubWalletAddress    ApiUri = "/mpc/sub_wallet/get/address/list"
	GetSubWalletAssets     ApiUri = "/mpc/sub_wallet/assets"
	SubWalletChangeShow    ApiUri = "/mpc/sub_wallet/change_show_status"
	SubWalletAddressInfo   ApiUri = "/mpc/sub_wallet/address/info"

	SubWalletBillWithdraw     ApiUri = "/api/mpc/billing/withdraw"
	SubWalletWithdrawList     ApiUri = "/api/mpc/billing/withdraw_list"
	SubWalletSyncWithdrawList ApiUri = "/api/mpc/billing/sync_withdraw_list"

	DepositList     ApiUri = "/api/mpc/billing/deposit_list"
	SyncDepositList ApiUri = "/api/mpc/billing/sync_deposit_list"

	GetCollectSubWallet  ApiUri = "/api/mpc/auto_collect/sub_wallets"
	SetAutoCollectSymbol ApiUri = "/api/mpc/billing/auto_collect/symbol/set"
	GetAutoCollectList   ApiUri = "/api/mpc/billing/sync_auto_collect_list"

	Web3TransCreate   ApiUri = "/api/mpc/web3/trans/create"
	Web3TransSpeed    ApiUri = "/api/mpc/web3/pending"
	Web3TransList     ApiUri = "/api/mpc/web3/trans_list"
	SyncWeb3TransList ApiUri = "/api/mpc/web3/sync_trans_list"
)
