package models

import "time"

type TokenBlacklist struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
