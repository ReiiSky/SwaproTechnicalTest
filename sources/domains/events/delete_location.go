package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type DeleteLocation struct {
	ID objects.Identifier[int]
	objects.Changelog
}

func (evt DeleteLocation) Eventname() string {
	return "DeleteLocation"
}
