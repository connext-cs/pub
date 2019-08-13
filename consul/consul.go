package consul

import (
	"errors"
	"github.com/connext-cs/pub/logs"

	"github.com/hashicorp/consul/api"
)

/*
  写入数据入consul
consulAddr consul的链接地址 127.0.0.1:8500
key        consul存储的key值
value      consul存储的value值
write      consul写入条件, 默认为nil
*/
func WriteKV(consulAddr string, key string, value []byte, write *api.WriteOptions) error {
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	if err != nil {
		logs.Error("api.NewClient error:", err.Error())
		return err
	}
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, write)
	if err != nil {
		logs.Error("consulClient.KV().Put error:", err.Error())
		return err
	}
	return nil
}

/*
  读取数据出consul
consulAddr consul的链接地址 127.0.0.1:8500
key        consul存储的key值
value      consul存储的value值
query      查询条件 默认 nil
*/
func ReadKV(consulAddr string, key string, query *api.QueryOptions) ([]byte, error) {
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	if err != nil {
		return nil, err
	}
	if kvPaira, _, err := consulClient.KV().Get(key, query); err == nil {
		if kvPaira != nil {
			return kvPaira.Value, nil
		}
	} else {
		err := errors.New("kvPaira == nil")
		return nil, err
	}

	return nil, nil
}
