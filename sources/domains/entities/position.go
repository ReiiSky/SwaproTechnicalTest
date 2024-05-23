package entities

import (
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
)

type Position struct {
	Entity[int]
	departmentID  objects.Identifier[int]
	name          string
	changelog     objects.Changelog
	employeeCount int
}

func NewPosition(
	id,
	departmentID int,
	name string,
	changelog objects.ChangelogParam,
	employeeCount int,
) Position {
	return Position{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		objects.NewIdentifier(departmentID),
		name,
		objects.NewChangelog(changelog),
		employeeCount,
	}
}

func (p Position) ID() objects.Identifier[int] {
	return p.identifier
}

func (p Position) Name() string {
	return p.name
}

func (p Position) Changelog() objects.Changelog {
	return p.changelog
}

func (p *Position) ChangeName(name string) {
	p.name = name
}

type PositionInfo struct {
	ID            int
	Name          string
	EmployeeCount int
	CreatedAt     objects.SwaproTime
	UpdatedAt     *objects.SwaproTime
}

func (p Position) Info() PositionInfo {
	return PositionInfo{
		objects.GetNumberIdentifier(p.identifier),
		p.name,
		p.employeeCount,
		p.changelog.CreatedAt(),
		p.changelog.UpdatedAt(),
	}
}
