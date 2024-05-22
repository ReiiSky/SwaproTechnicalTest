package entities

import "github.com/ReiiSky/SwaproTechnical/sources/domains/objects"

type Entity[T string | int] struct {
	identifier objects.Identifier[T]
}

func (e Entity[T]) ID() objects.Identifier[T] {
	return e.identifier
}
