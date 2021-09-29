package svc

import (
	"github.com/acger/pair-svc/database"
	"github.com/acger/pair-svc/internal/config"
	"github.com/acger/user-svc/userclient"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Cache   *redis.Redis
	UserSvc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      database.NewMysql(&c),
		Cache:   redis.NewRedis(c.Cache[0].Host, c.Cache[0].Type, c.Cache[0].Pass),
		UserSvc: userclient.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
