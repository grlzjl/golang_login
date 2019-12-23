package models

import (
	"log"
	"regexp"

	db "my-blog-by-go/database"
)

type MCategory struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Category   string `gorm:"column:category" json:"category"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

func (p *MCategory) TableName() string {
	return "categories"
}

// GetLabels 获取所有的标签
func GetCategories() []*MCategory {
	var categories []*MCategory

	err := db.ORM.Find(&categories).Error
	if err != nil {
		log.Println("MCategory Find ERROR:", err)
		return nil
	}
	return categories
}

func InsertCategory(categories string, ch chan []int64) {
	categoriesArr := regexp.MustCompile(`\s*,\s*`).Split(categories, -1)
	ids := make([]int64, len(categoriesArr))
	for i, v := range categoriesArr {
		if len(v) <= 0 {
			continue
		}
		newCategoriesQuery := &MCategory{}
		db.ORM.Where("category=?", v).First(newCategoriesQuery)
		if newCategoriesQuery.Id > 0 {
			ids[i] = newCategoriesQuery.Id
			continue
		}
		newCategories := &MCategory{Category: v}
		db.ORM.Create(newCategories)
		if newCategories.Id > 0 {
			ids[i] = newCategories.Id
		}
	}
	ch <- ids
}
