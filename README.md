# ChainUp Custody Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.21-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

ChainUp Custody 官方 Go SDK - 为数字资产托管提供完整的解决方案。

## ✨ 特性

- 🔐 **WaaS（钱包即服务）** - 完整的托管钱包 API 集成
- 🔑 **MPC（多方计算）** - 安全的分布式密钥管理
- 🏗️ **现代架构** - 面向对象设计，使用 Builder 模式
- 📝 **类型安全** - 所有 API 使用强类型结构体
- ✅ **精确金额** - 使用 `decimal.Decimal` 处理金融数据
- 🚀 **易于集成** - 简单直观的 API
- 🔒 **安全可靠** - RSA 加密通信

## 📦 安装

```bash
go get chainup.com/go-sdk
```

## 🚀 快速开始

### WaaS（托管）API

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
    // 使用 Builder 模式创建 WaaS 客户端
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

    // 用户注册 - 使用强类型参数
    userAPI := client.GetUserAPI()
    userResult, err := userAPI.RegisterEmailUser("user@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User registered: UID=%d\n", userResult.Data.UID)

    // 账户余额查询
    accountAPI := client.GetAccountAPI()
    account, err := accountAPI.GetUserAccount(12345, "BTC")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Balance: %s %s\n", account.Data.Balance.String(), account.Data.Symbol)

    // 提币请求
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

### MPC 钱包 API

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
    // 创建 MPC 客户端
    client, err := mpc.NewMpcClientBuilder().
        SetDomain("https://mpc-api.custody.chainup.com").
        SetAppID("your-app-id").
        SetRsaPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetWaasPublicKey("-----BEGIN PUBLIC KEY-----\n...").
        SetApiKey("your-api-key").
        SetSignPrivateKey("-----BEGIN PRIVATE KEY-----\n..."). // 可选：用于交易签名
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // 创建钱包
    walletAPI := client.GetWalletAPI()
    walletResult, err := walletAPI.CreateWallet("My Wallet", types.AppShowStatusShow)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Wallet created: ID=%d\n", walletResult.Data.SubWalletID)

    // 创建地址
    addressResult, err := walletAPI.CreateWalletAddress(walletResult.Data.SubWalletID, "ETH")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Address: %s\n", addressResult.Data.Address)

    // 提币（带交易签名）
    withdrawAPI := client.GetWithdrawAPI()
    withdrawReq := &types.WithdrawRequest{
        RequestID:   "unique-request-id",
        SubWalletID: walletResult.Data.SubWalletID,
        Symbol:      "ETH",
        Amount:      decimal.NewFromFloat(0.1),
        AddressTo:   "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0",
    }
    withdrawResult, err := withdrawAPI.Withdraw(withdrawReq, true) // true = 启用交易签名
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Withdraw ID: %d\n", withdrawResult.Data.WithdrawID)

    // 创建 Web3 交易
    web3API := client.GetWeb3API()
    web3Req := &types.Web3TransRequest{
        RequestID:           "web3-request-id",
        SubWalletID:         walletResult.Data.SubWalletID,
        MainChainSymbol:     "ETH",
        InteractiveContract: "0xabcdef...",
        Amount:              decimal.NewFromFloat(0.1),
        GasPrice:            decimal.NewFromInt(20000000000), // 20 Gwei
        GasLimit:            21000,
        InputData:           "0x",
        TransType:           "1",
    }
    web3Result, err := web3API.CreateWeb3Trans(web3Req, false)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Web3 Trans ID: %d\n", web3Result.Data.TransID)
}
```

## 📋 类型定义

SDK 使用强类型结构体，所有时间字段使用 `time.Time`，金额字段使用 `decimal.Decimal`：

### MPC 类型 (`mpc/types`)

```go
// 提币请求
type WithdrawRequest struct {
    RequestID   string          `json:"request_id"`
    SubWalletID int64           `json:"sub_wallet_id"`
    Symbol      string          `json:"symbol"`
    Amount      decimal.Decimal `json:"amount"`
    AddressTo   string          `json:"address_to"`
    // ...
}

// 提币记录
type WithdrawRecord struct {
    ID          int64           `json:"id"`
    Amount      decimal.Decimal `json:"amount"`
    CreatedAt   Timestamp       `json:"created_at"`    // time.Time 包装类型
    ConfirmedAt Timestamp       `json:"confirmed_at"`  // time.Time 包装类型
    // ...
}

// Web3 交易请求
type Web3TransRequest struct {
    GasLimit int64           `json:"gas_limit"` // int64 类型
    GasPrice decimal.Decimal `json:"gas_price"`
    // ...
}
```

### WaaS 类型 (`custody/types`)

```go
// 提币信息
type Withdraw struct {
    Amount    decimal.Decimal `json:"amount"`
    CreatedAt Timestamp       `json:"created_at"` // time.Time 包装类型
    UpdatedAt Timestamp       `json:"updated_at"` // time.Time 包装类型
    // ...
}

// 充值信息
type Deposit struct {
    Amount    decimal.Decimal `json:"amount"`
    CreatedAt Timestamp       `json:"created_at"` // time.Time 包装类型
    // ...
}
```

## 📁 项目结构

```
go-sdk/
├── chainup.go           # 主入口点和类型别名
├── custody/             # WaaS API 模块
│   ├── client.go        # WaaS 客户端
│   ├── config.go        # WaaS 配置
│   ├── api/             # API 实现
│   │   ├── user.go      # 用户管理
│   │   ├── account.go   # 账户管理
│   │   ├── billing.go   # 充提管理
│   │   ├── coin.go      # 币种信息
│   │   ├── transfer.go  # 转账操作
│   │   └── async_notify.go # 异步通知
│   └── types/           # 类型定义
│       └── types.go
├── mpc/                 # MPC API 模块
│   ├── client.go        # MPC 客户端
│   ├── config.go        # MPC 配置
│   ├── api/             # API 实现
│   │   ├── wallet.go    # 钱包管理
│   │   ├── deposit.go   # 充值记录
│   │   ├── withdraw.go  # 提币操作
│   │   ├── web3.go      # Web3 交易
│   │   ├── auto_sweep.go # 自动归集
│   │   ├── notify.go    # 通知处理
│   │   ├── workspace.go # 工作区信息
│   │   └── tron_resource.go # TRON 资源
│   └── types/           # 类型定义
│       └── types.go
├── utils/               # 工具包
│   ├── constants.go     # 常量定义
│   ├── crypto.go        # RSA 加解密
│   ├── http_client.go   # HTTP 客户端
│   └── mpcsign/         # MPC 签名
│       └── mpcsign.go
└── examples/            # 示例代码
    ├── waas_example.go
    └── mpc_example.go
```

## 🔧 API 参考

### WaaS API

| API         | 方法                                  | 说明         |
| ----------- | ------------------------------------- | ------------ |
| UserAPI     | `RegisterMobileUser(country, mobile)` | 手机注册     |
| UserAPI     | `RegisterEmailUser(email)`            | 邮箱注册     |
| AccountAPI  | `GetUserAccount(uid, symbol)`         | 获取账户余额 |
| AccountAPI  | `GetUserAddress(uid, symbol)`         | 获取充值地址 |
| BillingAPI  | `Withdraw(args)`                      | 发起提币     |
| BillingAPI  | `DepositList(ids)`                    | 获取充值记录 |
| TransferAPI | `AccountTransfer(args)`               | 内部转账     |

### MPC API

| API          | 方法                                    | 说明           |
| ------------ | --------------------------------------- | -------------- |
| WalletAPI    | `CreateWallet(name, status)`            | 创建钱包       |
| WalletAPI    | `CreateWalletAddress(walletID, symbol)` | 创建地址       |
| WithdrawAPI  | `Withdraw(req, needSign)`               | 发起提币       |
| Web3API      | `CreateWeb3Trans(req, needSign)`        | 创建 Web3 交易 |
| DepositAPI   | `GetDepositRecords(ids)`                | 获取充值记录   |
| AutoSweepAPI | `AutoCollectSubWallets(ids, symbol)`    | 自动归集       |

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 支持

如有问题，请联系 ChainUp 技术支持或提交 GitHub Issue。
