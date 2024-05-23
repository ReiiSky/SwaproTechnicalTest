package objects

import "time"

type SwaproTime struct {
	t time.Time
}

func NewSwaproTime(t time.Time) SwaproTime {
	return SwaproTime{
		t: t,
	}
}

func (swaproTime SwaproTime) ToISOUTC() string {
	return swaproTime.t.UTC().Format(time.RFC3339)
}
