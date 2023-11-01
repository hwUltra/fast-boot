package middleware

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type CheckUrlMiddleware struct {
	Redis *redis.Redis
}

func NewCheckUrlMiddleware(Redis *redis.Redis) *CheckUrlMiddleware {
	return &CheckUrlMiddleware{Redis: Redis}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("CheckUrlMiddleware --- begin")

		next(w, r)
	}
}
