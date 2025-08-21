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
	// vnd
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "111",
	//	CurrencyCoinName: "VND", //"VND",
	//	//ChannelCode:      "BankDirect", ////网银扫码:ScanQRCode, 银行直连:BankDirect
	//	Amount:         "100000000",
	//	BankCode:       "ACB",
	//	BankName:       "ACB",
	//	BankBranchName: "aa",
	//	BankUserName:   "cy",
	//	BankAccount:    "107719719971",
	//}

	// THB
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "1742913308094",
	//	CurrencyCoinName: "THB", //"VND",
	//	//ChannelCode:      "BankDirect", ////网银扫码:ScanQRCode, 银行直连:BankDirect
	//	Amount:         "100000",
	//	BankName:       "ธนาคารกรุงศรีอยุธยา(Bank of Ayudhaya Public Company Limited)",
	//	BankCode:       "BAY",
	//	BankBranchName: "aa",
	//	BankUserName:   "jane",
	//	BankAccount:    "107719719971",
	//	Memo:           "test",
	//}

	// IDR
	return ExglobalWithdrawReq{
		MerchantOrderNo:  "111",
		CurrencyCoinName: "IDR",
		Amount:           "100000",
		BankCode:         "016",
		BankName:         "Bank Maybank",
		BankBranchName:   "Bank Maybank",
		BankUserName:     "jane",
		BankAccount:      "107719719971",
		CustomerPhone:    "43623326",
		CustomerEmail:    "jane.y@logtec.com",
		Memo:             "test",
	}

	// INR
	//return ExglobalWithdrawReq{
	//	MerchantOrderNo:  "111",
	//	CurrencyCoinName: "INR", //"VND",
	//	//ChannelCode:      "BankDirect", ////网银扫码:ScanQRCode, 银行直连:BankDirect
	//	Amount:         "1",
	//	BankCode:       "ACB",
	//	BankName:       "ACB",
	//	BankBranchName: "aa",
	//	BankUserName:   "cy",
	//	BankAccount:    "107719719971",
	//	CustomerPhone:  "43623326",
	//	Memo:           "test",
	//}
}
