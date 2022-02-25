package database

import (
	"encoding/json"
	"github.com/acger/pair-svc/pair"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

func GetCacheElement(uid uint64, client *redis.Redis) *pair.EleSaveReq {

	cacheKey := "element"
	uidStr := strconv.FormatInt(int64(uid), 10)
	eleCache, _ := client.Hget(cacheKey, uidStr)

	element := pair.EleSaveReq{}
	json.Unmarshal([]byte(eleCache), &element)

	return &element
}
