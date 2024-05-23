package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type UpdateLocation struct {
	ID      objects.Identifier[int]
	NewName string
	objects.Changelog
}

func (evt UpdateLocation) Eventname() string {
	return "UpdateLocation"
}
