package go_exglobal

import (
	"github.com/asaka1234/go-exglobal/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params ExglobalInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params ExglobalInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
