package dal

import (
	"JT_CLUB/internal/models"
	"JT_CLUB/pkg/db"
	"fmt"
)

func InsertMsg(msg *models.Message) error {
	sql := `
		insert into message(
			message_id, avatar, from_username, 'from', 'to', content, content_type,
            type, message_type, url, file_suffix, file_path
        ) values (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	_, err := db.Conn.Exec(sql,
		msg.MessageId, msg.Avatar, msg.FromUsername, msg.From, msg.To, msg.Content,
		msg.ContentType, msg.Type, msg.MessageType, msg.Url, msg.FileSuffix, msg.FilePath,
	)
	if err != nil {
		return fmt.Errorf("insert user exec: %w", err)
	}
	return nil
}
