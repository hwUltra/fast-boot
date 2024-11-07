package cachex

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type Store struct {
	cache.Cache
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("fastBootCache")
	cacheErr      = errors.New("not found")
)

func NewStore(conf cache.CacheConf) *Store {
	cc := cache.New(conf, singleFlights, stats, cacheErr)
	return &Store{
		cc,
	}
}

func (s *Store) FormatPrimary(keyPrefix string, primary any) string {
	return fmt.Sprintf("%s%v", keyPrefix, primary)
}
