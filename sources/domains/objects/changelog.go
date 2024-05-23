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
	at SwaproTime
	by string
}

type Changelog struct {
	createdAt SwaproTime
	createdBy string
	update    *updateLog
	deletedAt *SwaproTime
}

func NewChangelog(param ChangelogParam) Changelog {
	var update *updateLog
	var deletedAt *SwaproTime

	if param.Update != nil {
		update = &updateLog{
			at: NewSwaproTime(param.Update.At),
			by: param.Update.By,
		}
	}

	if param.DeletedAt != nil {
		d := NewSwaproTime(*param.DeletedAt)
		deletedAt = &d
	}

	return Changelog{
		createdAt: NewSwaproTime(param.CreatedAt),
		createdBy: param.CreatedBy,
		update:    update,
		deletedAt: deletedAt,
	}
}

func NewCreateChangelog(createdByEmployeeCode string) Changelog {
	now := time.Now()

	return Changelog{
		createdAt: NewSwaproTime(now),
		createdBy: createdByEmployeeCode,
	}
}
