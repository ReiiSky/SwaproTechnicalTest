package main

import (
	"time"

	"github.com/ReiiSky/SwaproTechnical/sources/applications/services/auth"
	"github.com/ReiiSky/SwaproTechnical/sources/applications/services/hash"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/environments"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences/postgres"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/persistences/postgresevent"
	"github.com/ReiiSky/SwaproTechnical/sources/interfaces/adapters"
)

func FiberAPI(env environments.EnvVarContainer) adapters.Fiber {
	repositories := persistences.NewPostgreContainer(env).
		Scope(persistences.ScopeEmployee).
		AddSpecImpl(postgres.GetByEmployeeID{}).
		AddSpecImpl(postgres.GetByEmpty{}).
		AddSpecImpl(postgres.GetByEmployeeName{}).
		AddEventImpl(postgresevent.RegisterImpl{})

	kernelParam := infrastructures.KernelParam{
		ContextDuration: time.Second * time.Duration(env.GetInt("PROCESS_TIMEOUT_IN_SECOND")),
		Auth:            auth.NewJWTAuthentication(env.GetString("AUTH_SECRET")),
		Hasher:          hash.NewMD5(env.GetString("PASSWORD_SECRET")),
		Repositories:    repositories,
	}

	kernel := infrastructures.NewKernel(kernelParam)
	return adapters.NewFiber(kernel, env.GetInt("SERVER_PORT"))
}
