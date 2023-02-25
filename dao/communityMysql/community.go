package communityMysql

import "awesomeProject/models"

func GetCommunityLis() (*[]models.Community, error) {
	allCommunity, err := models.FindAllCommunity()
	community := allCommunity
	return community, err

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	allCommunityDetails, err := models.FindallCommunityDetails(id)
	return allCommunityDetails, err
}
