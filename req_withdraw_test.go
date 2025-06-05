package go_exglobal5

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {

	//构造client
	cli := NewClient(nil, ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL})

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
		MerchantOrderNo:  "111",
		CurrencyCoinName: "VND",
		ChannelCode:      "ScanQRCode", ////网银扫码:ScanQRCode, 银行直连:BankDirect
		Amount:           "100015",
		BankCode:         "ACB",
		BankName:         "ACB",
		BankBranchName:   "aa",
		BankUserName:     "cy",
		BankAccount:      "107719719971",
	}
}
