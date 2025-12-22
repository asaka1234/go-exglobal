package go_exglobal

import (
	"fmt"
	"testing"
)

func TestWdBack(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL, PAYIN_BANKLIST_URL, PAYOUT_BANKLIST_URL})

	//发请求
	err := cli.WithdrawCallBack(GenWdBackReqDemo(), Process)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func Process(req ExglobalWithdrawBackReq) error {
	fmt.Println(666666)
	return nil
}

func GenWdBackReqDemo() ExglobalWithdrawBackReq {
	return ExglobalWithdrawBackReq{
		RecordId:        877007760748208128,
		UID:             5588506,
		OrderAmount:     "9385",
		Signature:       "a1b538dd38fe34b3df00763916fd0fcd",
		TradeStatus:     4,
		MerchantOrderNo: "202512090649110461",
	}
}
