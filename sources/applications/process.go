package applications

import (
	"github.com/ReiiSky/SwaproTechnical/sources/applications/services"
)

type Process interface {
	Repositories() Repositories
	Services() Services
}

type Services interface {
	Auth() services.Auth
	Hasher() services.Hasher
}

type Repositories interface {
	Employee() QueryRepository
	CommandRepository
}
