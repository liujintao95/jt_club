package dal

import (
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
)

func SaveMsg(session db.Session, msg *models.Message) error {
	var (
		err error
		sql = `
			insert into message(
				message_id, avatar, from_username, 'from', 'to', content, content_type,
				type, message_type, url, file_suffix, file_path
			) values (?,?,?,?,?,?,?,?,?,?,?,?)
		`
	)
	_, err = session.Exec(sql,
		msg.MessageId, msg.Avatar, msg.FromUsername, msg.From, msg.To, msg.Content,
		msg.ContentType, msg.Type, msg.MessageType, msg.Url, msg.FileSuffix, msg.FilePath,
	)
	return err
}
