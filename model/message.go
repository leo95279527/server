package model

import "time"

type Message struct {
	Id       uint `gorm:"AUTO_INCREMENT;primary_key;index"`
	TokenId  string
	Message  string
	Title    string
	Priority int
	Date     time.Time
}