package events

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type UpdateSuperior struct {
	EmployeeID objects.Identifier[int]
	SuperiorID objects.Identifier[int]
}

func (evt UpdateSuperior) Eventname() string {
	return "UpdateSuperior"
}
