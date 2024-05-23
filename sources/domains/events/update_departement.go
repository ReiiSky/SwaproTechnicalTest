package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type UpdateDepartment struct {
	ID      objects.Identifier[int]
	NewName string
	objects.Changelog
}

func (evt UpdateDepartment) Eventname() string {
	return "UpdateDepartment"
}
