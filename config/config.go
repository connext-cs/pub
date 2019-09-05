package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"

	"github.com/connext-cs/pub/log"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source/consul"
)

const (
	ConsulKey       = "micro/config/cache"
	ConsulAddrLocal = "http://127.0.0.1:8500"
)

var conf = make(map[string]interface{})

func initConfig(consulAddr string) error {
	configServer := config.NewConfig()
	if err := configServer.Load(consul.NewSource(
		consul.WithAddress(consulAddr),
		consul.WithPrefix("micro/config/"),
		consul.StripPrefix(true),
	)); err != nil {
		log.Error("err:", err.Error())
		return err
	}
	var defaultConfig = api.DefaultConfig()
	defaultConfig.Address = consulAddr
	consulClient, err := api.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}
	value := configServer.Get("cache")
	if err := value.Scan(&conf); err != nil {
		log.Error("1111cache %+v", err)
		return err
	}
	watcher, err := configServer.Watch("cache")
	if err != nil {
		return errors.New(fmt.Sprintf("unable to create initconfig: %s", err.Error()))
	}
	go func() {
		for {
			_, err := watcher.Next()
			if err != nil {
				fmt.Println("watcher next() error", err)
				time.Sleep(time.Second)
				continue
			}
			if kvPair, _, err := consulClient.KV().Get(ConsulKey, nil); err != nil {
				log.Error("%+v", err)
			} else {
				if err := json.Unmarshal(kvPair.Value, &conf); err != nil {
					log.Error("%+v", err)
				}
			}
		}
	}()
	return nil
}

func cStr(key string, defaultvalue string) (value string) {
	if key == "consul_addr" {
		return ConsulAddress()
	}
	if len(conf) == 0 {
		consulAddr := ConsulAddress()
		if err := initConfig(consulAddr); err != nil {
			fmt.Println("InitConfig err:" + err.Error() + " use default, key:" + key + " defaultvalue:" + defaultvalue)
			return defaultvalue
		}
	}
	for k, v := range conf {
		if k == key {
			switch v.(type) {
			case string:
				return v.(string)
			default:
				return defaultvalue
			}
		}
	}
	return defaultvalue
}

func cInt(key string, defaultvalue int) (value int) {
	if len(conf) == 0 {
		consulAddr := ConsulAddress()
		if err := initConfig(consulAddr); err != nil {
			fmt.Println("InitConfig err:" + err.Error() + " use default, key:" + key + " defaultvalue:" + strconv.Itoa(defaultvalue))
			return defaultvalue
		}
	}

	for k, v := range conf {
		if k == key {
			switch v.(type) {
			case string:
				value, _ = strconv.Atoi(v.(string))
				return
			default:
				return defaultvalue
			}
		}
	}
	return defaultvalue
}

func ConsulAddress() string {
	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = ConsulAddrLocal
	}

	return consulAddr
}

func URLLive(url string) bool {
	if !strings.Contains(strings.ToLower(url), "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func ConsulAlive() bool {
	return URLLive(ConsulAddress())
}
