package cachex

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
)

type Store struct {
	Conf  cache.CacheConf
	Cache cache.Cache
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("fastBootCache")
	cacheErr      = errors.New("not found")
)

func NewStore(conf cache.CacheConf) *Store {
	cc := cache.New(conf, singleFlights, stats, cacheErr)
	return &Store{
		conf,
		cc,
	}
}

func (s *Store) FormatPrimary(keyPrefix string, primary any) string {
	return fmt.Sprintf("%s%v", keyPrefix, primary)
}

func (s *Store) ClearRedisPrefix(prefix string) {
	redisClient := redis.MustNewRedis(s.Conf[0].RedisConf)
	if list, _, err := redisClient.Scan(0, prefix, 0); err == nil {
		for _, item := range list {
			_, _ = redisClient.Del(item)
		}
	}
}
