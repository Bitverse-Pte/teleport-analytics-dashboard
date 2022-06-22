package model

import "time"

type TwitterAccount struct {
	Id   int `gorm:"primaryKey"`
	Name string
	// Username is used for login, and is unique
	Username       string    `gorm:"unique"`
	AccountID      string    `gorm:"column:accountId"`
	AccessToken    string    `gorm:"column:accessToken"`
	RefreshToken   string    `gorm:"column:refreshToken"`
	ExpiresAt      time.Time `gorm:"column:expiresAt"`
	FollowersCount int       `gorm:"column:followersCount"`
	TweetCount     int       `gorm:"column:tweetCount"`
}

func (TwitterAccount) TableName() string {
	return "TwitterAccount"
}

type TwitterAccountRealTimeStat struct {
	Id               int    `gorm:"primaryKey"`
	TwitterAccountId string `gorm:"column:twitterAccountId"`
	Date             time.Time
	FollowersCount   int `gorm:"column:followersCount"`
	FollowingCount   int `gorm:"column:followingCount"`
	TweetCount       int `gorm:"column:tweetCount"`
	ListedCount      int `gorm:"column:listedCount"`
}

func (TwitterAccountRealTimeStat) TableName() string {
	return "TwitterAccountRealTimeStat"
}

type TwitterAccountDailyStat struct {
	Id                int    `gorm:"primaryKey"`
	TwitterAccountId  string `gorm:"column:twitterAccountId"`
	Date              time.Time
	FollowersCount    int `gorm:"column:followersCount"`
	NewFollowersCount int `gorm:"column:newFollowersCount"`
	TweetCount        int `gorm:"column:tweetCount"`
	NewTweetCount     int `gorm:"column:newTweetCount"`
}

func (TwitterAccountDailyStat) TableName() string {
	return "TwitterAccountDailyStat"
}
