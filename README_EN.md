# ChainUp Custody Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.21-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Official ChainUp Custody Go SDK - Complete solution for digital asset custody.

[中文文档](./README.md)

## ✨ Features

- 🔐 **WaaS (Wallet-as-a-Service)** - Complete custody wallet API integration
- 🔑 **MPC (Multi-Party Computation)** - Secure distributed key management
- 🏗️ **Modern Architecture** - Object-oriented design with Builder pattern
- 📝 **Complete Documentation** - Detailed code comments and examples
- ✅ **Production Ready** - Follows Go official specifications
- 🚀 **Easy Integration** - Simple and intuitive API
- 🔒 **Secure & Reliable** - RSA encrypted communication

## 📦 Installation

```bash
go get chainup.com/go-sdk
```

## 🚀 Quick Start

### WaaS (Custody) API

```go
package main

import (
    "fmt"
    "log"

    "chainup.com/go-sdk/custody"
)

func main() {
    // Create WaaS client using Builder pattern
    client, err := custody.NewWaasClientBuilder().
        SetHost("https://api.custody.chainup.com").
        SetAppID("your-app-id").
        SetPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetPublicKey("-----BEGIN PUBLIC KEY-----\n...").
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // User operations
    userAPI := client.GetUserAPI()
    user, err := userAPI.RegisterEmailUser(map[string]interface{}{
        "email": "user@example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User registered: %+v\n", user)

    // Account operations
    accountAPI := client.GetAccountAPI()
    account, err := accountAPI.GetUserAccount(map[string]interface{}{
        "uid": 12345,
        "symbol": "BTC",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Account: %+v\n", account)
}
```

### MPC Wallet API

```go
package main

import (
    "fmt"
    "log"

    "chainup.com/go-sdk/mpc"
)

func main() {
    // Create MPC client
    client, err := mpc.NewMpcClientBuilder().
        SetAppID("your-app-id").
        SetRsaPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetApiKey("your-api-key").
        SetDomain("https://mpc-api.custody.chainup.com").
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // Create wallet
    walletAPI := client.GetWalletAPI()
    wallet, err := walletAPI.CreateWallet(map[string]interface{}{
        "sub_wallet_name": "My Wallet",
        "app_show_status": 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Wallet created: %+v\n", wallet)

    // Create address
    address, err := walletAPI.CreateWalletAddress(map[string]interface{}{
        "sub_wallet_id": wallet["sub_wallet_id"],
        "symbol": "ETH",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Address created: %+v\n", address)
}
```

## 📚 API Documentation

### WaaS API

- **UserAPI** - User management and registration

  - `RegisterMobileUser` - Register user with mobile
  - `RegisterEmailUser` - Register user with email
  - `GetMobileUser` - Get user info by mobile
  - `GetEmailUser` - Get user info by email
  - `SyncUserList` - Sync user list

- **AccountAPI** - Account and balance management

  - `GetUserAccount` - Get user account balance
  - `GetUserAddress` - Get user deposit address
  - `GetCompanyAccount` - Get merchant account balance
  - `GetUserAddressInfo` - Query address info
  - `SyncUserAddressList` - Sync user address list

- **BillingAPI** - Bills and transaction records

  - `Withdraw` - Withdraw
  - `WithdrawList` - Withdrawal list
  - `SyncWithdrawList` - Sync withdrawal list
  - `DepositList` - Deposit list
  - `SyncDepositList` - Sync deposit list
  - `MinerFeeList` - Miner fee list
  - `SyncMinerFeeList` - Sync miner fee list

- **CoinAPI** - Coin information

  - `GetCoinList` - Get coin list

- **TransferAPI** - Transfer operations

  - `AccountTransfer` - Account transfer
  - `GetAccountTransferList` - Get transfer list
  - `SyncAccountTransferList` - Sync transfer list

- **AsyncNotifyAPI** - Async notifications
  - `NotifyRequest` - Decrypt async notification

### MPC API

- **WalletAPI** - Wallet management

  - `CreateWallet` - Create wallet
  - `CreateWalletAddress` - Create wallet address
  - `QueryWalletAddress` - Query wallet addresses
  - `GetWalletAssets` - Get wallet assets
  - `ChangeWalletShowStatus` - Change wallet show status
  - `WalletAddressInfo` - Get address info

- **DepositAPI** - Deposit records

  - `GetDepositRecords` - Get deposit records
  - `SyncDepositRecords` - Sync deposit records

- **WithdrawAPI** - Withdrawal operations

  - `Withdraw` - Withdraw
  - `GetWithdrawRecords` - Get withdrawal records
  - `SyncWithdrawRecords` - Sync withdrawal records

- **Web3API** - Web3 transactions

  - `CreateWeb3Trans` - Create Web3 transaction
  - `AccelerationWeb3Trans` - Accelerate Web3 transaction
  - `GetWeb3Records` - Get Web3 records
  - `SyncWeb3Records` - Sync Web3 records

- **AutoSweepAPI** - Auto-sweep

  - `AutoCollectSubWallets` - Collect from sub-wallets
  - `SetAutoCollectSymbol` - Set auto-collect symbol
  - `SyncAutoCollectRecords` - Sync auto-collect records

- **WorkSpaceAPI** - Workspace

  - `GetSupportMainChain` - Get supported main chains
  - `GetCoinDetails` - Get coin details
  - `GetLastBlockHeight` - Get latest block height

- **TronResourceAPI** - TRON resources

  - `CreateTronDelegate` - Create TRON delegate
  - `GetBuyResourceRecords` - Get buy resource records
  - `SyncBuyResourceRecords` - Sync buy resource records

- **NotifyAPI** - Async notifications
  - `NotifyRequest` - Decrypt async notification

## 🔐 Security Notes

1. **Private Key Security**: Keep your RSA private key secure, do not commit to version control
2. **Encrypted Communication**: All API requests use RSA encryption
3. **Public Key Verification**: Response data is decrypted with ChainUp public key

## 📝 License

[MIT License](LICENSE)

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📧 Contact

- Website: https://custody.chainup.com
- Support: support@chainup.com
