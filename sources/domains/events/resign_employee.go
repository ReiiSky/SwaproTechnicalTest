package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type ResignEmployee struct {
	ID objects.Identifier[int]
	objects.Changelog
}

func (evt ResignEmployee) Eventname() string {
	return "ResignEmployee"
}
