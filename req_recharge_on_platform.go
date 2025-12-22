package go_exglobal

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-exglobal/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// dai
func (cli *Client) RechargeOnPlatform(req ExglobalRechargeOnPlatformReq) (*ExglobalRechargeOnPlatformRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.Params.MerchantId

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey, req.CurrencyCoinName)
	params["signature"] = signStr

	//返回值会放到这里
	var result ExglobalRechargeOnPlatformRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#exglobal#deposit#platform->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
