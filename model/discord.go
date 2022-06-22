package model

import (
	"database/sql/driver"
	"time"
)

type DiscordGuild struct {
	Id string `gorm:"primary_key"`
}

func (DiscordGuild) TableName() string {
	return "DiscordGuild"
}

type DiscordGuildDailyStat struct {
	ID                       uint `gorm:"primary_key"`
	Date                     time.Time
	DiscordGuildId           string       `gorm:"column:discordGuildId"`
	DiscordGuild             DiscordGuild `gorm:"foreignKey:DiscordGuildId"`
	StartTotalMemberCount    int          `gorm:"column:startTotalMemberCount"`
	StartOnlineMemberCount   int          `gorm:"column:startOnlineMemberCount"`
	EndTotalMemberCount      int          `gorm:"column:endTotalMemberCount"`
	EndOnlineMemberCount     int          `gorm:"column:endOnlineMemberCount"`
	HighestOnlineMemberCount int          `gorm:"column:highestOnlineMemberCount"`
	LowestOnlineMemberCount  int          `gorm:"column:lowestOnlineMemberCount"`
}

func (DiscordGuildDailyStat) TableName() string {
	return "DiscordGuildDailyStat"
}

type DiscordGuildStat struct {
	ID                uint         `gorm:"primary_key"`
	CreatedAt         time.Time    `gorm:"column:createdAt"`
	DiscordGuildId    string       `gorm:"column:discordGuildId"`
	DiscordGuild      DiscordGuild `gorm:"foreignKey:DiscordGuildId"`
	TotalMemberCount  int          `gorm:"column:totalMemberCount"`
	OnlineMemberCount int          `gorm:"column:onlineMemberCount"`
}

func (DiscordGuildStat) TableName() string {
	return "DiscordGuildStat"
}

type DiscordGuildMember struct {
	Id             string       `gorm:"primary_key"`
	DiscordGuildId string       `gorm:"column:discordGuildId"`
	DiscordGuild   DiscordGuild `gorm:"foreignKey:DiscordGuildId"`
	MessageQty     int          `gorm:"column:messageQty"`
	LastSeen       time.Time    `gorm:"column:lastSeen"`
}

func (DiscordGuildMember) TableName() string {
	return "DiscordGuildMember"
}

type DiscordChannel struct {
	Id             string `gorm:"primary_key"`
	Name           string
	Type           DiscordChannelType
	DiscordGuildId string       `gorm:"column:discordGuildId"`
	DiscordGuild   DiscordGuild `gorm:"foreignKey:DiscordGuildId"`
	CreatedAt      time.Time    `gorm:"column:createdAt"`
}

type DiscordChannelType string

const (
	DiscordChannelTypeGuildText          = "GUILD_TEXT"
	DiscordChannelTypeDM                 = "DM"
	DiscordChannelTypeGuildVoice         = "GUILD_VOICE"
	DiscordChannelTypeGroupDM            = "GROUP_DM"
	DiscordChannelTypeGuildCategory      = "GUILD_CATEGORY"
	DiscordChannelTypeNews               = "GUILD_NEWS"
	DiscordChannelTypeStore              = "GUILD_STORE"
	DiscordChannelTypeUnUsed1            = "UNUSED1"
	DiscordChannelTypeUnUsed2            = "UNUSED2"
	DiscordChannelTypeUnUsed3            = "UNUSED3"
	DiscordChannelTypeGuildNewsThread    = "GUILD_NEWS_THREAD"
	DiscordChannelTypeGuildPublicThread  = "GUILD_PUBLIC_THREAD"
	DiscordChannelTypeGuildPrivateThread = "GUILD_PRIVATE_THREAD"
	DiscordChannelTypeStageVoice         = "GUILD_STAGE_VOICE"
	DiscordChannelTypeGuildDirectory     = "GUILD_DIRECTORY"
)

func (t *DiscordChannelType) Scan(value interface{}) error {
	*t = DiscordChannelType(value.([]byte))
	return nil
}

func (t DiscordChannelType) Value() (driver.Value, error) {
	return string(t), nil
}

type DiscordGuildChannelStat struct {
	ID                uint           `gorm:"primary_key"`
	CreatedAt         time.Time      `gorm:"column:createdAt"`
	DiscordChannelId  string         `gorm:"column:discordChannelId"`
	DiscordChannel    DiscordChannel `gorm:"foreignKey:DiscordChannelId"`
	TotalMemberCount  int            `gorm:"column:totalMemberCount"`
	OnlineMemberCount int            `gorm:"column:onlineMemberCount"`
}

func (DiscordGuildChannelStat) TableName() string {
	return "DiscordGuildChannelStat"
}

type DiscordGuildChannelDailyStat struct {
	ID                       uint `gorm:"primary_key"`
	Date                     time.Time
	DiscordChannelId         string         `gorm:"column:discordChannelId"`
	DiscordChannel           DiscordChannel `gorm:"foreignKey:DiscordChannelId"`
	StartTotalMemberCount    int            `gorm:"column:startTotalMemberCount"`
	StartOnlineMemberCount   int            `gorm:"column:startOnlineMemberCount"`
	EndTotalMemberCount      int            `gorm:"column:endTotalMemberCount"`
	EndOnlineMemberCount     int            `gorm:"column:endOnlineMemberCount"`
	HighestOnlineMemberCount int            `gorm:"column:highestOnlineMemberCount"`
	LowestOnlineMemberCount  int            `gorm:"column:lowestOnlineMemberCount"`
}

func (DiscordGuildChannelDailyStat) TableName() string {
	return "DiscordGuildChannelDailyStat"
}
