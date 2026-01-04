package endpoint

import (
	"time"

	"github.com/google/uuid"
)

type Endpoint struct {
	ID                 uuid.UUID
	Name               string
	Token              string
	CreatedAt          time.Time
	UserID             uuid.UUID
	NotificationEnable bool
}
