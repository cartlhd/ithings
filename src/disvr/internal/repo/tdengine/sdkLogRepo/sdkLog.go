package sdkLogRepo

import (
	"github.com/i-Things/things/shared/clients"
	"github.com/i-Things/things/shared/store"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

type SDKLogRepo struct {
	t *clients.Td
	store.SDKLogStore
}

func NewSDKLogRepo(dataSource string) *SDKLogRepo {
	td, err := clients.NewTDengine(dataSource)
	if err != nil {
		logx.Error("NewTDengine err", err)
		os.Exit(-1)
	}
	return &SDKLogRepo{t: td}
}
