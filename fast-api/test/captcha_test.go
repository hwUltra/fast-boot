package test

import (
	"fast-boot/common/captcha"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

var ct = captchaTool.NewCaptchaTool(captchaTool.CaptchaConf{
	Type:      captchaTool.MathType,
	Store:     captchaTool.RedisType,
	RedisConf: redis.RedisConf{Host: "192.168.3.88:16379", Type: "node"},
})

func TestCaptcha(t *testing.T) {

	if id, b64s, answer, err := ct.Make(); err != nil {
		t.Errorf("TestCaptcha: %v", err)
		t.FailNow()
	} else {
		fmt.Printf(" id = %s \n b64s = %s \n answer = %s \n", id, b64s, answer)
	}

}

func TestVerify(t *testing.T) {
	b := ct.VerifyCaptcha("Td39eOouS9FA9s9x0D9B", "q597", true)
	fmt.Println("TestVerifyCommon = ", b)
}
