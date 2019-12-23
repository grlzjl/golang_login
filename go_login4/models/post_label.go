package models

import (
	"time"

	db "my-blog-by-go/database"
)

type MPostAndLabel struct {
	Id         int64 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PostId     int64 `gorm:"column:post_id" json:"post_id"`
	LabelId    int64 `gorm:"column:label_id" json:"label_id"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

func (p *MPostAndLabel) TableName() string {
	return "post_label"
}

// InsertPostAndLabel 将标题插入到post_label表
func InsertPostAndLabel(PostId, LabelId int64) (int64, error) {
	mPostAndLabel := &MPostAndLabel{
		PostId:     PostId,
		LabelId:    LabelId,
		CreateTime: time.Now().Unix(),
	}

	err := db.ORM.Model(&MPostAndLabel{}).Create(mPostAndLabel).Error
	return mPostAndLabel.Id, err
}

// GetPostsByLabelID 根据label查询post_id
func GetPostsByLabelID(labelId int64) ([]int64, error) {
	var result []int64
	var labelAndPosts []*MPostAndLabel

	err := db.ORM.Model(&MPostAndLabel{}).Where("label_id = ?", labelId).
		Find(&labelAndPosts).Error
	if err != nil {
		return result, err
	}

	for idx := range labelAndPosts {
		result = append(result, labelAndPosts[idx].PostId)
	}

	return result, nil
}

// RemoveByPostID 根据postid删除对应的记录
func RemovePLByPostID(postID int64) error {
	return db.ORM.Model(&MPostAndLabel{}).Delete("post_id = ?", postID).Error
}
