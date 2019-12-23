package conf

import (
	"webProject/cache"
)

func InitCache(dbName string) {
	cache.NewCacheHandle(dbName)
}
