package consulops

import (
	"errors"
	"path"
	proconfig "github.com/connext-cs/protocol/config"
	zhcnx "github.com/connext-cs/pub/zhlog"

	"github.com/hashicorp/consul/api"
)

func WriteConsulInfo(key, val string) (e error) {
	traceID := zhcnx.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhcnx.Error(traceID, "写入Consul配置发生错误:", e.(error))
		}
	}()
	if path.IsAbs(key) {
		zhcnx.Error(traceID, "key应为相对路径")
		return errors.New("key应为相对路径")
	}
	value := []byte(val)
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: proconfig.CConsulAddr()})
	zhcnx.Assert(err)
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, nil)
	zhcnx.Assert(err)
	return nil
}
