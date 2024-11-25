package models

import "time"

type History struct {
	Activity  string    `json:"activity"`
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	Time      time.Time `json:"time"`
}
