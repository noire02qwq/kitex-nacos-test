package common

import (
	"fmt"

	"github.com/kitex-contrib/registry-nacos/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	NacosAddr = "127.0.0.1"
	NacosPort = 8848
)

// NewNacosRegistry creates a nacos registry
func NewNacosRegistry() (*nacos.Registry, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosAddr, uint64(NacosPort)),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("new nacos client failed: %v", err)
	}

	r := nacos.NewNacosRegistry(cli)
	return r, nil
}

// NewNacosResolver creates a nacos resolver
func NewNacosResolver() (*nacos.Resolver, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosAddr, uint64(NacosPort)),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("new nacos client failed: %v", err)
	}

	r := nacos.NewNacosResolver(cli)
	return r, nil
}
