package infrastructures

import (
	"context"
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/services"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces"
)

type RepositoriesContainer interface {
	New(ctx context.Context) persistences.RepositoriesContext
}

type Kernel struct {
	contextDuration     time.Duration
	repositoryContainer RepositoriesContainer
	auth                services.Auth
	hasher              services.Hasher
}

type KernelParam struct {
	ContextDuration time.Duration
	Repositories    RepositoriesContainer
	Auth            services.Auth
	Hasher          services.Hasher
}

func NewKernel(param KernelParam) interfaces.IKernel {
	return Kernel{
		contextDuration:     param.ContextDuration,
		repositoryContainer: param.Repositories,
		auth:                param.Auth,
		hasher:              param.Hasher,
	}
}

func (k Kernel) NewProcess() interfaces.StopableProcess {
	ctx, cancel := context.WithTimeout(context.Background(), k.contextDuration)

	go func() {
		select {
		case <-time.After(k.contextDuration):
			cancel()
		}
	}()

	kernelContext := Process{
		ctx:               ctx,
		repositoryContext: k.repositoryContainer.New(ctx),
		auth:              k.auth,
		hasher:            k.hasher,
	}

	return kernelContext
}

type Process struct {
	ctx               context.Context
	auth              services.Auth
	hasher            services.Hasher
	repositoryContext persistences.RepositoriesContext
}

func (p Process) Repositories() applications.Repositories {
	return p.repositoryContext
}

type Service struct {
	auth   services.Auth
	hasher services.Hasher
}

func (s Service) Auth() services.Auth {
	return s.auth
}

func (s Service) Hasher() services.Hasher {
	return s.hasher
}

func (p Process) Services() applications.Services {
	return Service{
		auth:   p.auth,
		hasher: p.hasher,
	}
}

func (p Process) Ctx() context.Context {
	return p.ctx
}

func (p Process) Close() {
	p.repositoryContext.Close()
}
