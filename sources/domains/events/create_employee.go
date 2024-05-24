package events

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/entities"
)

type CreateEmployee struct {
	Employee entities.Employee
}

func (evt CreateEmployee) Eventname() string {
	return "CreateEmployee"
}
