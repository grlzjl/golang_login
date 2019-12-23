package models

import (
	"log"
	"regexp"

	db "my-blog-by-go/database"
)

type MLabel struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Label      string `gorm:"column:label" json:"label"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

func (p *MLabel) TableName() string {
	return "labels"
}

// GetLabels 获取所有的标签
func GetLabels() []*MLabel {
	var labels []*MLabel
	err := db.ORM.Find(&labels).Error
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return labels
}

//插入标签(以","分隔). 将新插入的标签标识传入 ch channel 中.
func InsertLabel(labels string, ch chan []int64) {
	labelsArr := regexp.MustCompile(`\s*,\s*`).Split(labels, -1)
	ids := make([]int64, len(labelsArr))
	for i, v := range labelsArr {
		if len(v) <= 0 {
			continue
		}
		labelQuery := &MLabel{}
		db.ORM.Where("label=?", v).First(labelQuery)
		if labelQuery.Id > 0 {
			ids[i] = labelQuery.Id
			continue
		}
		newLabel := &MLabel{Label: v}
		db.ORM.Create(newLabel)
		if newLabel.Id > 0 {
			ids[i] = newLabel.Id
		}
	}
	ch <- ids
}
