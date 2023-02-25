package postMysql

import (
	"awesomeProject/models"
)

func CreatPost(p *models.Post) error {
	err := models.CreatePost(p)
	return err

}

func GetPostDetail(id int64) (*models.Post, error) {

	//查询帖子信息
	post, err := models.GETDetailOfPost(id)
	return post, err

	//根据作者id查询作者信息

}
