package mysql

import (
	"database/sql"
	"lovemarket/models"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "liwenzhou.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(userId) from user where userName = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	/*	user.Password = encryptPassword(user.Password)*/
	// 执行SQL语句入库
	sqlStr := `insert into user( username, password , email) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.Username, user.Password, user.Email)
	return
}

/*// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum(nil))
}*/

func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select userId, userName, password from user where userName=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	/*	password := encryptPassword(oPassword)*/
	if oPassword != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select userId, userName from user where userId = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
