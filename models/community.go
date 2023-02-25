/*
   @Author: StudentCWZ
   @Description:
   @File: community
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/22 15:24
*/

package models

import (
	"awesomeProject/dao/mysql"
	"fmt"
	"time"
)

type Community struct {
	ID   int64  `json:"id" gorm:"community_id"`
	Name string `json:"name" gorm:"column:community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" gorm:"community_id"`
	Name         string    `json:"name" gorm:"column:community_name"`
	Introduction string    `json:"introduction,omitempty" gorm:"introduction"`
	CreateTime   time.Time `json:"create_time" gorm:"create_time"`
}

func FindAllCommunity() (*[]Community, error) {
	//var communities
	//Community := new([]Community)
	var Community []Community
	err := mysql.Db.Find(&Community).Error

	fmt.Println(Community)

	return &Community, err

}

//func FindUserList() []*UserIm2 {
//	var users []*UserIm2
//	users = make([]*UserIm2, 20)
//	initialization.InitMysql().Find(&users)
//	return users
//}

func FindallCommunityDetails(id int64) (*CommunityDetail, error) {
	var community *CommunityDetail
	//c := new(CommunityDetail)
	err := mysql.Db.Where("community_id=?", id).First(community).Error
	return community, err
}
