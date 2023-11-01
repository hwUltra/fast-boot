package test

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

type UserInfo struct {
	Uid      int64  `json:"uid"`
	UserName string `json:"username"`
}

func TestRedisBuilder(t *testing.T) {

	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"})

	//userInfo := map[string]any{
	//	"name": "kyle",
	//	"age":  10,
	//}
	userInfo := UserInfo{1, "kyle"}
	bytes, _ := json.Marshal(userInfo)
	err = redisClient.Setex("user_1", string(bytes), 1200)
	if err != nil {
		return
	}

	exists, err := redisClient.Exists("user_1")
	if err != nil {
		fmt.Println("exists err:", err)
	}
	if exists {
		infos, err := redisClient.Get("user_1")
		if err == nil {
			//var result map[string]any
			//if err = json.Unmarshal([]byte(infos), &result); err != nil {
			//	fmt.Println("exists err:", err)
			//}
			//fmt.Println("result:", result, result["age"], result["name"])
			var result UserInfo
			if err = json.Unmarshal([]byte(infos), &result); err != nil {
				fmt.Println("exists err:", err)
			}
			fmt.Println("result:", result, result.Uid, result.UserName)

		}

	}

}
