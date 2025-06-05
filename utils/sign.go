package utils

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"net/url"
	"sort"
)

// 计算请求签名
func Sign(params map[string]interface{}, accessKey string) string {

	keys := lo.Keys(params)
	sort.Strings(keys)

	//拼接签名原始字符串
	rawString := ""
	lo.ForEach(keys, func(x string, index int) {
		value := cast.ToString(params[x])
		value = url.QueryEscape(value)
		rawString += fmt.Sprintf("%s=%s", x, value)

		if index != len(params)-1 {
			//不是最后一个,则拼接
			rawString += "&"
		}
	})
	// 3. 将secretKey拼接到最后
	rawString += accessKey

	fmt.Printf("[rawString]%s\n", rawString)

	//4. 计算md5签名
	signResult := GetMD5([]byte(rawString))
	return signResult
}

// 验证签名
func VerifySign(params map[string]interface{}, signKey string) bool {
	// Check if sign exists in params
	signValue, exists := params["signature"]
	if !exists {
		return false
	}

	// Get the sign value and remove it from params
	key := fmt.Sprintf("%v", signValue)
	delete(params, "signature")

	// Generate current signature
	currentKey := Sign(params, signKey)

	// Compare the signatures
	return key == currentKey
}
