package model

import "time"

type Tweet struct {
	Id                int       `gorm:"primaryKey"`
	TweetId           string    `gorm:"unique;column:tweetId"`
	CreatedAt         time.Time `gorm:"column:createdAt"`
	Text              string
	Impressions       int
	Retweets          int
	QuoteTweets       int `gorm:"column:quoteTweets"`
	Likes             int
	Replies           int
	UrlLinkClicks     int `gorm:"column:urlLinkClicks"`
	UserProfileClicks int `gorm:"column:userProfileClicks"`
	MediaViews        int `gorm:"column:mediaViews"`
}

func (Tweet) TableName() string {
	return "Tweet"
}
