package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Membership struct {
	Entity[int]
	name      string
	password  objects.CryptedInformation
	address   string
	isActive  bool
	changelog objects.Changelog
}

func NewMembership(
	id int,
	name,
	password,
	address string,
	isActive bool,
	changelog objects.ChangelogParam,
) Membership {
	return Membership{
		Entity[int]{
			identifier: objects.NewIdentifier(id),
		},
		name,
		objects.NewMD5CryptedInformation(password),
		address,
		isActive,
		objects.NewChangelog(changelog),
	}
}

func (mem Membership) ID() objects.Identifier[int] {
	return mem.identifier
}

func (mem Membership) Password() objects.CryptedInformation {
	return mem.password
}

func (mem Membership) Changelog() objects.Changelog {
	return mem.changelog
}
