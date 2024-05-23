package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type CreateOrUsePosition struct {
	// Can be use department id or name, for
	// attaching position name, if not department exist and using name.
	// New Department name will be created.
	DepartmentID   objects.Identifier[int]
	DepartmentName string

	PositionName string
	objects.Changelog
}

func (evt CreateOrUsePosition) Eventname() string {
	return "CreateOrUsePosition"
}
