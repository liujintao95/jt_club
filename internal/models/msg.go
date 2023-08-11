package models

import "time"

type Message struct {
	Id           int64
	MessageId    string
	Avatar       string
	FromUsername string
	From         string
	To           string
	Content      string
	ContentType  int32
	Type         string
	MessageType  int32
	Url          string
	FileSuffix   string
	FilePath     string
	Ctime        time.Time
	Utime        time.Time
	Deleted      int
}
