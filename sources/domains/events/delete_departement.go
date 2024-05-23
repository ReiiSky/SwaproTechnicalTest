package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type DeleteDepartment struct {
	ID objects.Identifier[int]
	objects.Changelog
}

func (evt DeleteDepartment) Eventname() string {
	return "DeleteDepartment"
}
