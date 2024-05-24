package postgres

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/ReiiSky/SwaproTechnical/sources/domains/objects"
	"github.com/jmoiron/sqlx"
)

func upsertChangelog(ctx context.Context, db *sql.DB, logname string, id int, log objects.Changelog) {
	var (
		updatedAt *string
		updatedBy *string
		deletedAt *string
	)

	if log.UpdatedAt() != nil {
		u := log.UpdatedAt().ToISOUTC()
		updatedAt = &u
	}

	if log.UpdatedBy() != nil {
		updatedBy = log.UpdatedBy()
	}

	if log.DeletedAt() != nil {
		da := log.DeletedAt().ToISOUTC()
		deletedAt = &da
	}

	sq.Insert("changelog").
		Columns("log_name", "id", "created_at", "created_by").
		Values(logname, id, log.CreatedAt(), log.CreatedBy()).
		Suffix(fmt.Sprintf(
			"on conflict (log_name, id) do update set updated_at = '%v',"+
				"updated_by = '%v', deleted_at = '%v'",
			updatedAt, updatedBy, deletedAt),
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		ExecContext(ctx)
}

type changelog struct {
	LogName   string  `db:"log_name"`
	ID        int     `db:"id"`
	CreatedAt string  `db:"created_at"`
	CreatedBy string  `db:"created_by"`
	UpdatedAt *string `db:"updated_at"`
	UpdatedBy *string `db:"updated_by"`
	DeletedAt *string `db:"deleted_at"`
}

var changelogColumns = []string{"log_name", "id", "created_at", "created_by", "updated_at", "updated_by", "deleted_at"}

func getChangelog(ctx context.Context, db *sql.DB, logname string, id int) *changelog {
	result, _ := sq.Select(changelogColumns...).
		From("changelog").
		Where(sq.Eq{"log_name": logname}).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		QueryContext(ctx)

	changelogs := []changelog{}
	sqlx.StructScan(result, &changelogs)

	if len(changelogs) <= 0 {
		return nil
	}

	return &changelogs[0]
}
