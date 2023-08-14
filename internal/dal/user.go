package dal

import (
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
	"fmt"
)

func SelectUserThroughEmail(email string) (models.Users, error) {
	var user models.Users
	sql := "select * from user where email=?"
	err := db.Conn.Get(&user, sql, email)
	if err != nil {
		return user, fmt.Errorf("select user: %w", err)
	}
	return user, nil
}

func InsertUser(user *models.Users, passwordHash string) error {
	sql := "insert into user(uid, name, email, password) values (?,?,?,?)"
	_, err := db.Conn.Exec(sql, user.Uid, user.Name, user.Email, passwordHash)
	if err != nil {
		return fmt.Errorf("insert user exec: %w", err)
	}
	return nil
}

func SelectUserThroughUid(uid string) (*models.Users, error) {
	var user *models.Users
	sql := "select * from user where uid=?"
	err := db.Conn.Get(&user, sql, uid)
	if err != nil {
		return user, fmt.Errorf("select user: %w", err)
	}
	return user, nil
}

func SelectGroupUser(gid string) ([]*models.Users, error) {
	var users []*models.Users
	sql := "SELECT `user`.* FROM `user` INNER JOIN user_group_map WHERE user_group_map.gid = ?"
	err := db.Conn.Get(&users, sql, gid)
	if err != nil {
		return users, fmt.Errorf("select group users: %w", err)
	}
	return users, nil
}
