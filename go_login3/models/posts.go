package models

import (
	"fmt"
	"go_learn/database"
)

type MPost struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	LikeCount  int32  `gorm:"column:like_count" json:"like_count"`
}

func (p *MPost) TableName() string {
	return "posts"
}

func GetPosts() []*MPost {
	var results []*MPost

	err := database.DBInstance.Order("like_count desc").Find(&results).Error
	if err != nil {
		fmt.Println("find posts error", err)
		return nil
	}
	return results
}
