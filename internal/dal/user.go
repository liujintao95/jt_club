package dal

import (
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
)

func GetUserByEmail(session db.Session, email string) (models.User, error) {
	var (
		user models.User
		err  error
		sql  = "select * from user where email=?"
	)
	err = session.Get(&user, sql, email)
	return user, err
}

func SaveUser(session db.Session, user *models.User, passwordHash string) error {
	var (
		err error
		sql = "insert into user(uid, name, email, password) values (?,?,?,?)"
	)
	_, err = session.Exec(sql, user.Uid, user.Name, user.Email, passwordHash)
	return err
}

func GetUserByUid(session db.Session, uid string) (*models.User, error) {
	var (
		user *models.User
		err  error
		sql  = "select * from user where uid=?"
	)
	err = session.Get(&user, sql, uid)
	return user, err
}

func GetUserContacts(session db.Session, uid string) ([]*models.UserContacts, error) {
	var (
		contacts []*models.UserContacts
		err      error
		sql      = `
			select * from user_contacts 
			left join user 
			on user_contacts.contact_type = ?
			and user_contacts.contact_id = user.uid
			left join user_group 
			on user_contacts.contact_type = ?
			and user_contacts.contact_id = user_group.gid
			where user_contacts.uid=?
		`
	)
	err = session.Select(&contacts, sql, constant.ContactsUserType, constant.ContactsGroupType, uid)
	return contacts, err
}
