package go_exglobal

import (
	"fmt"
	"testing"
)

func TestClient_RechargeOnPlatform(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.RechargeOnPlatform(GenRechargeOnPlateformRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenRechargeOnPlateformRequestDemo() ExglobalRechargeOnPlatformReq {
	return ExglobalRechargeOnPlatformReq{
		MerchantOrderNo:  "20251222039384595",
		CurrencyCoinName: "JPY",
		Amount:           100000,
	}

	// return ExglobalRechargeOnPlatformReq{
	// 	MerchantOrderNo:  "20251222039384591", //商户id
	// 	CurrencyCoinName: "INR",
	// 	Amount:           100000,
	// }
}
