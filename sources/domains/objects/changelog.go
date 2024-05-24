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
	return Changelog{
		createdAt: NewSwaproTimeNow(),
		createdBy: createdByEmployeeCode,
	}
}

func (c Changelog) UpdatedNow(updatedByEmployeeCode string) Changelog {
	c.update = &updateLog{
		at: NewSwaproTimeNow(),
		by: updatedByEmployeeCode,
	}

	return c
}

func (c Changelog) DeletedNow() Changelog {
	t := NewSwaproTimeNow()
	c.deletedAt = &t

	return c
}

func (c Changelog) CreatedAt() SwaproTime {
	return c.createdAt
}

func (c Changelog) CreatedBy() string {
	return c.createdBy
}

func (c Changelog) UpdatedAt() *SwaproTime {
	if c.update == nil {
		return nil
	}

	return &c.update.at
}

func (c Changelog) UpdatedBy() *string {
	if c.update == nil {
		return nil
	}

	return &c.update.by
}

func (c Changelog) DeletedAt() *SwaproTime {
	if c.deletedAt == nil {
		return nil
	}

	return c.deletedAt
}
