package test

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestLog(t *testing.T) {
	logx.WithContext(context.Background()).Errorf("【API-ERR】 : %+v ", "xxx")
}
