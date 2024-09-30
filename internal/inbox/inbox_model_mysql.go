package inbox

import "time"

type InboxModel struct {
	EventID string    `gorm:"column:event_id;primaryKey"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (InboxModel) TableName() string {
	return "inboxs"
}
