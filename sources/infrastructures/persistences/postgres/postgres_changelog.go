package postgres

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

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
