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

func GenWithdrawRequestDemo() ExglobalWithdrawReq {
	return ExglobalWithdrawReq{
		MerchantOrderNo:  "2025468273680231",
		CurrencyCoinName: "INR",
		Amount:           "2005",
		BankName:         "ACB",
		BankBranchName:   "AIRP0000001",
		BankUserName:     "jane",
		BankAccount:      "107719719971",
		CustomerPhone:    "43623326",
		Memo:             "test",
	}

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
