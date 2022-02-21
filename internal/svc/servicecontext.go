package svc

import (
	"github.com/acger/pair-svc/database"
	"github.com/acger/pair-svc/internal/config"
	"github.com/acger/user-svc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Cache   *redis.Redis
	UserSvc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      database.NewMysql(&c),
		Cache:   redis.New(c.Cache[0].Host, redis.WithPass(c.Cache[0].Pass)),
		UserSvc: user.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
