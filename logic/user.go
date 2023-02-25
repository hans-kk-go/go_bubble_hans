package logic

import (
	"awesomeProject/dao/mysql/userMysql"
	"awesomeProject/models"
	"awesomeProject/pkg/jwt_"
	"awesomeProject/pkg/md5"
	"errors"
)

// 存放业务逻辑的代码

func SignUp(p1 *models.ParamSignUp) error {
	//判断用户存不存在
	userExist, _, _ := userMysql.CheckUserExist(p1.Username)
	exist := userExist
	//if err2 != nil {
	//	return err2
	//}
	if exist {
		return errors.New("用户已存在")
	}
	//生成uid

	//3.密码要加密

	//4，保存进数据库
	err := userMysql.InsertUser(p1)
	if err != nil {
		return err
	}
	return err
}

func Login(p *models.ParamLogin) (err error, token string) {
	var user models.User
	opassword := p.Password
	//判断用户存不存在
	userExist, err2, user := userMysql.CheckUserExist(p.Username)
	password := user.Password
	exist := userExist
	if err2 != nil {
		return err2, ""
	}
	if !exist {
		return errors.New("用户不存在"), ""
	}
	//判断密码是否正确
	password1 := md5.ToMd5(opassword)
	if password != password1 {
		return errors.New("密码错误"), ""
	}

	//生成token
	token1 := jwt_.JwtGettoken(user.UserID)

	return err, token1

}
