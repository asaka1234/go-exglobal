package go_exglobal5

import (
	"crypto/tls"
	"github.com/asaka1234/go-exglobal5/utils"
)

func (cli *Client) GetOrderList() (*ExglobalOrderListRsp, error) {

	rawURL := cli.Params.OrderListUrl

	params := map[string]interface{}{
		"sys_no": cli.Params.MerchantId,
	}

	//签名
	signStr := utils.SignDeposit(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result ExglobalOrderListRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
