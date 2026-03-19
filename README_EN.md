# ChainUp Custody Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.21-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Official ChainUp Custody Go SDK - Complete solution for digital asset custody.

[中文文档](./README.md)

## ✨ Features

- 🔐 **WaaS (Wallet-as-a-Service)** - Complete custody wallet API integration
- 🔑 **MPC (Multi-Party Computation)** - Secure distributed key management
- 🏗️ **Modern Architecture** - Object-oriented design with Builder pattern
- 📝 **Type Safe** - All APIs use strongly-typed structs
- ✅ **Precise Amounts** - Uses `decimal.Decimal` for financial data
- 🚀 **Easy Integration** - Simple and intuitive API
- 🔒 **Secure & Reliable** - RSA encrypted communication

## 📦 Installation

```bash
[chainup.com/go-sdk-demo](https://github.com/HiCoinCom/go-sdk-demo)
```

## 🚀 Quick Start

### WaaS (Custody) API

```go
package main

import (
    "fmt"
    "log"

    "chainup.com/go-sdk/custody"
    "chainup.com/go-sdk/custody/api"
    "github.com/shopspring/decimal"
)

func main() {
    // Create WaaS client using Builder pattern
    client, err := custody.NewWaasClientBuilder().
        SetAppID("your-app-id").
        SetPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetPublicKey("-----BEGIN PUBLIC KEY-----\n...").
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // User registration
    userAPI := client.GetUserAPI()
    userResult, err := userAPI.RegisterEmailUser("user@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User registered: UID=%d\n", userResult.Data.UID)

    // Get account balance
    accountAPI := client.GetAccountAPI()
    account, err := accountAPI.GetUserAccount(12345, "BTC")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance: %s\n", account.Data.NormalBalance.String())

    // Withdraw request
    billingAPI := client.GetBillingAPI()
    withdrawResult, err := billingAPI.Withdraw(&api.WithdrawArgs{
        RequestID: "unique-request-id",
        FromUID:   12345,
        ToAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
        Amount:    decimal.NewFromFloat(0.1),
        Symbol:    "BTC",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Withdraw ID: %d\n", withdrawResult.Data.ID)
}
```

### MPC Wallet API

```go
package main

import (
    "fmt"
    "log"

    "chainup.com/go-sdk/mpc"
    "chainup.com/go-sdk/mpc/types"
    "github.com/shopspring/decimal"
)

func main() {
    // Create MPC client
    client, err := mpc.NewMpcClientBuilder().
        SetAppID("your-app-id").
        SetRsaPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetWaasPublicKey("-----BEGIN PUBLIC KEY-----\n...").
        SetSignPrivateKey("-----BEGIN PRIVATE KEY-----\n..."). // Optional: for transaction signing
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // Create wallet
    walletAPI := client.GetWalletAPI()
    walletResult, err := walletAPI.CreateWallet("My Wallet", types.AppShowStatusShow)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Wallet created: ID=%d\n", walletResult.Data.WalletID)

    // Create address
    addressResult, err := walletAPI.CreateWalletAddress(walletResult.Data.WalletID, "ETH")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Address: %s\n", addressResult.Data.Address)

    // Withdraw with transaction signing
    withdrawAPI := client.GetWithdrawAPI()
    withdrawReq := &types.WithdrawRequest{
        RequestID: "unique-request-id",
        WalletID:  walletResult.Data.WalletID,
        Symbol:    "ETH",
        Amount:    decimal.NewFromFloat(0.1),
        AddressTo: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0",
    }
    withdrawResult, err := withdrawAPI.Withdraw(withdrawReq, true) // true = enable transaction signing
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Withdraw success: OrderID=%s\n", withdrawResult.OrderID)

    // Create Web3 transaction
    web3API := client.GetWeb3API()
    web3Req := &types.Web3TransRequest{
        RequestID:           "web3-request-id",
        WalletID:            walletResult.Data.WalletID,
        MainChainSymbol:     "ETH",
        InteractiveContract: "0xabcdef...",
        Amount:              decimal.NewFromFloat(0.1),
        GasPrice:            decimal.NewFromInt(20000000000), // 20 Gwei
        GasLimit:            21000,
        InputData:           "0x",
        TransType:           "1",
    }
    web3Result, err := web3API.CreateWeb3Trans(web3Req, true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Web3 transaction created: OrderID=%s\n", web3Result.OrderID)
}
```

## 📁 Project Structure

```
go-sdk/
├── chainup.go              # Main entry point
├── custody/                # WaaS API module
│   ├── client.go           # WaaS client
│   ├── config.go           # WaaS configuration
│   ├── api/                # API implementations
│   │   ├── base.go         # Base API
│   │   ├── user.go         # User management
│   │   ├── account.go      # Account management
│   │   ├── billing.go      # Deposit/Withdraw management
│   │   ├── coin.go         # Coin information
│   │   ├── transfer.go     # Transfer operations
│   │   └── async_notify.go # Async notifications
│   └── types/              # Type definitions
│       └── types.go
├── mpc/                    # MPC API module
│   ├── client.go           # MPC client
│   ├── config.go           # MPC configuration
│   ├── api/                # API implementations
│   │   ├── client.go       # Base client
│   │   ├── config.go       # Config interface
│   │   ├── errors.go       # Error types
│   │   ├── response.go     # Response handling
│   │   ├── wallet.go       # Wallet management
│   │   ├── deposit.go      # Deposit records
│   │   ├── withdraw.go     # Withdraw operations
│   │   ├── web3.go         # Web3 transactions
│   │   ├── auto_sweep.go   # Auto sweep
│   │   ├── notify.go       # Notification handling
│   │   ├── workspace.go    # Workspace info
│   │   └── tron_resource.go# TRON resources
│   └── types/              # Type definitions
│       └── types.go
├── utils/                  # Utilities
│   ├── constants.go        # Constants
│   ├── crypto.go           # RSA encryption
│   ├── http_client.go      # HTTP client
│   └── mpcsign/            # MPC signing
│       └── mpcsign.go
└── examples/               # Example code
    ├── waas_example.go
    └── mpc_example.go
```

## 🔧 API Reference

### WaaS API

| API Class      | Method                                | Description                |
| -------------- | ------------------------------------- | -------------------------- |
| UserAPI        | `RegisterMobileUser(country, mobile)` | Register by mobile         |
| UserAPI        | `RegisterEmailUser(email)`            | Register by email          |
| UserAPI        | `GetMobileUser(country, mobile)`      | Get user by mobile         |
| UserAPI        | `GetEmailUser(email)`                 | Get user by email          |
| UserAPI        | `SyncUserList(maxID)`                 | Sync user list             |
| AccountAPI     | `GetUserAccount(uid, symbol)`         | Get account balance        |
| AccountAPI     | `GetUserAddress(uid, symbol)`         | Get deposit address        |
| AccountAPI     | `GetCompanyAccount(symbol)`           | Get company account        |
| AccountAPI     | `GetUserAddressInfo(address)`         | Get address info           |
| AccountAPI     | `SyncUserAddressList(maxID)`          | Sync address list          |
| BillingAPI     | `Withdraw(args)`                      | Initiate withdraw          |
| BillingAPI     | `WithdrawList(requestIDs)`            | Get withdraw records       |
| BillingAPI     | `SyncWithdrawList(maxID)`             | Sync withdraw records      |
| BillingAPI     | `DepositList(ids)`                    | Get deposit records        |
| BillingAPI     | `SyncDepositList(maxID)`              | Sync deposit records       |
| BillingAPI     | `MinerFeeList(ids)`                   | Get miner fee records      |
| BillingAPI     | `SyncMinerFeeList(maxID)`             | Sync miner fee records     |
| CoinAPI        | `GetCoinList()`                       | Get coin list              |
| TransferAPI    | `AccountTransfer(args)`               | Internal transfer          |
| TransferAPI    | `GetAccountTransferList(requestIDs)`  | Get transfer records       |
| TransferAPI    | `SyncAccountTransferList(maxID)`      | Sync transfer records      |
| AsyncNotifyAPI | `NotifyRequest(req)`                  | Decrypt async notification |

### MPC API

| API Class       | Method                                  | Description                    |
| --------------- | --------------------------------------- | ------------------------------ |
| WalletAPI       | `CreateWallet(name, status)`            | Create wallet                  |
| WalletAPI       | `CreateWalletAddress(walletID, symbol)` | Create address                 |
| WalletAPI       | `QueryWalletAddress(args)`              | Query wallet addresses         |
| WalletAPI       | `GetWalletAssets(walletID, symbol)`     | Get wallet assets              |
| WalletAPI       | `ChangeWalletShowStatus(ids, status)`   | Change wallet visibility       |
| WalletAPI       | `WalletAddressInfo(address, memo)`      | Get address info               |
| DepositAPI      | `GetDepositRecords(ids)`                | Get deposit records            |
| DepositAPI      | `SyncDepositRecords(maxID)`             | Sync deposit records           |
| WithdrawAPI     | `Withdraw(req, needSign)`               | Initiate withdraw              |
| WithdrawAPI     | `GetWithdrawRecords(requestIDs)`        | Get withdraw records           |
| WithdrawAPI     | `SyncWithdrawRecords(maxID)`            | Sync withdraw records          |
| Web3API         | `CreateWeb3Trans(req, needSign)`        | Create Web3 transaction        |
| Web3API         | `AccelerationWeb3Trans(args)`           | Accelerate Web3 transaction    |
| Web3API         | `GetWeb3Records(requestIDs)`            | Get Web3 records               |
| Web3API         | `SyncWeb3Records(maxID)`                | Sync Web3 records              |
| AutoSweepAPI    | `AutoCollectSubWallets(ids, symbol)`    | Auto collect                   |
| AutoSweepAPI    | `SetAutoCollectSymbol(args)`            | Set auto collect symbol        |
| AutoSweepAPI    | `SyncAutoCollectRecords(maxID)`         | Sync collect records           |
| WorkSpaceAPI    | `GetSupportMainChain()`                 | Get supported main chains      |
| WorkSpaceAPI    | `GetCoinDetails(args)`                  | Get coin details               |
| WorkSpaceAPI    | `GetLastBlockHeight(mainChainSymbol)`   | Get latest block height        |
| TronResourceAPI | `CreateTronDelegate(args)`              | Create TRON delegate           |
| TronResourceAPI | `GetBuyResourceRecords(requestIDs)`     | Get resource purchase records  |
| TronResourceAPI | `SyncBuyResourceRecords(maxID)`         | Sync resource purchase records |
| NotifyAPI       | `NotifyRequest(req)`                    | Decrypt async notification     |

## 📋 Type Definitions

### MPC Types (`mpc/types`)

```go
// Withdraw request
type WithdrawRequest struct {
    RequestID string          `json:"request_id"`
    WalletID  int64           `json:"sub_wallet_id"`
    Symbol    string          `json:"symbol"`
    Amount    decimal.Decimal `json:"amount"`
    AddressTo string          `json:"address_to"`
    Memo      string          `json:"memo,omitempty"`
    Remark    string          `json:"remark,omitempty"`
}

// Web3 transaction request
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
}

// Web3 transaction acceleration arguments
// Doc: https://custodydocs-en.chainup.com/api-references/mpc-apis/apis/web3/web3-pending
type Web3AccelerationArgs struct {
    TransID  int    `json:"trans_id"`   // Web3 transaction ID (required)
    GasPrice string `json:"gas_price"`  // Gas fee in Gwei (required)
    GasLimit string `json:"gas_limit"`  // Gas limit (required)
}

// Wallet display status
type AppShowStatus int
const (
    AppShowStatusShow   AppShowStatus = 1  // Show
    AppShowStatusHidden AppShowStatus = 2  // Hidden
)
```

### WaaS Types (`custody/types`)

```go
// User information
type UserInfo struct {
    UID      FlexInt `json:"uid"`
    Nickname string  `json:"nickname"`
}

// Account information
type Account struct {
    DepositAddress string          `json:"deposit_address"`
    LockBalance    decimal.Decimal `json:"lock_balance"`
    NormalBalance  decimal.Decimal `json:"normal_balance"`
}
```

## 🔐 Security Notes

1. **Private Key Security**: Keep your RSA private key secure, do not commit to version control
2. **Encrypted Communication**: All API requests use RSA encryption
3. **Signature Verification**: Transaction signing uses SHA256 algorithm

## 📄 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📞 Support

- Website: https://custody.chainup.com
- Support: custody@chainup.com
