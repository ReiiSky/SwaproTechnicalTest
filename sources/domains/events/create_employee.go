package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/entities"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type CreateEmployee struct {
	Employee   entities.Employee
	Position   *entities.Position
	Department *entities.Department
	SuperiorID *objects.Identifier[int]
}

func (evt CreateEmployee) Eventname() string {
	return "CreateEmployee"
}
