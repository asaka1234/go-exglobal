package go_exglobal5

import (
	"github.com/asaka1234/go-exglobal5/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params Buy365InitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params Buy365InitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
