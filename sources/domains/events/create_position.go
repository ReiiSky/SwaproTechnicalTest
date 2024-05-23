package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type CreateOrUsePosition struct {
	DepartmentID objects.Identifier[int]
	PositionName string
}

func (evt CreateOrUsePosition) Eventname() string {
	return "CreateOrUsePosition"
}
