package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Position struct {
	Entity[int]
	departmentID objects.Identifier[int]
	name         string
	changelog    objects.Changelog
}

func NewPosition(
	id,
	departmentID int,
	name string,
	changelog objects.ChangelogParam,
) Position {
	return Position{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		objects.NewIdentifier(departmentID),
		name,
		objects.NewChangelog(changelog),
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
