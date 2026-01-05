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
	NotificationEnable bool // false의 경우 push하지않고 notification 테이블에만 데이터를 넣음
}
