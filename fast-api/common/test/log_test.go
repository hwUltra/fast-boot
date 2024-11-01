package test

import (
	"context"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	logx.WithContext(context.Background()).Errorf("【API-ERR】 : %+v ", "xxx")
}

func Test_Time(t *testing.T) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeStr := time.Now().Format(globalkey.DateTimeFormatYMDTime)
	nowTimeStr := time.Now().In(cstSh).Format(globalkey.DateTimeFormatStandardTime)

	todayBeginTimeStr := timeStr + " 00:00:00"
	todayEndTimeStr := timeStr + " 23:59:59"

	formatTime1, _ := time.Parse(globalkey.DateTimeFormatStandardTime, nowTimeStr)
	formatTime2, _ := time.Parse(globalkey.DateTimeFormatStandardTime, todayBeginTimeStr)
	formatTime3, _ := time.Parse(globalkey.DateTimeFormatStandardTime, todayEndTimeStr)

	fmt.Println("now", formatTime1)
	fmt.Println("begin", formatTime2)
	fmt.Println("end", formatTime3)

	td := time.Date(2024, 10, 29, 18, 30, 0, 0, cstSh)
	fmt.Println("指定时间：", td)

	tp, _ := time.ParseDuration("24h")
	m1 := time.Now().Add(tp).Format(globalkey.DateTimeFormatYMDTime) + " 12:00:00"
	fmt.Println("m1：", m1)
}
