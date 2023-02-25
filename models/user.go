package models

import "awesomeProject/dao/mysql"

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func FindOneByUserId(userId int64) (*User, error) {
	var user User
	err := mysql.Db.Where("user_id=?", userId).First(&user).Error
	return &user, err
}

func FindByName(name string) (error, bool, User) {
	var user User
	err := mysql.Db.Where("username=?", name).First(&user).Error
	if user.UserID != 0 {
		return err, true, user
	}
	return err, false, user
}

//func FindByPhone(phone string) (UserIm2, bool) {
//	var user UserIm2
//	initialization.InitMysql().Where("phone=?", phone).First(&user)
//	if user.ID != 0 {
//		return user, true
//	}
//	return user, false
//}
//
//func FindByEmail(email string) UserIm2 {
//	var user UserIm2
//	initialization.InitMysql().Where("email=?", email).First(&user)
//	return user
//}
//
//func FindUserList() []*UserIm2 {
//	var users []*UserIm2
//	users = make([]*UserIm2, 20)
//	initialization.InitMysql().Find(&users)
//	return users
//}

func SaveOne(user *User) error {

	err := mysql.Db.Create(user).Error
	return err

}
