package benefit

import "database/sql"

type Cycle struct {
	CycleId        sql.NullInt64
	StartPeriod    sql.NullTime
	Active         sql.NullBool
	CreatedDate    sql.NullTime
	ModifiedDate   sql.NullTime
	ModifiedUser   sql.NullString
	BenefitSetting sql.NullString
	Title          sql.NullString
	Visible        sql.NullBool
}

func (cycle *Cycle) GetCycleId() int64 {
	return cycle.CycleId.Int64
}

func (cycle *Cycle) GetBenefitSetting() string {
	return cycle.BenefitSetting.String
}

func (cycle *Cycle) IsValid() bool {
	return cycle.CycleId.Valid && (cycle != &Cycle{})
}
