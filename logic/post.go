package logic

import (
	"awesomeProject/dao/communityMysql"
	"awesomeProject/dao/mysql/userMysql"
	"awesomeProject/dao/postMysql"
	"awesomeProject/models"
	"awesomeProject/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	//1，生成post id
	p.ID = snowflake.GetSnowflakeId()

	//2.保存到数据库
	err := postMysql.CreatPost(p)

	//返回响应
	return err
}

func GetPostDetailHandler(id int64) (*models.ApiPostDetail, error) {
	//查询并组合我们接口想用的数据

	//查询帖子接口信息
	post, err := postMysql.GetPostDetail(id)
	if err != nil {
		zap.L().Error("err", zap.Error(err))
	}
	//查询作者信息
	user, err := userMysql.QueryUserByUserID(post.AuthorID)
	if err != nil {
		zap.L().Error("err", zap.Error(err))
	}

	//查询社区信息
	communityDetail, err := communityMysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("err", zap.Error(err))
	}
	data := new(models.ApiPostDetail)
	data.Post = post
	data.CommunityDetail = communityDetail
	data.AuthorName = user.Username
	return data, err

}
