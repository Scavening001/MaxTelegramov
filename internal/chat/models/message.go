package models

import "time"

type Message struct {
	Type     string    `json:"type"`
	Username string    `json:"username"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}
