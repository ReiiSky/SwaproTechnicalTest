package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type UpdatePosition struct {
	ID      objects.Identifier[int]
	NewName string
	objects.Changelog
}

func (evt UpdatePosition) Eventname() string {
	return "UpdatePosition"
}
