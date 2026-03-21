package client

import (
	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
)

func New() *ad_open_sdk_go.Client {
	cfg := config.NewConfiguration()
	return ad_open_sdk_go.Init(cfg)
}
