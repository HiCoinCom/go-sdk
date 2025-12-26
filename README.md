# ChainUp Custody Go SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.21-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

ChainUp Custody 官方 Go SDK - 为数字资产托管提供完整的解决方案。

[English Documentation](./README_EN.md)

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
        SetAppID("your-app-id").
        SetPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetPublicKey("-----BEGIN PUBLIC KEY-----\n...").
        SetDebug(true).
        Build()
    if err != nil {
        log.Fatal(err)
    }

    // 用户注册
    userAPI := client.GetUserAPI()
    userResult, err := userAPI.RegisterEmailUser("user@example.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("用户注册成功: UID=%d\n", userResult.Data.UID)

    // 账户余额查询
    accountAPI := client.GetAccountAPI()
    account, err := accountAPI.GetUserAccount(12345, "BTC")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("余额: %s\n", account.Data.NormalBalance.String())

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
    fmt.Printf("提币ID: %d\n", withdrawResult.Data.ID)
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
        SetAppID("your-app-id").
        SetRsaPrivateKey("-----BEGIN PRIVATE KEY-----\n...").
        SetWaasPublicKey("-----BEGIN PUBLIC KEY-----\n...").
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
    fmt.Printf("钱包创建成功: ID=%d\n", walletResult.Data.WalletID)

    // 创建地址
    addressResult, err := walletAPI.CreateWalletAddress(walletResult.Data.WalletID, "ETH")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("地址: %s\n", addressResult.Data.Address)

    // 提币（带交易签名）
    withdrawAPI := client.GetWithdrawAPI()
    withdrawReq := &types.WithdrawRequest{
        RequestID: "unique-request-id",
        WalletID:  walletResult.Data.WalletID,
        Symbol:    "ETH",
        Amount:    decimal.NewFromFloat(0.1),
        AddressTo: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0",
    }
    withdrawResult, err := withdrawAPI.Withdraw(withdrawReq, true) // true = 启用交易签名
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("提币成功: OrderID=%s\n", withdrawResult.OrderID)

    // 创建 Web3 交易
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
    fmt.Printf("Web3 交易创建成功: OrderID=%s\n", web3Result.OrderID)
}
```

## 📁 项目结构

```
go-sdk/
├── chainup.go              # 主入口点
├── custody/                # WaaS API 模块
│   ├── client.go           # WaaS 客户端
│   ├── config.go           # WaaS 配置
│   ├── api/                # API 实现
│   │   ├── base.go         # 基础 API
│   │   ├── user.go         # 用户管理
│   │   ├── account.go      # 账户管理
│   │   ├── billing.go      # 充提管理
│   │   ├── coin.go         # 币种信息
│   │   ├── transfer.go     # 转账操作
│   │   └── async_notify.go # 异步通知
│   └── types/              # 类型定义
│       └── types.go
├── mpc/                    # MPC API 模块
│   ├── client.go           # MPC 客户端
│   ├── config.go           # MPC 配置
│   ├── api/                # API 实现
│   │   ├── client.go       # 基础客户端
│   │   ├── config.go       # 配置接口
│   │   ├── errors.go       # 错误类型
│   │   ├── response.go     # 响应处理
│   │   ├── wallet.go       # 钱包管理
│   │   ├── deposit.go      # 充值记录
│   │   ├── withdraw.go     # 提币操作
│   │   ├── web3.go         # Web3 交易
│   │   ├── auto_sweep.go   # 自动归集
│   │   ├── notify.go       # 通知处理
│   │   ├── workspace.go    # 工作区信息
│   │   └── tron_resource.go# TRON 资源
│   └── types/              # 类型定义
│       └── types.go
├── utils/                  # 工具包
│   ├── constants.go        # 常量定义
│   ├── crypto.go           # RSA 加解密
│   ├── http_client.go      # HTTP 客户端
│   └── mpcsign/            # MPC 签名
│       └── mpcsign.go
└── examples/               # 示例代码
    ├── waas_example.go
    └── mpc_example.go
```

## 🔧 API 参考

### WaaS API

| API 类       | 方法                                  | 说明         |
| ------------ | ------------------------------------- | ------------ |
| UserAPI      | `RegisterMobileUser(country, mobile)` | 手机注册     |
| UserAPI      | `RegisterEmailUser(email)`            | 邮箱注册     |
| UserAPI      | `GetMobileUser(country, mobile)`      | 根据手机获取用户 |
| UserAPI      | `GetEmailUser(email)`                 | 根据邮箱获取用户 |
| UserAPI      | `SyncUserList(maxID)`                 | 同步用户列表 |
| AccountAPI   | `GetUserAccount(uid, symbol)`         | 获取账户余额 |
| AccountAPI   | `GetUserAddress(uid, symbol)`         | 获取充值地址 |
| AccountAPI   | `GetCompanyAccount(symbol)`           | 获取商户账户 |
| AccountAPI   | `GetUserAddressInfo(address)`         | 获取地址信息 |
| AccountAPI   | `SyncUserAddressList(maxID)`          | 同步地址列表 |
| BillingAPI   | `Withdraw(args)`                      | 发起提币     |
| BillingAPI   | `WithdrawList(requestIDs)`            | 获取提币记录 |
| BillingAPI   | `SyncWithdrawList(maxID)`             | 同步提币记录 |
| BillingAPI   | `DepositList(ids)`                    | 获取充值记录 |
| BillingAPI   | `SyncDepositList(maxID)`              | 同步充值记录 |
| BillingAPI   | `MinerFeeList(ids)`                   | 获取矿工费记录 |
| BillingAPI   | `SyncMinerFeeList(maxID)`             | 同步矿工费记录 |
| CoinAPI      | `GetCoinList()`                       | 获取币种列表 |
| TransferAPI  | `AccountTransfer(args)`               | 内部转账     |
| TransferAPI  | `GetAccountTransferList(requestIDs)`  | 获取转账记录 |
| TransferAPI  | `SyncAccountTransferList(maxID)`      | 同步转账记录 |
| AsyncNotifyAPI | `NotifyRequest(req)`                | 解密异步通知 |

### MPC API

| API 类           | 方法                                      | 说明             |
| ---------------- | ----------------------------------------- | ---------------- |
| WalletAPI        | `CreateWallet(name, status)`              | 创建钱包         |
| WalletAPI        | `CreateWalletAddress(walletID, symbol)`   | 创建地址         |
| WalletAPI        | `QueryWalletAddress(args)`                | 查询钱包地址     |
| WalletAPI        | `GetWalletAssets(walletID, symbol)`       | 获取钱包资产     |
| WalletAPI        | `ChangeWalletShowStatus(ids, status)`     | 修改钱包显示状态 |
| WalletAPI        | `WalletAddressInfo(address, memo)`        | 获取地址信息     |
| DepositAPI       | `GetDepositRecords(ids)`                  | 获取充值记录     |
| DepositAPI       | `SyncDepositRecords(maxID)`               | 同步充值记录     |
| WithdrawAPI      | `Withdraw(req, needSign)`                 | 发起提币         |
| WithdrawAPI      | `GetWithdrawRecords(requestIDs)`          | 获取提币记录     |
| WithdrawAPI      | `SyncWithdrawRecords(maxID)`              | 同步提币记录     |
| Web3API          | `CreateWeb3Trans(req, needSign)`          | 创建 Web3 交易   |
| Web3API          | `AccelerationWeb3Trans(args)`             | 加速 Web3 交易   |
| Web3API          | `GetWeb3Records(requestIDs)`              | 获取 Web3 记录   |
| Web3API          | `SyncWeb3Records(maxID)`                  | 同步 Web3 记录   |
| AutoSweepAPI     | `AutoCollectSubWallets(ids, symbol)`      | 自动归集         |
| AutoSweepAPI     | `SetAutoCollectSymbol(args)`              | 设置自动归集币种 |
| AutoSweepAPI     | `SyncAutoCollectRecords(maxID)`           | 同步归集记录     |
| WorkSpaceAPI     | `GetSupportMainChain()`                   | 获取支持的主链   |
| WorkSpaceAPI     | `GetCoinDetails(args)`                    | 获取币种详情     |
| WorkSpaceAPI     | `GetLastBlockHeight(mainChainSymbol)`     | 获取最新区块高度 |
| TronResourceAPI  | `CreateTronDelegate(args)`                | 创建 TRON 代理   |
| TronResourceAPI  | `GetBuyResourceRecords(requestIDs)`       | 获取资源购买记录 |
| TronResourceAPI  | `SyncBuyResourceRecords(maxID)`           | 同步资源购买记录 |
| NotifyAPI        | `NotifyRequest(req)`                      | 解密异步通知     |

## 📋 类型定义

### MPC 类型 (`mpc/types`)

```go
// 提币请求
type WithdrawRequest struct {
    RequestID string          `json:"request_id"`
    WalletID  int64           `json:"sub_wallet_id"`
    Symbol    string          `json:"symbol"`
    Amount    decimal.Decimal `json:"amount"`
    AddressTo string          `json:"address_to"`
    Memo      string          `json:"memo,omitempty"`
    Remark    string          `json:"remark,omitempty"`
}

// Web3 交易请求
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

// 钱包显示状态
type AppShowStatus int
const (
    AppShowStatusShow   AppShowStatus = 1  // 显示
    AppShowStatusHidden AppShowStatus = 2  // 隐藏
)
```

### WaaS 类型 (`custody/types`)

```go
// 用户信息
type UserInfo struct {
    UID      FlexInt `json:"uid"`
    Nickname string  `json:"nickname"`
}

// 账户信息
type Account struct {
    DepositAddress string          `json:"deposit_address"`
    LockBalance    decimal.Decimal `json:"lock_balance"`
    NormalBalance  decimal.Decimal `json:"normal_balance"`
}
```

## 🔐 安全说明

1. **私钥安全**: 请妥善保管 RSA 私钥，不要提交到版本控制
2. **加密通信**: 所有 API 请求使用 RSA 加密
3. **签名验证**: 交易签名使用 SHA256 算法

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 支持

- 官网: https://custody.chainup.com
- 技术支持: custody@chainup.com
