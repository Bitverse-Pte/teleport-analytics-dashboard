package model

import (
	"math/big"
	"time"
)

type TelegramGroup struct {
	ID     uint  `gorm:"primary_key"`
	ChatId int64 `gorm:"type:decimal(65,0);column:chatId"`
}

func (TelegramGroup) TableName() string {
	return "TelegramGroup"
}

type TelegramGroupStats struct {
	ID                uint          `gorm:"primary_key"`
	ChatGroupID       int64         `gorm:"column:groupId"`
	ChatGroup         TelegramGroup `gorm:"foreignKey:ChatGroupID"`
	Date              time.Time
	NewMemberCount    int `gorm:"column:newMemberCount"`
	MessageCount      int `gorm:"column:messageCount"`
	TotalMemberCount  int `gorm:"column:totalMemberCount"`
	ActiveMemberCount int `gorm:"column:activeMemberCount"`
}

func (TelegramGroupStats) TableName() string {
	return "TelegramGroupStats"
}

type TelegramGroupDailyStat struct {
	ID                uint          `gorm:"primary_key"`
	GroupID           big.Int       `gorm:"column:groupId"`
	Group             TelegramGroup `gorm:"foreignKey:GroupID"`
	Date              time.Time
	NewMemberCount    int `gorm:"column:newMemberCount"`
	MessageCount      int `gorm:"column:messageCount"`
	TotalMemberCount  int `gorm:"column:totalMemberCount"`
	ActiveMemberCount int `gorm:"column:activeMemberCount"`
}

func (TelegramGroupDailyStat) TableName() string {
	return "TelegramGroupDailyStat"
}

type TelegramChatMember struct {
	ID           uint          `gorm:"primary_key"`
	UserId       big.Int       `gorm:"column:userId"`
	GroupID      big.Int       `gorm:"column:groupId"`
	Group        TelegramGroup `gorm:"foreignKey:GroupID"`
	JoinAt       time.Time     `gorm:"column:joinAt"`
	MessageCount int           `gorm:"column:messageCount"`
	LastSeen     time.Time     `gorm:"column:lastSeen"`
	ActiveDays   int           `gorm:"column:activeDays"`
}

func (TelegramChatMember) TableName() string {
	return "TelegramChatMember"
}
