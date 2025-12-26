// Package main demonstrates WaaS API usage with typed requests and responses
// To run: go run examples/waas_example.go
//go:build ignore
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
		SetAppID("").
		SetPrivateKey(``).
		SetPublicKey(``).
		SetDebug(false).
		Build()

	if err != nil {
		log.Fatalf("Failed to create WaaS client: %v", err)
	}

	uid := int64(15036904) // 替换为实际用户ID
	// ============================================================
	// User API Examples
	// ============================================================

	// Example 1: Register a user by email
	fmt.Println("=== Example 1: Register User by Email ===")
	userAPI := client.GetUserAPI()
	userResult, err := userAPI.RegisterEmailUser("user21@example.com")
	if err != nil {
		log.Printf("Failed to register user: %v", err)
	} else {
		fmt.Printf("User registered - Code: %s, UID: %d\n", userResult.Code, userResult.Data.UID)
	}

	// Example 2: Register a user by mobile
	fmt.Println("\n=== Example 2: Register User by Mobile ===")
	mobileUser, err := userAPI.RegisterMobileUser("86", "13800138000")
	if err != nil {
		log.Printf("Failed to register mobile user: %v", err)
	} else {
		fmt.Printf("Mobile user registered - Code: %s, UID: %d\n", mobileUser.Code, mobileUser.Data.UID)
	}

	// Example 3: Get user by email
	fmt.Println("\n=== Example 3: Get User by Email ===")
	emailUser, err := userAPI.GetEmailUser("user@example.com")
	if err != nil {
		log.Printf("Failed to get email user: %v", err)
	} else {
		fmt.Printf("Email user - UID: %d, Nickname: %s\n", emailUser.Data.UID, emailUser.Data.Nickname)
	}

	// Example 4: Get user by mobile
	fmt.Println("\n=== Example 4: Get User by Mobile ===")
	mobileUserInfo, err := userAPI.GetMobileUser("86", "13800138000")
	if err != nil {
		log.Printf("Failed to get mobile user: %v", err)
	} else {
		fmt.Printf("Mobile user - UID: %d, Nickname: %s\n", mobileUserInfo.Data.UID, mobileUserInfo.Data.Nickname)
	}

	// Example 5: Sync user list
	fmt.Println("\n=== Example 5: Sync User List ===")
	userList, err := userAPI.SyncUserList(0) // 0 for first sync
	if err != nil {
		log.Printf("Failed to sync user list: %v", err)
	} else {
		fmt.Printf("Synced %d users\n", len(userList.Data))
		for _, u := range userList.Data {
			fmt.Printf("  - UID: %d, Nickname: %s\n", u.UID, u.Nickname)
			break
		}
	}

	// ============================================================
	// Account API Examples
	// ============================================================

	// Example 6: Get user account balance
	fmt.Println("\n=== Example 6: Get User Account Balance ===")
	accountAPI := client.GetAccountAPI()
	account, err := accountAPI.GetUserAccount(uid, "APTOS")
	if err != nil {
		log.Printf("Failed to get account: %v", err)
	} else {
		fmt.Printf("Account - Symbol: %s, Balance: %s\n", account.Data.NormalBalance, account.Data.LockBalance)
	}

	// Example 7: Get user deposit address
	fmt.Println("\n=== Example 7: Get User Deposit Address ===")
	address, err := accountAPI.GetUserAddress(uid, "APTOS")
	if err != nil {
		log.Printf("Failed to get address: %v", err)
	} else {
		fmt.Printf("Deposit address - Uid: %s, Address: %s\n", address.Data.UID, address.Data.Address)
	}

	// Example 8: Get company account
	fmt.Println("\n=== Example 8: Get Company Account ===")
	companyAccount, err := accountAPI.GetCompanyAccount("APTOS")
	if err != nil {
		log.Printf("Failed to get company account: %v", err)
	} else {
		fmt.Printf("Company account - Symbol: %s, Balance: %s\n", companyAccount.Data.Symbol, companyAccount.Data.Balance.String())
	}

	// Example 9: Get user address info
	fmt.Println("\n=== Example 9: Get User Address Info ===")
	addressInfo, err := accountAPI.GetUserAddressInfo("0xd4036730fd450237b8fea382bd887c4c96a8453a")
	if err != nil {
		log.Printf("Failed to get address info: %v", err)
	} else {
		fmt.Printf("Address info - id: %d, Symbol: %s\n", addressInfo.Data.Id, addressInfo.Data.Symbol)
	}

	// Example 10: Sync user address list
	fmt.Println("\n=== Example 10: Sync User Address List ===")
	addressList, err := accountAPI.SyncUserAddressList(0)
	if err != nil {
		log.Printf("Failed to sync address list: %v", err)
	} else {
		fmt.Printf("Synced %d addresses\n", len(addressList.Data))
	}

	// ============================================================
	// Coin API Examples
	// ============================================================

	// Example 11: Get coin list
	fmt.Println("\n=== Example 11: Get Coin List ===")
	coinAPI := client.GetCoinAPI()
	coins, err := coinAPI.GetCoinList()
	if err != nil {
		log.Printf("Failed to get coin list: %v", err)
	} else {
		fmt.Printf("Supported coins: %d\n", len(coins.Data))
		for _, coin := range coins.Data {
			fmt.Printf("  - Symbol: %s, Base symbol: %s, Decimals: %d\n", coin.Symbol, coin.BaseSymbol, coin.Decimals)
			break
		}
	}

	// ============================================================
	// Transfer API Examples
	// ============================================================

	// Example 12: Account transfer
	fmt.Println("\n=== Example 12: Account Transfer ===")
	transferAPI := client.GetTransferAPI()
	transferArgs := &api.TransferArgs{
		RequestID: "unique-transfer-id-123",
		FromUID:   12345,
		ToUID:     67890,
		Symbol:    "USDT",
		Amount:    decimal.NewFromFloat(100.5),
		Remark:    "Transfer example",
	}
	transfer, err := transferAPI.AccountTransfer(transferArgs)
	if err != nil {
		log.Printf("Failed to transfer: %v", err)
	} else {
		fmt.Printf("Transfer result - Code: %s, ID: %d\n", transfer.Code, transfer.Data.ID)
	}

	// Example 13: Get transfer list
	fmt.Println("\n=== Example 13: Get Transfer List ===")
	transferList, err := transferAPI.GetAccountTransferList([]string{"unique-transfer-id-123"})
	if err != nil {
		log.Printf("Failed to get transfer list: %v", err)
	} else {
		fmt.Printf("Transfer records: %d\n", len(transferList.Data))
		for _, t := range transferList.Data {
			fmt.Printf("  - RequestID: %s, Symbol: %s, Amount: %s\n", t.RequestID, t.Symbol, t.Amount.String())
			break
		}
	}

	// Example 14: Sync transfer list
	fmt.Println("\n=== Example 14: Sync Transfer List ===")
	syncTransferList, err := transferAPI.SyncAccountTransferList(0)
	if err != nil {
		log.Printf("Failed to sync transfer list: %v", err)
	} else {
		fmt.Printf("Synced transfer records: %d\n", len(syncTransferList.Data))
	}

	// ============================================================
	// Billing API Examples
	// ============================================================

	// Example 15: Withdraw
	fmt.Println("\n=== Example 15: Withdraw ===")
	billingAPI := client.GetBillingAPI()
	withdrawArgs := &api.WithdrawArgs{
		RequestID: "1234567803",
		FromUID:   uid,
		ToAddress: "0x0f1dc222af5ea2660ff84ae91adc48f1cb2d4991f1e6569dd24d94599c335a06",
		Amount:    decimal.NewFromFloat(0.001),
		Symbol:    "APtOS",
	}
	withdraw, err := billingAPI.Withdraw(withdrawArgs)
	if err != nil {
		log.Printf("Failed to withdraw: %v", err)
	} else {
		fmt.Printf("Withdraw result - Code: %s, ID: %d\n", withdraw.Code, withdraw.Data.ID)
	}

	// Example 16: Get withdrawal list
	fmt.Println("\n=== Example 16: Get Withdrawal List ===")
	withdrawList, err := billingAPI.WithdrawList([]string{"12345678"})
	if err != nil {
		log.Printf("Failed to get withdrawal list: %v", err)
	} else {
		fmt.Printf("Withdrawal records: %d\n", len(withdrawList.Data))
		for _, w := range withdrawList.Data {
			fmt.Printf("  - RequestID: %s, Symbol: %s, Amount: %s, Status: %d\n", w.RequestID, w.Symbol, w.Amount.String(), w.Status)
			break
		}
	}

	// Example 17: Sync withdrawal list
	fmt.Println("\n=== Example 17: Sync Withdrawal List ===")
	syncWithdrawList, err := billingAPI.SyncWithdrawList(0)
	if err != nil {
		log.Printf("Failed to sync withdrawal list: %v", err)
	} else {
		fmt.Printf("Synced withdrawal records: %d\n", len(syncWithdrawList.Data))
	}

	// Example 18: Get deposit list
	fmt.Println("\n=== Example 18: Get Deposit List ===")
	depositList, err := billingAPI.DepositList([]int64{123, 456, 3294170})
	if err != nil {
		log.Printf("Failed to get deposit list: %v", err)
	} else {
		fmt.Printf("Deposit records: %d\n", len(depositList.Data))
		for _, d := range depositList.Data {
			fmt.Printf("  - ID: %d, Symbol: %s, Amount: %s, Status: %d\n", d.ID, d.Symbol, d.Amount.String(), d.Status)
			break
		}
	}

	// Example 19: Sync deposit list
	fmt.Println("\n=== Example 19: Sync Deposit List ===")
	syncDepositList, err := billingAPI.SyncDepositList(0)
	if err != nil {
		log.Printf("Failed to sync deposit list: %v", err)
	} else {
		fmt.Printf("Synced deposit records: %d\n", len(syncDepositList.Data))
		for _, d := range syncDepositList.Data {
			fmt.Printf("  - ID: %d, Symbol: %s, Amount: %s, Status: %d\n", d.ID, d.Symbol, d.Amount.String(), d.Status)
			break
		}
	}

	// Example 20: Get miner fee list
	fmt.Println("\n=== Example 20: Get Miner Fee List ===")
	minerFeeList, err := billingAPI.MinerFeeList([]int64{1001, 1002})
	if err != nil {
		log.Printf("Failed to get miner fee list: %v", err)
	} else {
		fmt.Printf("Miner fee records: %d\n", len(minerFeeList.Data))
		for _, m := range minerFeeList.Data {
			fmt.Printf("  - ID: %d, Symbol: %s, Fee: %s\n", m.ID, m.Symbol, m.Fee.String())
			break
		}
	}

	// Example 21: Sync miner fee list
	fmt.Println("\n=== Example 21: Sync Miner Fee List ===")
	syncMinerFeeList, err := billingAPI.SyncMinerFeeList(0)
	if err != nil {
		log.Printf("Failed to sync miner fee list: %v", err)
	} else {
		fmt.Printf("Synced miner fee records: %d\n", len(syncMinerFeeList.Data))
	}

	// ============================================================
	// Async Notification Example
	// ============================================================

	// Example 22: Handle async notification
	fmt.Println("\n=== Example 22: Handle Async Notification ===")
	asyncNotifyAPI := client.GetAsyncNotifyAPI()
	// This cipher would come from a webhook callback
	cipher := "jhoA9MtGotqWxqEtB27SwCtJCo9JSIxh2B6m8CItrPQj2gsm6rw-ti1qY5tNP52qXg60FLK49cFj-a84m-57z8aT-Vo-YyJPTcM8Qpuyjj5Pf8tAcbBjBHganULYNPjCCkzgH5n5dlMZIp0tmpc7nV7Pp6hi63KjGGNTfAAbWp7QOVukAsQeQyBFPeKhlVEhq8xqQEN2yg_T1jHRUjIdlTDn2LG_i2tI0MlDpPg5FHL6cViSVM23WBPhJnAFOOrGhaqq06YtVG2m8_x_pLTyI5ZK61Bv0HnDUuIkDuRqNXyhko0sG9uGuKWJ3maWfUc9bSb0VcWPHeWnYUrcE2M9TVtwTEKdcImqZnvjc12YUh_Oz2a9VNls_XN_gTRbeIiTUGsiXX1Yq6OkCCxrsCgD0AXz0KOX4uphZldXq17ZO7sU21-b1y0rsk0qY6PbKRYpp4hhdeKpEfB2gckhf1rc9h17j0ufri4LqsE4EccGuQD4JcSrT5RLY4QRil4wdIO9ZPmhb-Od3zqT9OYPSvPg0QVCVpw-Tn17WfsZw2xB9gO8uzvGcvz9TfUrI8zKg6b6roTR9xt0m0oqMCyhrjAlU35QUh54MHAWI22A3WJkR4d4KhTOrq-2KuCg7Obi3SCoZmVWb28tztUwN6ttc4PJmM370g_YNCiv5Q6F95QgozYAGpu7Kc8ckcsORixNAUpqTCYaZHmST7bxCXDGPaL45H4zHe6IkU-Tf06rY7DoKeMgjGTz3Pb8hrXRXdSCYz9y0MjwGledXqnLiww0Dn_q-qWgOqQs6NeiLG5IqWKJG2e0buav2l_fH-biflRHjpidaTvFnTMUPf9k9-ygWwiWDzM9OD0X-mNdEI6WNe_27O9CtmUTxlBgRJ2tYyhF32a3flQXaA4m34PPXD_HyxFYRQXfqTt_7uaV7NinsnwN8Ll9ccFdXw8BuANu8j24zvBP0zvUyo9d1ywqn0Cw2wt-vPUWF7sZifTLkdr9O7mcAN08ByaIc1MR5ULI-lUsfi6U"
	notification, err := asyncNotifyAPI.NotifyRequest(cipher)
	if err != nil {
		log.Printf("Failed to decrypt notification: %v", err)
	} else {
		fmt.Printf("Notification - Side: %s, Symbol: %s, Amount: %s, Status: %d\n",
			notification.Side, notification.Symbol, notification.Amount.String(), notification.Status)
	}

	withdrawArgs, err = asyncNotifyAPI.VerifyRequest(cipher)
	if err != nil {
		log.Printf("Failed to verify notification: %v", err)
	} else {
		fmt.Printf("Verified Notification - Symbol: %s, Amount: %s, RequestId: %s\n",
			withdrawArgs.Symbol, withdrawArgs.Amount, withdrawArgs.RequestID)
	}

	VerifyResponse, err := asyncNotifyAPI.VerifyResponse(withdrawArgs)
	if err != nil {
		log.Printf("Failed to verify notification: %v", err)
	} else {
		fmt.Printf("Verified Notification - Side: %s, \n",
			VerifyResponse)
	}
	fmt.Println("\n=== All Examples Completed ===")
}
