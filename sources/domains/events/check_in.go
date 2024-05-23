package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type CheckIn struct {
	EmployeeID objects.Identifier[int]
	Name       string
	In         objects.SwaproTime
	objects.Changelog
}

func (evt CheckIn) Eventname() string {
	return "CheckIn"
}
