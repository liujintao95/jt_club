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
		sql  = "select * from user where email=? and deleted = false"
	)
	err = session.Get(&user, sql, email)
	return user, err
}

func SetUser(session db.Session, user *models.User, passwordHash string) error {
	var (
		err error
		sql = "insert into user(uid, name, email, password) values (?,?,?,?)"
	)
	_, err = session.Exec(sql, user.Uid, user.Name, user.Email, passwordHash)
	return err
}

func SelectUsers(session db.Session, query string) ([]*models.User, error) {
	var (
		user      []*models.User
		err       error
		likeQuery = "%" + query + "%"
		sql       = `
			select * 
			from user 
			where deleted = false
			and	(
			    uid like ?
			    or email like ?
			)
		`
	)
	err = session.Select(&user, sql, likeQuery, likeQuery)
	return user, err
}

func SelectContacts(session db.Session, currentUser *models.User, query string) ([]*models.User, error) {
	var (
		contacts  []*models.User
		err       error
		likeQuery = "%" + query + "%"
		sql       = `
			select user.* 
			from user
			inner join user_contacts
			on user.uid = user_contacts.contact_id
			and	user_contacts.contact_type = ?
			and user_contacts.deleted = false
			where user.deleted = false
			and	(
			    user.uid like ?
			    or user.email like ?
			)
			and user_contacts.uid = ?
		`
	)
	err = session.Select(&contacts, sql, constant.ContactsUserType, likeQuery, likeQuery, currentUser.Uid)
	return contacts, err
}

func GetUserByUid(session db.Session, uid string) (*models.User, error) {
	var (
		user *models.User
		err  error
		sql  = `select * from user where uid=? and deleted = false`
	)
	err = session.Get(&user, sql, uid)
	return user, err
}

func GetUserContacts(session db.Session, uid string) ([]*models.UserContacts, error) {
	var (
		contacts []*models.UserContacts
		err      error
		sql      = `
			select * 
			from user_contacts 
			left join user 
			on user_contacts.contact_type = ?
			and user_contacts.contact_id = user.uid
			and user.deleted = false
			left join user_group 
			on user_contacts.contact_type = ?
			and user_contacts.contact_id = user_group.gid
			and	user_group.deleted = false
			where user_contacts.uid=?
			and	user_contacts.deleted = false
		`
	)
	err = session.Select(&contacts, sql, constant.ContactsUserType, constant.ContactsGroupType, uid)
	return contacts, err
}

func SetContactRequest(session db.Session, ContactsRequest *models.ContactsRequest) error {
	var (
		err error
		sql = `
			insert into contacts_request(
				request_id, uid, contact_id, contact_type, status, notice
			) values (?,?,?,?,?,?)`
	)
	_, err = session.Exec(sql,
		ContactsRequest.RequestId, ContactsRequest.Uid, ContactsRequest.ContactId,
		ContactsRequest.ContactType, ContactsRequest.Status, ContactsRequest.Notice,
	)
	return err
}
