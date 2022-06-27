package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	"github.com/teleport-network/teleport-analytics-dashboard/model"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open("mysql", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}

	orm.AutoMigrate(&model.WalletDaily{})

}

func GetDB() *gorm.DB {
	return orm
}
