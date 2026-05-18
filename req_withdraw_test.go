package go_exglobal

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

// [rawString]amount=8365&bankAccount=6767484&bankBranchName=鹿兒島銀行&bankName=鹿兒島銀行&bankUserName=jane jane&currencyCoinName=JPY&customerPhone=43623326&merchantOrderNo=202603161237070802&paymentType=VirtualAccount&uid=5588432&key=AUWAwePaoiOXetbpv2q7XJsevtCnUE9V

func GenWithdrawRequestDemo() ExglobalWithdrawReq {

	// KRW 韩国
	// 	{
	//    "amount": "74210",
	//    "bankAccount": "325236",
	//    "bankBranchName": "BUSAN",
	//    "bankUserName": "張三",
	//    "currencyCoinName": "KRW",
	//    "customerEmail": "jane.y1@yopmail.com",
	//    "customerPhone": "1125300231651",
	//    "memo": "prod",
	//    "merchantOrderNo": "202605151113370550",
	//    "paymentType": "VirtualAccount",
	//    "signature": "c1fe7b4722100528c53e25429d669732",
	//    "uid": 5588789
	// }
	return ExglobalWithdrawReq{
		MerchantOrderNo:  "202605151113370551",
		CurrencyCoinName: "KRW",
		Amount:           "74210",
		BankCode:         "BUSAN",
		BankName:         "BNK부산은행(BNKBusanBank)",
		BankAccount:      "325236",
		BankBranchName:   "BUSAN",
		BankUserName:     "張三",
		CustomerEmail:    "jane.y1@yopmail.com",
		CustomerPhone:    "1125300231651",
		Memo:             "prod",
		PaymentType:      "BankTransfer",
	}

	// BRL
	// return ExglobalWithdrawReq{
	// 	MerchantOrderNo:  "202603161237070809",
	// 	CurrencyCoinName: "BRL",
	// 	Amount:           "8000",
	// 	BankCode:         "CPF",
	// 	BankName:         "CPF",
	// 	BankBranchName:   "KARACHI BRANCH",
	// 	BankUserName:     "jane jane",
	// 	BankAccount:      "6767484",
	// 	// CustomerEmail:    "afsf@gmail.com",
	// 	// CustomerPhone: "43623326",
	// 	Memo: "prod",
	// }

	// PKR
	// return ExglobalWithdrawReq{
	// 	MerchantOrderNo:  "202603161237070808",
	// 	CurrencyCoinName: "PKR",
	// 	Amount:           "8000",
	// 	BankCode:         "HBL",
	// 	BankName:         "HBL",
	// 	BankBranchName:   "KARACHI BRANCH",
	// 	BankUserName:     "jane jane",
	// 	BankAccount:      "6767484",
	// 	CustomerEmail:    "afsf@gmail.com",
	// 	CustomerPhone:    "43623326",
	// 	Memo:             "prod",
	// }

	// JPY
	// 	{
	//    "amount": "8365",
	//    "bankAccount": "6767484",
	//    "bankBranchName": "鹿兒島銀行",
	//    "bankCode": "鹿兒島銀行",
	//    "bankName": "鹿兒島銀行",
	//    "bankUserName": "jane jane",
	//    "currencyCoinName": "JPY",
	//    "memo": "prod",
	//    "merchantOrderNo": "202603161237070801",
	//    "paymentType": "VirtualAccount",
	//    "signature": "c612b81619b813120d120a7291e76981",
	//    "uid": 5588789
	// }

	return ExglobalWithdrawReq{
		MerchantOrderNo:  "202603161237070802",
		CurrencyCoinName: "JPY",
		Amount:           "8000",
		BankCode:         "鹿兒島銀行",
		BankName:         "鹿兒島銀行",
		BankBranchName:   "鹿兒島銀行",
		BankUserName:     "jane jane",
		BankAccount:      "6767484",
		CustomerEmail:    "afsf@gmail.com",
		CustomerPhone:    "43623326",
		Memo:             "prod",
	}

	// return ExglobalWithdrawReq{
	// 	MerchantOrderNo:  "2025468273680231",
	// 	CurrencyCoinName: "INR",
	// 	Amount:           "2005",
	// 	BankName:         "ACB",
	// 	BankBranchName:   "AIRP0000001",
	// 	BankUserName:     "jane",
	// 	BankAccount:      "107719719971",
	// 	CustomerPhone:    "43623326",
	// 	Memo:             "test",
	// }

	// php
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "202512110928168559",
	//	CurrencyCoinName: "PHP",
	//	Amount:           "200",
	//	BankUserName:     "jane",
	//	BankAccount:      "107719719971",
	//	BankCode:         "",
	//	BankName:         "",
	//	BankBranchName:   "",
	//	Memo:             "test",
	//}

	// vnd
	//amount=100000000&bankAccount=107719719971&bankBranchName=aa&bankName=ACB&bankUserName=cy&currencyCoinName=VND&merchantOrderNo=111&paymentType=BankDirect&uid=5588506&key=M2gjZTVutMlhtHYygMnKFQHh58JNW0zm
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "20251211092816855",
	//	CurrencyCoinName: "VND", //"VND",
	//	Amount:           "100000000",
	//	//BankName:         "ACB - NGAN HANG TMCP A CHAU (ACB)",
	//	//BankBranchName:   "ACB",
	//	BankUserName: "jane",
	//	BankAccount:  "107719719971",
	//	Memo:         "test",
	//}

	//{
	//	"amount": "131935",
	//	"bankAccount": "845574",
	//	"bankUserName": "feng",
	//	"currencyCoinName": "VND",
	//	"memo": "prod",
	//	"merchantOrderNo": "202512220318060792",
	//	"paymentType": "BankDirect",
	//	"signature": "926253db582f3fb6c7329857cb1a5a91",
	//	"uid": 5588506
	//}

	// INR
	// amount=1&bankAccount=107719719971&bankBranchName=aa&bankName=ACB&bankUserName=cy&currencyCoinName=INR&customerPhone=43623326&merchantOrderNo=111&paymentType=BankPayout&uid=5588506&key=M2gjZTVutMlhtHYygMnKFQHh58JNW0zm
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "20254682736802",
	//	CurrencyCoinName: "INR",
	//	Amount:           "100",
	//	BankName:         "ACB",
	//	BankBranchName:   "aa",
	//	BankUserName:     "jane",
	//	BankAccount:      "107719719971",
	//	CustomerPhone:    "43623326",
	//	Memo:             "test",
	//}

	// IDR  ==> amount low credit
	// amount=100000&bankAccount=107719719971&bankBranchName=Bank Maybank&bankCode=016&bankName=Bank Maybank&bankUserName=jane&currencyCoinName=IDR&customerEmail=jane.y@logtec.com&customerPhone=43623326&merchantOrderNo=111&paymentType=BankDirect&uid=5588506&key=M2gjZTVutMlhtHYygMnKFQHh58JNW0zm
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "202512015270938",
	//	CurrencyCoinName: "IDR",
	//	Amount:           "100000",
	//	BankCode:         "016",
	//	BankName:         "Bank Maybank",
	//	BankBranchName:   "Bank Maybank",
	//	BankUserName:     "jane",
	//	BankAccount:      "107719719971",
	//	CustomerPhone:    "43623326",
	//	CustomerEmail:    "jane.y@logtec.com",
	//	Memo:             "test",
	//}

	// THB  ==> amount low credit
	// amount=100000&bankAccount=107719719971&bankBranchName=aa&bankCode=BAY&bankName=ธนาคารกรุงศรีอยุธยา(Bank of Ayudhaya Public Company Limited)&bankUserName=jane&currencyCoinName=THB&merchantOrderNo=1742913308094&paymentType=BankDirect&uid=5588506&key=M2gjZTVutMlhtHYygMnKFQHh58JNW0zm
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "1742913308094",
	//	CurrencyCoinName: "THB", //"VND",
	//	Amount:           "100000",
	//	BankName:         "ธนาคารกรุงศรีอยุธยา(Bank of Ayudhaya Public Company Limited)",
	//	BankCode:         "BAY",
	//	BankBranchName:   "aa",
	//	BankUserName:     "jane",
	//	BankAccount:      "107719719971",
	//	Memo:             "test",
	//}

}
