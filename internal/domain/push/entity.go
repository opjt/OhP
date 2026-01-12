package push

import "github.com/google/uuid"

type Subscription struct {
	UserID   uuid.UUID
	Endpoint string
	P256dh   string
	Auth     string
}

type Token struct {
	Endpoint string
	P256dh   string
	Auth     string
}
