package userMysql

import (
	"awesomeProject/models"
	"awesomeProject/pkg/md5"
	"awesomeProject/pkg/snowflake"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

func SignUp() {

}

func QueryUserByUserID(id int64) (*models.User, error) {
	user, err := models.FindOneByUserId(id)
	return user, err
}

func InsertUser(p *models.ParamSignUp) error {
	user := new(models.User)
	user.Username = p.Username
	user.Password = p.Password
	user.UserID = snowflake.GetSnowflakeId()
	//密码加密
	user.Password = md5.ToMd5(user.Password)

	err := models.SaveOne(user)
	if err != nil {
		return err
	}
	return err
}

func CheckUserExist(p string) (bool, error, models.User) {
	var user models.User
	err, b, user := models.FindByName(p)

	if b {
		return true, err, user
	}
	return false, err, user

}
