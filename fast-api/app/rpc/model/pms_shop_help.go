package model

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"strings"
)

type PmsShopModelHelp struct {
	PmsShop PmsShopModel
	Db      *gorm.DB
	Rsc     *redis.Redis
}

var (
	cacheBootPmsShopIdPrefix = "cache:boot:pms_shop:id:"
)

func CreatePmsShopModelHelp(db *gorm.DB, rsd *redis.Redis) *PmsShopModelHelp {
	return &PmsShopModelHelp{
		PmsShop: PmsShopModel{},
		Db:      db,
		Rsc:     rsd,
	}
}

func (m *PmsShopModelHelp) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBootPmsShopIdPrefix, primary)
}

func (m *PmsShopModelHelp) Insert(data *PmsShopModel) (int64, error) {

	if err := m.Db.Create(data).Error; err != nil {
		return 0, err
	}
	bootPmsShopIdKey := m.formatPrimary(data.Id)
	bytes, _ := json.Marshal(data)
	err := m.Rsc.Set(bootPmsShopIdKey, string(bytes))
	if err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *PmsShopModelHelp) Delete(ids string) {
	idsArr := strings.Split(ids, ",")
	m.Db.Delete(&m.PmsShop, idsArr)
	for _, id := range idsArr {
		bootPmsShopIdKey := m.formatPrimary(id)
		_, _ = m.Rsc.Del(bootPmsShopIdKey)

	}
}

func (m *PmsShopModelHelp) Update(data *PmsShopModel) error {
	if err := m.Db.Save(data).Error; err != nil {
		return err
	}
	bootPmsShopIdKey := m.formatPrimary(data.Id)
	bytes, _ := json.Marshal(data)
	err := m.Rsc.Set(bootPmsShopIdKey, string(bytes))
	if err != nil {
		return err
	}
	return nil
}

func (m *PmsShopModelHelp) Find(id int64) (*PmsShopModel, error) {
	bootPmsShopIdKey := m.formatPrimary(id)
	item := PmsShopModel{}
	if exists, err := m.Rsc.Exists(bootPmsShopIdKey); err == nil {
		if exists {
			res, err := m.Rsc.Get(bootPmsShopIdKey)
			if err == nil {
				if err = json.Unmarshal([]byte(res), &item); err == nil {
					return &item, nil
				}
			}
		}
	}
	if err := m.Db.First(&item, id).Error; err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(item)
	_ = m.Rsc.Set(bootPmsShopIdKey, string(bytes))
	return &item, nil

}
