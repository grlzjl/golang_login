package models

import (
	"time"

	db "my-blog-by-go/database"
)

type MPostAndCategory struct {
	Id         int64 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PostId     int64 `gorm:"column:post_id" json:"post_id"`
	CategoryId int64 `gorm:"column:category_id" json:"category_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

func (p *MPostAndCategory) TableName() string {
	return "post_category"
}

// InsertPostAndCategory 将标题插入到post_category表
func InsertPostAndCategory(PostId, CategoryId int64) (int64, error) {
	mPostAndCategory := &MPostAndCategory{
		PostId:     PostId,
		CategoryId: CategoryId,
		CreateTime: time.Now().Unix(),
	}
	err := db.ORM.Model(&MPostAndCategory{}).Create(mPostAndCategory).Error
	return mPostAndCategory.Id, err
}

// GetPostsByCategoryId 根据category_id查询post_id
func GetPostsByCategoryId(categoryId int64) ([]int64, error) {
	var result []int64
	var categoryAndPosts []*MPostAndCategory
	err := db.ORM.Model(&MPostAndCategory{}).Where("category_id = ?", categoryId).
		Find(&categoryAndPosts).Error
	if err != nil {
		return result, err
	}
	for idx := range categoryAndPosts {
		result = append(result, categoryAndPosts[idx].PostId)
	}

	return result, nil
}

// RemovePCByPostID 根据post_id删除对应的记录
func RemovePCByPostID(postID int64) error {
	return db.ORM.Model(&MPostAndCategory{}).Delete("post_id = ?", postID).Error
}
