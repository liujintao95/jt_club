package dal

import (
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
)

func SelectUserThroughEmail(email string) (models.Users, error) {
	var (
		user models.Users
		err  error
		sql  = "select * from user where email=?"
	)
	err = db.Conn.Get(&user, sql, email)
	return user, err
}

func InsertUser(user *models.Users, passwordHash string) error {
	var (
		err error
		sql = "insert into user(uid, name, email, password) values (?,?,?,?)"
	)
	_, err = db.Conn.Exec(sql, user.Uid, user.Name, user.Email, passwordHash)
	return err
}

func SelectUserThroughUid(uid string) (*models.Users, error) {
	var (
		user *models.Users
		err  error
		sql  = "select * from user where uid=?"
	)
	err = db.Conn.Get(&user, sql, uid)
	return user, err
}

func SelectGroupUser(gid string) ([]*models.Users, error) {
	var (
		users []*models.Users
		err   error
		sql   = "SELECT `user`.* FROM `user` INNER JOIN user_group_map WHERE user_group_map.gid = ?"
	)
	err = db.Conn.Get(&users, sql, gid)
	return users, err
}
