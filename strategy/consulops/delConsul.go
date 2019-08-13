package consulops

import (
	proconfig "github.com/connext-cs/protocol/config"
	zhcnx "github.com/connext-cs/pub/zhlog"

	"github.com/hashicorp/consul/api"
)

func DelConsulInfo(key string) (e error) {
	traceID := zhcnx.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhcnx.Error(traceID, "删除Consul配置发生错误:", e.(error))
			return
		}
	}()
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: proconfig.CConsulAddr()})
	zhcnx.Assert(err)
	_, err = consulClient.KV().Delete(key, nil)
	zhcnx.Assert(err)
	return nil
}
