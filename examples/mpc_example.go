// Package main demonstrates MPC API usage
// To run: go run examples/mpc_example.go

package main

import (
	"fmt"
	"log"

	"chainup.com/go-sdk/mpc"
	"chainup.com/go-sdk/mpc/types"
	"github.com/shopspring/decimal"
)

func main() {
	// Create MPC client using Builder pattern
	client, err := mpc.NewMpcClientBuilder().
		SetAppID("").
		SetRsaPrivateKey(``).
		SetWaasPublicKey(``).
		SetSignPrivateKey(``).
		SetDebug(false).
		Build()

	if err != nil {
		log.Fatalf("Failed to create MPC client: %v", err)
	}

	// Example 1: Create a wallet
	fmt.Println("=== Example 1: Create Wallet ===")
	walletAPI := client.GetWalletAPI()
	walletResult, err := walletAPI.CreateWallet("My Wallet1", types.AppShowStatusShow)
	if err != nil {
		log.Printf("Failed to create wallet: %v", err)
	} else if walletResult.Data != nil {
		fmt.Printf("Wallet created: Code=%s, Msg=%s, WalletID=%d\n",
			walletResult.Code, walletResult.Msg, walletResult.Data.WalletID)
	} else {
		fmt.Printf("Wallet response: Code=%s, Msg=%s\n", walletResult.Code, walletResult.Msg)
	}

	// Example 2: Create wallet address
	fmt.Println("\n=== Example 2: Create Wallet Address ===")
	addressResult, err := walletAPI.CreateWalletAddress(1000537, "ETH")
	if err != nil {
		log.Printf("Failed to create address: %v", err)
	} else if addressResult.Data != nil {
		fmt.Printf("Address created: Address=%s, AddrType=%d\n",
			addressResult.Data.Address, addressResult.Data.AddrType)
	} else {
		fmt.Printf("Address response: Code=%s, Msg=%s\n", addressResult.Code, addressResult.Msg)
	}

	// Example 3: Query wallet addresses
	fmt.Println("\n=== Example 3: Query Wallet Addresses ===")
	addressList, err := walletAPI.QueryWalletAddress(&types.QueryWalletAddressArgs{
		WalletID: 1000537,
		Symbol:   "ETH",
		MaxID:    0,
	})
	if err != nil {
		log.Printf("Failed to query addresses: %v", err)
	} else {
		fmt.Printf("Found %d addresses\n", len(addressList.Data))
		for _, addr := range addressList.Data {
			fmt.Printf("  - ID=%d, Address=%s, AddrType=%d\n", addr.ID, addr.Address, addr.AddrType)
			break
		}
	}

	// Example 4: Get wallet assets
	fmt.Println("\n=== Example 4: Get Wallet Assets ===")
	assetsResult, err := walletAPI.GetWalletAssets(1000537, "APTOS")
	if err != nil {
		log.Printf("Failed to get assets: %v", err)
	} else if assetsResult.Data != nil {
		fmt.Printf("Assets: Symbol=%s, Balance=%s, Frozen=%s\n",
			assetsResult.Data.Symbol, assetsResult.Data.Balance.String(), assetsResult.Data.FrozenAmount.String())
	} else {
		fmt.Printf("Assets response: Code=%s, Msg=%s\n", assetsResult.Code, assetsResult.Msg)
	}

	// Example 5: Change wallet show status
	fmt.Println("\n=== Example 5: Change Wallet Show Status ===")
	success, err := walletAPI.ChangeWalletShowStatus([]int64{1000537, 123457}, types.AppShowStatusShow)
	if err != nil {
		log.Printf("Failed to change status: %v", err)
	} else {
		fmt.Printf("Change status result: %v\n", success)
	}

	// walletaddressinfo
	addressInfo, err := walletAPI.WalletAddressInfo("0x633A84Ee0ab29d911e5466e5E1CB9cdBf5917E72", "")
	if err != nil {
		log.Printf("Failed to WalletAddressInfo: %v", err)
	} else {
		fmt.Printf("WalletAddressInfo result: %v %+v, %+v\n",
			addressInfo.Data.WalletID, addressInfo.Data.AddrType, addressInfo.Data.MergeAddressSymbol)
	}

	// Example 6: Withdraw (with structured types)
	fmt.Println("\n=== Example 6: Withdraw ===")
	withdrawAPI := client.GetWithdrawAPI()
	// Example 7: Withdraw with transaction signature
	fmt.Println("\n=== Example 7: Withdraw with Signature ===")
	withdrawReqWithSign := &types.WithdrawRequest{
		RequestID: "12345678949",
		WalletID:  1000537,
		Symbol:    "Sepolia",
		Amount:    decimal.NewFromFloat(0.001),
		AddressTo: "0xdcb0D867403adE76e75a4A6bBcE9D53C9d05B981",
		Remark:    "Signed withdrawal",
	}

	// Pass true to enable transaction signing
	withdrawResult2, err := withdrawAPI.Withdraw(withdrawReqWithSign, true)
	if err != nil {
		log.Printf("Failed to withdraw with signature: %v", err)
	} else if withdrawResult2.Data.WithdrawID != 0 {
		fmt.Printf("Withdraw with signature result: Code=%s, WithdrawID=%d\n",
			withdrawResult2.Code, withdrawResult2.Data.WithdrawID)
	} else {
		fmt.Printf("Withdraw with signature response: Code=%s, Msg=%s\n", withdrawResult2.Code, withdrawResult2.Msg)
	}

	// Example 8: Get withdraw records
	fmt.Println("\n=== Example 8: Get Withdraw Records ===")
	withdrawRecords, err := withdrawAPI.GetWithdrawRecords([]string{"12345678901", "12345678"})
	if err != nil {
		log.Printf("Failed to get withdraw records: %v", err)
	} else {
		fmt.Printf("Found %d withdraw records\n", len(withdrawRecords.Data))
	}

	// Example 9: Sync withdraw records
	fmt.Println("\n=== Example 9: Sync Withdraw Records ===")
	syncRecords, err := withdrawAPI.SyncWithdrawRecords(0)
	if err != nil {
		log.Printf("Failed to sync withdraw records: %v", err)
	} else {
		fmt.Printf("Synced %d withdraw records\n", len(syncRecords.Data))
	}

	// Example 10: Get deposit records
	fmt.Println("\n=== Example 10: Get Deposit Records ===")
	depositAPI := client.GetDepositAPI()
	depositRecords, err := depositAPI.GetDepositRecords([]int64{3294170, 456, 3})
	if err != nil {
		log.Printf("Failed to get deposit records: %v", err)
	} else {
		fmt.Printf("Found %d deposit records\n", len(depositRecords.Data))
	}

	depositRecords, err = depositAPI.SyncDepositRecords(100)
	if err != nil {
		log.Printf("Failed to sync deposit records: %v", err)
	} else {
		fmt.Printf("sync %d deposit records\n", len(depositRecords.Data))
	}

	// Example 11: Create Web3 Transaction
	fmt.Println("\n=== Example 11: Create Web3 Transaction ===")
	web3API := client.GetWeb3API()
	web3Req := &types.Web3TransRequest{
		RequestID:           "1234567890",
		WalletID:            123456,
		MainChainSymbol:     "ETH",
		InteractiveContract: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
		Amount:              decimal.NewFromFloat(10.1),
		GasPrice:            decimal.NewFromInt(20000000000), // 20 Gwei
		GasLimit:            21000,
		InputData:           "0x",
		TransType:           "1",
		DappName:            "Example DApp",
	}

	web3Result, err := web3API.CreateWeb3Trans(web3Req, true)
	if err != nil {
		log.Printf("Failed to create Web3 transaction: %v", err)
	} else if web3Result.Data.TransID != 0 {
		fmt.Printf("Web3 transaction created: Code=%s, TransID=%d\n",
			web3Result.Code, web3Result.Data.TransID)
	} else {
		fmt.Printf("Web3 transaction response: Code=%s, Msg=%s\n", web3Result.Code, web3Result.Msg)
	}

	web3Record, err := web3API.GetWeb3Records([]string{"123456578901"})
	if err != nil {
		log.Printf("Failed to get web3 record: %v", err)
	} else {
		fmt.Printf("Found %d web3 record\n", len(web3Record.Data))
	}

	web3Record, err = web3API.SyncWeb3Records(1)
	if err != nil {
		log.Printf("Failed to sync web3 record: %v", err)
	} else {
		fmt.Printf("Sync %d web3 record\n", len(web3Record.Data))
	}

	// Example 12: Get supported main chains
	fmt.Println("\n=== Example 12: Get Supported Main Chains ===")
	workspaceAPI := client.GetWorkSpaceAPI()
	chains, err := workspaceAPI.GetSupportMainChain()
	if err != nil {
		log.Printf("Failed to get supported chains: %v", err)
	} else if chains.Data != nil {
		fmt.Printf("Found %d open main chains, %d support main chains\n",
			len(chains.Data.OpenMainChain), len(chains.Data.SupportMainChain))
	} else {
		fmt.Println("No supported chain data returned")
	}

	coinList, err := workspaceAPI.GetCoinDetails(&types.GetCoinDetailsArgs{
		MaxID: 1,
		Limit: 100,
	})
	if err != nil {
		log.Printf("Failed to get coin list: %v", err)
	} else {
		fmt.Printf("Found %d coin list\n", len(coinList.Data))
		for _, v := range coinList.Data {
			log.Printf("%s %s %s", v.Symbol, v.BaseSymbol, v.ContractAddress)
			break
		}
	}

	// Example 13: Get block height
	fmt.Println("\n=== Example 13: Get Block Height ===")
	blockHeight, err := workspaceAPI.GetLastBlockHeight("ETH")
	if err != nil {
		log.Printf("Failed to get block height: %v", err)
	} else if blockHeight.Data != nil {
		fmt.Printf("ETH block height: %d\n", blockHeight.Data.BlockHeight)
	} else {
		fmt.Printf("Block height response: Code=%s, Msg=%s\n", blockHeight.Code, blockHeight.Msg)
	}

	// Example 14: Auto collect
	fmt.Println("\n=== Example 14: Auto Collect ===")
	autoSweepAPI := client.GetAutoSweepAPI()
	collectResult, err := autoSweepAPI.AutoCollectSubWallets([]int64{1000537, 123457}, "ETH")
	if err != nil {
		log.Printf("Failed to auto collect: %v", err)
	} else {
		fmt.Printf("Auto collect result: Code=%s, Msg=%s %v %v\n",
			collectResult.Code, collectResult.Msg,
			collectResult.Data.CollectWalletId, collectResult.Data.FuelingWalletId)
	}

	setResult, err := autoSweepAPI.SetAutoCollectSymbol(&types.SetAutoCollectSymbolArgs{
		Symbol:       "USDTERC20",
		CollectMin:   decimal.New(10, 0),
		FuelingLimit: decimal.New(10, 0),
	})

	if err != nil {
		log.Printf("Failed to set auto collect: %v", err)
	} else {
		fmt.Printf("set auto collect result: %v\n", setResult)
	}

	sweepRecord, err := autoSweepAPI.SyncAutoCollectRecords(0)
	if err != nil {
		log.Printf("Failed to get collect: %v", err)
	} else {
		fmt.Printf("get collect record: %v\n", len(sweepRecord.Data))
	}

	// Example 15: Tron Resource API
	fmt.Println("\n=== Example 15: Tron Resource API ===")
	tronAPI := client.GetTronResourceAPI()

	// 15.1 Create Tron delegate (buy resource)
	fmt.Println("--- 15.1 Create Tron Delegate ---")
	delegateArgs := &types.TronBuyResourceArgs{
		RequestID:         "tron_delegate_test_001",
		BuyType:           0,
		ResourceType:      1, // 0: energy
		EnergyNum:         100000,
		ServiceChargeType: "10010",
		AddressFrom:       "TPjJg9FnzQuYBd6bshgaq7rkH4s36zju5S",
		AddressTo:         "TGmBzYfBBtMfFF8v9PweTaPwn3WoB7aGPd",
		ContractAddress:   "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	}
	delegateResult, err := tronAPI.CreateTronDelegate(delegateArgs)
	if err != nil {
		log.Printf("Failed to create Tron delegate: %v", err)
	} else if delegateResult.Data != nil {
		fmt.Printf("Tron delegate created: Code=%s, TransID=%d\n",
			delegateResult.Code, delegateResult.Data.TransID)
	} else {
		fmt.Printf("Tron delegate response: Code=%s, Msg=%s\n", delegateResult.Code, delegateResult.Msg)
	}

	// 15.2 Get buy resource records
	fmt.Println("--- 15.2 Get Buy Resource Records ---")
	resourceRecords, err := tronAPI.GetBuyResourceRecords([]string{"1234567890", "tron_test_002"})
	if err != nil {
		log.Printf("Failed to get resource records: %v", err)
	} else {
		fmt.Printf("Found %d resource records\n", len(resourceRecords.Data))
		for _, record := range resourceRecords.Data {
			fmt.Printf("  - RequestID=%s, Status=%d, EnergyNum=%d\n",
				record.RequestID, record.Status, record.EnergyNum)
			break
		}
	}

	// 15.3 Sync buy resource records
	fmt.Println("--- 15.3 Sync Buy Resource Records ---")
	syncResourceRecords, err := tronAPI.SyncBuyResourceRecords(0)
	if err != nil {
		log.Printf("Failed to sync resource records: %v", err)
	} else {
		fmt.Printf("Synced %d resource records\n", len(syncResourceRecords.Data))
	}

	// Example 16: NotifyRequest - Decrypt async notification
	fmt.Println("\n=== Example 16: NotifyRequest (Decrypt Async Notification) ===")
	notifyAPI := client.GetNotifyAPI()

	// This is a mock encrypted cipher that would be received from the callback
	// In real usage, this cipher comes from WaaS platform's async notification
	mockCipher := "Af-uUJj8a2-Og7E5CwzANv4vo8NMf-z-DijwrIuK74Or8eRveM7G_-f0ErtX4WurcVrjdWC-tqU0BDhBwiDijbdyCFBvYB5UmLnHL_Rg13amhQTM-kaHoh-U9WPhYB3vGRwWkTwJ_aETERVVciAvoTf5CalqydMSe8G3KNz-ymrSVUe92DfW5ZdDKJm1hNYYteGJvg0hk--GRiPybPv2W78NlTLyWmXq094megsVzZv-KlsEGPUvPoBnEJ0Xu__AO-l-GfCG4rVO4rb8J01Nq_0Q9eRKcKWq0ci7MfnPPLMhtAWwRvSd3U8PUNHOLqGaJzOLraFnuFUHn90h7T23_DeAduA2W6dto99qb8YQ_iVnMnOKfE0Ls7Vv5S2qhgQJ0nl-BA3PPPOwW37cMb-wTbi3ZezU_S1NQEbrruEChkPhTaK0AqsM6mESV8wGflcWx3N9XPv6QatJ9zedBnkfJ4bJ4Vy2rUEtQF8eVc6zXhV8PuDRiSMf0V0yxzMjE6o9z0s087KSAqFphitlHvQMPJ29FUnyvCe_Czr5WPuhl89GOZjERE2uoNTfHqAlZVzMamoPv4y0qyIjJTufAQm-WwrQK9kGesky7eCiOXVdtR9UhEYpzEJSgXxENjUrHMx6D2AlEzlr17a2DgI-WrWB7oUnyiNnf__ElmLPPkJBdFUfzJByQkLxkUB0FLvTWdVbiIRPmPpdgb7jkhJsHUSOH0NmULqu8bYiEQtGfqRJh8I98qDzHWwfE_VAbqwATj2oD959Fm1eInBqh7eXGoy2WR3o00VpPrNvoE4eJNmw3WpVzlRF7ZVwOpcWRT-dHTShz9mB2Etk9P8D4rGmMZyXHkt4aGUJkE1b3cOEjzkOEFX8CaNe-VHiBYhIyFzMetn7mfIFB0hl565FGEumbhDKNNz_m9T2qPM5k4BQ9fLWUt_WJAVdC81_piIlBOQfYPDbdYoc_9ser1p-Jy5cgTyOMdWuSWC3jMsT09xr8dMcLkKmd39khGidAvGqOOPL1ST0"

	notifyData, err := notifyAPI.NotifyRequest(mockCipher)
	if err != nil {
		log.Printf("Failed to decrypt notification (expected with mock data): %v", err)
	} else {
		fmt.Printf("Notification decrypted successfully:\n")
		fmt.Printf("  - Side: %s\n", notifyData.Side)
		fmt.Printf("  - ID: %d\n", notifyData.ID)
		fmt.Printf("  - WalletID: %d\n", notifyData.WalletID)
		fmt.Printf("  - Symbol: %s\n", notifyData.Symbol)
		fmt.Printf("  - Amount: %s\n", notifyData.Amount.String())
		fmt.Printf("  - AddressFrom: %s\n", notifyData.AddressFrom)
		fmt.Printf("  - AddressTo: %s\n", notifyData.AddressTo)
		fmt.Printf("  - Txid: %s\n", notifyData.Txid)
		fmt.Printf("  - Status: %d\n", notifyData.Status)
		fmt.Printf("  - Confirmations: %d\n", notifyData.Confirmations)
	}

	// Demonstrate how to use NotifyRequest in a real HTTP handler
	fmt.Println("\n--- NotifyRequest Usage Example ---")
	fmt.Println(``)

	fmt.Println("\n=== All examples completed ===")
}
