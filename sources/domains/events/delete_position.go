package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type DeletePosition struct {
	ID objects.Identifier[int]
	objects.Changelog
}

func (evt DeletePosition) Eventname() string {
	return "DeletePosition"
}
