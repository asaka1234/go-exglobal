package go_exglobal

import (
	"fmt"
	"github.com/asaka1234/go-exglobal/utils"
	"github.com/go-resty/resty/v2"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

type MyRequestBody struct {
	Uid       int64  `json:"uid"`
	Signature string `json:"signature"`
}

func TestGetPaymentChannelList(t *testing.T) {
	client := resty.New()
	//url := "https://api.exlinked.global/coin/pay/proxy/query/paymentChannelList"
	url := "https://api.exlinked.global/coin/pay/proxy/query/paymentBank"

	params := make(map[string]interface{}, 1)
	params["uid"] = MERCHANT_ID
	signStr := utils.Sign(params, ACCESS_SECRET)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&MyRequestBody{
			Uid:       MERCHANT_ID,
			Signature: signStr,
		}).Post(url)

	if err != nil {
		fmt.Printf("request fail: %v\n", err)
		return
	}

	if resp.IsSuccess() {
		fmt.Printf("response success: %s\n", resp.String())
	} else {
		fmt.Printf("response fail: %s\n", resp.String())
	}
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ExglobalInitParams{MERCHANT_ID, ACCESS_SECRET, BACK_SECRET, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() ExglobalDepositReq {
	// VND
	//return ExglobalDepositReq{
	//	MerchantOrderNo:  "323231224", //商户id
	//	CurrencyCoinName: "VND",
	//	//ChannelCode:      "ScanQRCode",
	//	Amount:        100000,
	//	PaymentMethod: 1,
	//}

	// THB
	//return ExglobalDepositReq{
	//	MerchantOrderNo:  "3232312244", //商户id
	//	CurrencyCoinName: "THB",
	//	//ChannelCode:      "ScanQRCode",
	//	Amount:        100000,
	//	PaymentMethod: 1,
	//}

	// IDR
	return ExglobalDepositReq{
		MerchantOrderNo:  "3232312243", //商户id
		CurrencyCoinName: "IDR",
		CustomerName:     "jane",
		CustomerEmail:    "jane.y@logtec.com",
		CustomerPhone:    "34256646",
		Memo:             "test",
		//ChannelCode:      "ScanQRCode",
		Amount:        100000,
		PaymentMethod: 3,
	}
}
