package persistences

import (
	"errors"

	"github.com/ReiiSky/SwaproTechnical/sources/applications"
)

var (
	ErrEmptyScope          = errors.New("scope should not be empty")
	ErrSpecNotImplemented  = errors.New("specs not implemented")
	ErrEventNotImplemented = errors.New("event not implemented")
)

type ContextKeyType int

const (
	DBContextKey ContextKeyType = iota
	TXContextKey ContextKeyType = iota
)

type RepositoriesContext interface {
	applications.Repositories
	Close()
}
