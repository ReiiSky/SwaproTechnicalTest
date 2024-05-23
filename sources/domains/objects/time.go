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

func NewSwaproTimeNow() SwaproTime {
	return SwaproTime{
		t: time.Now(),
	}
}

func (swaproTime SwaproTime) ToISOUTC() string {
	return swaproTime.t.UTC().Format(time.RFC3339)
}
