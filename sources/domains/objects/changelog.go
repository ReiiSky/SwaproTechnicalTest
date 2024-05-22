package objects

import "time"

type ChangelogParam struct {
	CreatedAt time.Time
	CreatedBy string
	Update    *struct {
		At time.Time
		By string
	}
	DeletedAt *time.Time
}

type updateLog struct {
	at time.Time
	by string
}

type Changelog struct {
	createdAt time.Time
	createdBy string
	update    *updateLog
	deletedAt *time.Time
}

func NewChangelog(param ChangelogParam) Changelog {
	var update *updateLog

	if param.Update != nil {
		update = &updateLog{
			at: param.Update.At,
			by: param.Update.By,
		}
	}

	return Changelog{
		createdAt: param.CreatedAt,
		createdBy: param.CreatedBy,
		update:    update,
		deletedAt: param.DeletedAt,
	}
}
