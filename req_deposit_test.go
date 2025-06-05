package go_exglobal5

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() ExglobalDepositReq {
	return ExglobalDepositReq{
		MerchantOrderNo:  "323224", //商户id
		CurrencyCoinName: "VND",
		ChannelCode:      "18.29.120.32",
		Amount:           100,
		PaymentMethod:    3, //商户订单号
	}
}
