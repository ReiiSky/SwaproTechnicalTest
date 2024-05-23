package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type DeleteEmployee struct {
	ID objects.Identifier[int]
	objects.Changelog
}

func (evt DeleteEmployee) Eventname() string {
	return "DeleteEmployee"
}
