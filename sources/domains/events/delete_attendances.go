package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type DeleteAttendances struct {
	EmployeeID objects.Identifier[int]
}

func (evt DeleteAttendances) Eventname() string {
	return "DeleteAttendances"
}
