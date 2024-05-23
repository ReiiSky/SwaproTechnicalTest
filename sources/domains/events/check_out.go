package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type CheckOut struct {
	AttendanceID objects.Identifier[int]
	objects.Changelog
}

func (evt CheckOut) Eventname() string {
	return "CheckOut"
}
