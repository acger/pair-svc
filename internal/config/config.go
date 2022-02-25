package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type ElasticsearchConf struct {
	Addresses []string
	Username  string
	Password  string
}

type Config struct {
	zrpc.RpcServerConf
	Datasource    string
	Cache         cache.CacheConf
	UserSvc       zrpc.RpcClientConf
	Elasticsearch ElasticsearchConf
}
