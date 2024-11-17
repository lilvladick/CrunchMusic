package models

import (
	"time"

	"github.com/lib/pq"
)

type Timestamp struct {
	Time time.Time
}

func (t *Timestamp) ToNullTime() pq.NullTime {
	return pq.NullTime{
		Time:  t.Time,
		Valid: !t.Time.IsZero(),
	}
}
