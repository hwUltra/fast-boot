package test

import (
	"fast-boot/app/rpc/model"
	"fmt"
	"github.com/hwUltra/fb-tools/gormV2"
	"testing"
)

func TestPgSql(t *testing.T) {
	gormConn, err := gormV2.GetSqlDriver(
		gormV2.Conf{
			SqlType:       1,
			IsOpenReadDb:  false,
			SlowThreshold: 30,
			Write: gormV2.ConfigParamsDetail{
				Host:               "192.168.3.88",
				DataBase:           "mall-boot",
				Port:               15432,
				Prefix:             "",
				User:               "kyle",
				Pass:               "123123",
				Charset:            "utf8",
				SetMaxIdleConn:     10,
				SetMaxOpenConn:     128,
				SetConnMaxLifetime: 60,
			},
			Read: gormV2.ConfigParamsDetail{
				Host:               "192.168.3.88",
				DataBase:           "mall-boot",
				Port:               15432,
				Prefix:             "",
				User:               "kyle",
				Pass:               "123123",
				Charset:            "utf8",
				SetMaxIdleConn:     10,
				SetMaxOpenConn:     128,
				SetConnMaxLifetime: 60,
			},
		})
	fmt.Println(gormConn, err)
	items := make([]*model.SysMenuModel, 0)
	gormConn.Model(&model.SysMenuModel{}).
		Order("id asc").Find(&items)

	//list := make([]*model.SysMenuModel, 0)
	for _, item := range items {
		fmt.Println(item)
	}

}
