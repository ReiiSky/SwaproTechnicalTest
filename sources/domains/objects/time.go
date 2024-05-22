package objects

import "time"

type SwaproTime struct {
	t        time.Time
	platform time.Time
}

func NewSwaproTime(t time.Time) SwaproTime {
	return SwaproTime{
		t: t,
	}
}

func (swaproTime SwaproTime) ToISO() string {
	return swaproTime.t.Format(time.RFC3339)
}

func (swaproTime SwaproTime) PlatformToISO() string {
	return swaproTime.platform.Format(time.RFC3339)
}

func (swaproTime SwaproTime) SetPlatformTime(platform time.Time) SwaproTime {
	return SwaproTime{
		swaproTime.t,
		platform,
	}
}
