package models

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	db "my-blog-by-go/database"
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

// GetPostByID 根据ID获取文章
func GetPostByID(id int64) *MPost {
	post := &MPost{}

	err := db.ORM.Where("id = ?", id).First(post).Error
	if err != nil || post.Id <= 0 {
		return nil
	}

	return post
}

// GetPostByIDs 根据IDS获取文章列表
func GetPostByIDs(ids []int64) []*MPost {
	var posts []*MPost
	if len(ids) <= 0 {
		return posts
	}
	err := db.ORM.Where("id in (?)", ids).Find(&posts).Error
	if err != nil {
		log.Println("GetPostByIDs ERROR:", err)
		return nil
	}

	return posts
}

// GetPostByTitle 根据标题获取post
func GetPostByTitle(title string) *MPost {
	post := &MPost{}
	err := db.ORM.Where("title = ?", title).First(post).Error
	if err != nil || post.Id <= 0 {
		return nil
	}

	return post
}

// GetPosts 获取所有的文章
func GetPosts() []*MPost {
	var post []*MPost
	err := db.ORM.Order("like_count desc").Find(&post).Error
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return post
}

// InsertPost 将标题插入到posts表
func InsertPost(title, content string, ch chan int64) {

	id, err := InsertPostCore(title, content)
	defer func() {
		ch <- id
	}()
	if err != nil {
		log.Println("InsertPost ERROR:", err)
		return
	}
	return
}

// InsertPost 将标题插入到posts表
func InsertPostCore(title, content string) (int64, error) {
	newPost := &MPost{
		Title:      title,
		Content:    content,
		LikeCount:  0,
		CreateTime: time.Now().Unix(),
	}

	err := db.ORM.Create(newPost).Error
	if err != nil {
		log.Println("InsertPostCore ERROR:", err)
		return 0, err
	}
	return newPost.Id, nil
}

// RemovePostByID 根据ID删除post
func RemovePostByID(id int64) error {
	err := db.ORM.Model(&MPost{}).Delete("id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

// AddLikeByPostId 给某个post点赞
func AddLikeByPostId(id int64) error {

	err := db.ORM.Model(&MPost{}).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return err
	}

	return nil
}

// InsertPostByTransaction 利用事务写入发文数据
func InsertPostByTransaction(requestInfo InsertPostReq) (int64, error) {
	if requestInfo.Title == "" {
		return 0, errors.New("check title error")
	}
	newPost := &MPost{
		Title:      requestInfo.Title,
		Content:    requestInfo.Content,
		CreateTime: time.Now().Unix(),
		LikeCount:  0,
	}
	db := db.ORM.Begin()

	// 发文数据
	err := db.Model(&MPost{}).Create(newPost).Error
	if err != nil {
		db.Rollback()
		log.Println("InsertPostByTransaction ERROR:", err)
		return 0, err
	}

	log.Println("[INFO] postId: ", newPost.Id)
	if newPost.Id <= 0 {
		return 0, errors.New("check postID error")
	}

	// label数据
	labelIDs := make([]int64, len(requestInfo.Label))
	for i, v := range requestInfo.Label {
		labelQuery := &MLabel{}
		db.Model(&MLabel{}).Where("label = ?", v).First(labelQuery)
		if labelQuery.Id > 0 {
			labelIDs[i] = labelQuery.Id
			continue
		}
		newLabel := &MLabel{Label: v, CreateTime: time.Now().Unix()}
		db.Model(&MLabel{}).Create(newLabel)
		if newLabel.Id > 0 {
			labelIDs[i] = newLabel.Id

		}
	}

	// category数据
	categoryIDs := make([]int64, len(requestInfo.Categories))
	for i, v := range requestInfo.Categories {
		categoryQuery := &MCategory{}
		db.Model(&MCategory{}).Where("category = ?", v).First(categoryQuery)
		if categoryQuery.Id > 0 {
			categoryIDs[i] = categoryQuery.Id
			continue
		}
		newCategories := &MCategory{Category: v, CreateTime: time.Now().Unix()}
		db.Model(&MCategory{}).Create(newCategories)
		if newCategories.Id > 0 {
			categoryIDs[i] = newCategories.Id
		}
	}

	// post和label的关系
	for _, v := range labelIDs {
		if v <= 0 {
			continue
		}
		mPostAndLabel := &MPostAndLabel{
			PostId:     newPost.Id,
			LabelId:    v,
			CreateTime: time.Now().Unix(),
		}
		err := db.Model(&MPostAndLabel{}).Create(mPostAndLabel).Error
		if err != nil {
			db.Rollback()
			log.Println("InsertPostByTransaction ERROR:", err)
			return 0, err
		}
	}

	// post和category的关系
	for _, v := range categoryIDs {
		mPostAndCategory := &MPostAndCategory{
			PostId:     newPost.Id,
			CategoryId: v,
			CreateTime: time.Now().Unix(),
		}
		err := db.Model(&MPostAndCategory{}).Create(mPostAndCategory).Error
		if err != nil {
			db.Rollback()
			log.Println("InsertPostByTransaction ERROR:", err)
			return 0, err
		}
	}

	db.Commit()

	return newPost.Id, nil
}
