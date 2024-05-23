package entities

import (
	"strings"

	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Department struct {
	Entity[int]
	name          string
	changelog     objects.Changelog
	employeeCount int
	positionCount int
}

func NewDepartment(
	id int,
	name string,
	changelog objects.ChangelogParam,
	employeeCount,
	positionCount int,
) Department {
	return Department{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		name,
		objects.NewChangelog(changelog),
		employeeCount,
		positionCount,
	}
}

func (d Department) NameEqual(name string) bool {
	return strings.Compare(d.name, name) == 0
}

func (d Department) ID() objects.Identifier[int] {
	return d.identifier
}

func (d Department) Name() string {
	return d.name
}

func (d Department) Changelog() objects.Changelog {
	return d.changelog
}

func (d *Department) ChangeName(name string) {
	d.name = name
}

type DepartmentInfo struct {
	ID            int
	Name          string
	EmployeeCount int
	PositionCount int
	CreatedAt     objects.SwaproTime
	UpdatedAt     *objects.SwaproTime
}

func (d Department) Info() DepartmentInfo {
	return DepartmentInfo{
		objects.GetNumberIdentifier(d.identifier),
		d.name,
		d.employeeCount,
		d.positionCount,
		d.changelog.CreatedAt(),
		d.changelog.UpdatedAt(),
	}
}
