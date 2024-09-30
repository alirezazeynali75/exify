package outbox

import "time"

type Outbox struct {
	ID        uint64
	Payload   string
	Topic     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time 
}
