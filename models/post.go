package models

import (
	"awesomeProject/dao/mysql"
	"time"
)

// Post 内存对齐
type Post struct {
	ID          int64     `json:"id,string" gorm:"post_id"`
	AuthorID    int64     `json:"author_id" gorm:"author_id"`
	CommunityID int64     `json:"community_id" gorm:"community_id" binding:"required"`
	Status      int32     `json:"status" gorm:"status"`
	Title       string    `json:"title" gorm:"title" binding:"required"`
	Content     string    `json:"content" gorm:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" gorm:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}

func CreatePost(p *Post) error {
	err := mysql.Db.Create(p).Error
	return err
}

func GETDetailOfPost(id int64) (*Post, error) {
	var postValues Post
	err := mysql.Db.Where("post_id=?", id).First(&postValues).Error
	return &postValues, err

}
