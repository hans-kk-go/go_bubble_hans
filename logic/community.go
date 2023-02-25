package logic

import (
	"awesomeProject/dao/communityMysql"
	"awesomeProject/models"
)

func GetCommunityList() (*[]models.Community, error) {
	//查数据库 查找到所有的comminity 并返回

	communitylist, err := communityMysql.GetCommunityLis()
	//if err != nil {
	//	return nil,err
	//}
	return communitylist, err

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	//查数据库 查找到所有的communityDetail 并返回
	data, err := communityMysql.GetCommunityDetail(id)
	return data, err
}
