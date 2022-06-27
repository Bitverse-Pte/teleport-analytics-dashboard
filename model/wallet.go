package model

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
	"time"
)

type WalletDaily struct {
	gorm.Model
	Date             time.Time
	Type             WalletType
	DownloadCount    int
	NewDownloadCount int
}

type WalletType string

const (
	WalletTypeExtension = "Extension"
)

func (ct *WalletType) Scan(value interface{}) error {
	*ct = WalletType(value.([]byte))
	return nil
}

func (ct WalletType) Value() (driver.Value, error) {
	return string(ct), nil
}
