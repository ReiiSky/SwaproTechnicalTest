package persistences

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ReiiSky/SwaproTechnical/sources/applications"
	"github.com/ReiiSky/SwaproTechnical/sources/domains"
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/environments"
	_ "github.com/lib/pq"
)

type PostgreSQLConfig struct {
	Host       string
	Port       int
	Name, Pass string
	DB, Opts   string
}

type SpecImplFn func(context.Context, domains.ISpecification) []applications.Aggregate
type EventImplFn func(context.Context, domains.IEvent) error

type specsInScope map[string]SpecImplFn

type PostgreContainer struct {
	env              environments.EnvVarContainer
	currentSpecScope string
	specImpls        map[string]specsInScope
	eventImpls       map[string]EventImplFn
}

func NewPostgreContainer(env environments.EnvVarContainer) PostgreContainer {
	if env == nil {
		panic(errors.New("error env is nil in postgre container"))
	}

	// test availability of postgres config
	env.GetString("POSTGRE_HOST")
	env.GetString("POSTGRE_NAME")
	env.GetString("POSTGRE_PASS")
	env.GetString("POSTGRE_DB")
	env.GetString("POSTGRE_OPTS")
	env.GetInt("POSTGRE_PORT")

	return PostgreContainer{
		env,
		ScopeEmpty,
		make(map[string]specsInScope),
		make(map[string]EventImplFn),
	}
}

const (
	ScopeEmpty    = ""
	ScopeEmployee = "Employee"
)

func (pct PostgreContainer) Scope(scope string) PostgreContainer {
	pct.currentSpecScope = scope

	return pct
}

type specImpl interface {
	domains.ISpecification
	Fn() SpecImplFn
}

type eventImpl interface {
	domains.IEvent
	Fn() EventImplFn
}

func (pct PostgreContainer) AddSpecImpl(impl specImpl) PostgreContainer {
	if len(pct.currentSpecScope) <= 0 {
		panic(ErrEmptyScope)
	}

	currentSpecImpls, ok := pct.specImpls[pct.currentSpecScope]

	if !ok {
		currentSpecImpls = map[string]SpecImplFn{}
	}

	currentSpecImpls[impl.Specname()] = impl.Fn()
	pct.specImpls[pct.currentSpecScope] = currentSpecImpls

	return pct
}

func (pct PostgreContainer) AddEventImpl(impl eventImpl) PostgreContainer {
	pct.eventImpls[impl.Eventname()] = impl.Fn()

	return pct
}

func (pct PostgreContainer) New(ctx context.Context) RepositoriesContext {
	config := PostgreSQLConfig{
		Host: pct.env.GetString("POSTGRE_HOST"),
		Name: pct.env.GetString("POSTGRE_NAME"),
		Port: pct.env.GetInt("POSTGRE_PORT"),
		Pass: pct.env.GetString("POSTGRE_PASS"),
		DB:   pct.env.GetString("POSTGRE_DB"),
		Opts: pct.env.GetString("POSTGRE_OPTS"),
	}

	return &PostgreContext{
		ctx,
		nil,
		pct.specImpls,
		pct.eventImpls,
		config,
	}
}

type PostgreContext struct {
	ctx            context.Context
	db             *sql.DB
	allSpecScopes  map[string]specsInScope
	allEventScopes map[string]EventImplFn
	dbConfig       PostgreSQLConfig
}

func (pctx *PostgreContext) initDBOnce() {
	if pctx.db == nil {
		db, _ := sql.Open("postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s %s",
			pctx.dbConfig.Host,
			pctx.dbConfig.Port,
			pctx.dbConfig.Name,
			pctx.dbConfig.Pass,
			pctx.dbConfig.DB,
			pctx.dbConfig.Opts,
		))

		err := db.Ping()

		if err != nil {
			panic(err)
		}

		pctx.ctx = context.WithValue(pctx.ctx, DBContextKey, db)
		pctx.db = db
	}
}

func (pctx PostgreContext) Close() {
	if pctx.db != nil {
		// TODO: catch error here
		pctx.db.Close()
	}
}

type PostgreQueryImpl struct {
	specs specsInScope
	ctx   context.Context
}

func (qimpl PostgreQueryImpl) GetOne(
	spec domains.ISpecification,
) applications.Aggregate {
	fn, ok := qimpl.specs[spec.Specname()]

	if !ok {
		panic(ErrSpecNotImplemented)
	}

	aggregates := fn(qimpl.ctx, spec)

	if len(aggregates) <= 0 {
		return nil
	}

	return aggregates[0]
}

func (qimpl PostgreQueryImpl) Get(
	spec domains.ISpecification,
) []applications.Aggregate {
	fn, ok := qimpl.specs[spec.Specname()]

	if !ok {
		panic(ErrSpecNotImplemented)
	}

	return fn(qimpl.ctx, spec)
}

func (pctx *PostgreContext) Save(
	aggrs applications.Aggregate,
) {
	pctx.initDBOnce()

	events := aggrs.Events()

	tx, err := pctx.db.Begin()

	if err != nil {
		panic(err)
	}

	eventContext := context.WithValue(pctx.ctx, TXContextKey, tx)

	for _, event := range events {
		evt := event.Top()
		eventFn, ok := pctx.allEventScopes[evt.Eventname()]

		if !ok {
			panic(ErrEventNotImplemented)
		}

		err := eventFn(eventContext, evt)

		if err == nil {
			continue
		}

		tx.Rollback()
		return
	}

	tx.Commit()
}

func (pctx *PostgreContext) Employee() applications.QueryRepository {
	pctx.initDBOnce()

	return PostgreQueryImpl{
		specs: pctx.allSpecScopes[ScopeEmployee],
		ctx:   pctx.ctx,
	}
}
